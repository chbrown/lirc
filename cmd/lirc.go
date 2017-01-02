package main

import (
	"flag"
	"log"
	"github.com/sorcix/irc"
	"github.com/chbrown/lirc"
)

// reads from oldc, runs a function, writes to new chan
func listenChan(oldc chan *irc.Message, fn func (*irc.Message)) chan *irc.Message {
	c := make(chan *irc.Message)
	go func() {
		for {
			m := <-oldc
			fn(m)
			c <- m
		}
	}()
	return c
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

	outputListener := lirc.NewJsonListener()

	// set up pipeline
	inbox := make(chan *irc.Message)
	conn.ReadToChan(inbox)
	// wrap inbox in a listener
	inbox = listenChan(inbox, outputListener.Incoming)

	outbox := make(chan *irc.Message)
	// wrap outbox in a listener
	deliverbox := listenChan(outbox, outputListener.Outgoing)
	conn.WriteFromChan(deliverbox)

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
					Params:  []string{lirc.AddDefaultPrefix(channel)},
				}
			}
		}
	}
}
