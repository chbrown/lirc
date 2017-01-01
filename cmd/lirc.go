package main

import (
	"flag"
	"io"
	"log"
	"strings"
	"github.com/sorcix/irc"
	"github.com/chbrown/lirc"
)

// create a (sync) receiving chan and read all messages from conn into it
func decoderChan(conn *irc.Conn) <-chan *irc.Message {
	c := make(chan *irc.Message)
	outputTransformer := lirc.NewTextTransformer()
	go func() {
		for {
			// conn.Decode() returns a (*Message, error)
			m, err := conn.Decode()
			if err != nil {
				// we log but otherwise swallow the error here
				if err == io.EOF {
					log.Println("IRC Connection reached EOF")
				} else {
					log.Println("IRC Connection Decode error", err)
				}
				close(c)
				// avoid adding the Message accompanying the error to the chan
				break
			}
			outputTransformer.Incoming(m)
			c <- m
		}
	}()
	return c
}

// create a sending chan that writes its messages to conn
func encoderChan(conn *irc.Conn) chan<- *irc.Message {
	c := make(chan *irc.Message)
	outputTransformer := lirc.NewTextTransformer()
	go func() {
		for {
			m := <-c
			outputTransformer.Outgoing(m)
			err := conn.Encode(m)
			if err != nil {
				log.Println("IRC Connection Encode error", err)
				close(c)
				break
			}
		}
	}()
	return c
}

// prefix the given channel with a # if it does not already start with # or &
func addDefaultPrefix(channel string) string {
	if strings.HasPrefix(channel, "#") {
		return channel
	}
	if strings.HasPrefix(channel, "&") {
		return channel
	}
	return "#" + channel
}

func main() {
	var addr, nick, user string
	var tls bool
	flag.StringVar(&addr, "addr", "irc.freenode.net", "server in 'hostname:port' format")
	flag.BoolVar(&tls, "tls", false, "connect to <addr> with SSL/TLS")
	flag.StringVar(&nick, "nick", "", "nickname")
	flag.StringVar(&user, "user", "anonymous", "username")
	flag.Parse()

	channels := flag.Args()

	conn, err := lirc.IrcDial(addr, tls)
	if err != nil {
		log.Panicln("IRC Connection Dial error", err)
	}
	inbox := decoderChan(conn)
	outbox := encoderChan(conn)

	// send desired nickname to server
	outbox <- &irc.Message{
		Command: irc.NICK,
		Params:  []string{nick},
	}

	// send username and "real" name to server
	outbox <- &irc.Message{
		Command:  irc.USER,
		Params:   []string{user, "8", "*"},
		Trailing: "Anonymous",
	}

	// main loop
	for m := range inbox {
		switch m.Command {
		case irc.PING:
			outbox <- &irc.Message{
				Command:  irc.PONG,
				Params:   m.Params,
				Trailing: m.Trailing,
			}
		case irc.RPL_WELCOME:
			for _, channel := range channels {
				// TODO: JOIN all channels at once
				outbox <- &irc.Message{
					Command: irc.JOIN,
					Params:  []string{addDefaultPrefix(channel)},
				}
			}
		}
	}
}
