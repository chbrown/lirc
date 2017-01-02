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
			c <- m
			fn(m)
		}
	}()
	return c
}

var listenerFuncs = map[string](func() lirc.Listener){
	"raw":  lirc.NewRawListener,
	"text": lirc.NewTextListener,
	"json": lirc.NewJsonListener,
}

func main() {
	config := lirc.ParseFlags()

	conn, err := lirc.IrcDial(config.Addr, config.Tls)
	if err != nil {
		log.Panicln("IRC Connection Dial error", err)
	}

	listeners := make([]lirc.Listener, len(config.Outputs))
	for i, output := range config.Outputs {
		listeners[i] = listenerFuncs[output]()
	}

	// set up pipeline
	inbox := make(chan *irc.Message)
	conn.ReadToChan(inbox)
	// wrap inbox in a listener
	inbox = listenChan(inbox, func(m *irc.Message) {
		for _, listener := range listeners {
			listener.Incoming(m)
		}
	})

	outbox := make(chan *irc.Message)
	// wrap outbox in a listener
	deliverbox := listenChan(outbox, func(m *irc.Message) {
		for _, listener := range listeners {
			listener.Outgoing(m)
		}
	})
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
