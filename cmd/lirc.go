package main

import (
	"github.com/chbrown/lirc"
	"github.com/sorcix/irc"
	"log"
)

// reads from oldc, runs a function, writes to new chan
func listenChan(oldc chan *irc.Message, fn func(*irc.Message)) chan *irc.Message {
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

func contains(haystack []string, needle string) bool {
	for _, straw := range haystack {
		if straw == needle {
			return true
		}
	}
	return false
}

func main() {
	config := lirc.ParseFlags()

	conn, err := lirc.IrcDial(config.Addr, config.Tls)
	if err != nil {
		log.Panicln("IRC Connection Dial error", err)
	}

	var outputListener lirc.Listener
	if contains(config.Outputs, "text") {
		outputListener = lirc.NewTextListener()
	} else {
		outputListener = lirc.NewJsonListener()
	}

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
		Params:  []string{config.Nickname},
	}

	// send username and "real" name to server
	outbox <- &irc.Message{
		Command: irc.USER,
		// <mode> options: 12 = invisible+wallops, 8 = invisible, 4 = wallops, 0 = no special mode
		Params:   []string{config.Username, "8", "*"},
		Trailing: config.Realname,
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
			for _, channel := range config.Channels {
				// TODO: JOIN all channels at once
				outbox <- &irc.Message{
					Command: irc.JOIN,
					Params:  []string{channel},
				}
			}
		}
	}
}
