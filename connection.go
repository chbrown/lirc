package lirc

import (
	"io"
	"log"
	"fmt"
	"strings"
	"net"
	"crypto/tls"
	"github.com/sorcix/irc"
)

// prefix the given channel with a # if it does not already start with # or &
func AddDefaultPrefix(channel string) string {
	if strings.HasPrefix(channel, "#") {
		return channel
	}
	if strings.HasPrefix(channel, "&") {
		return channel
	}
	return "#" + channel
}

// return the default port for an IRC server, depending on TLS
func ircDefaultPort(use_tls bool) string {
	if use_tls {
		return "6697"
	}
	return "6667"
}

// add a :<port> suffix to addr if addr does not already contain a : character
func addDefaultPort(addr string, port string) string {
	if strings.Contains(addr, ":") {
		return addr
	}
	return fmt.Sprintf("%s:%s", addr, port)
}

// calls net.Dial or tls.Dial, depending on the use_tls argument
func netDial(network string, addr string, use_tls bool) (net.Conn, error) {
	if use_tls {
		// tls.Dial returns a (*tls.Conn, error)
		return tls.Dial(network, addr, &tls.Config{})
	}
	// net.Dial returns a (net.Conn, error)
	return net.Dial(network, addr)
}

type Conn struct {
	irc.Conn
}

// listen for and decode all messages from conn into c
func (conn *Conn) ReadToChan(c chan<- *irc.Message) {
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
			c <- m
		}
	}()
}

// write everything available on c to conn
func (conn *Conn) WriteFromChan(c <-chan *irc.Message) {
	go func() {
		for {
			m := <-c
			err := conn.Encode(m)
			if err != nil {
				log.Println("IRC Connection Encode error", err)
				break
			}
		}
	}()
}

// create a new irc.Conn, adding the default IRC port to addr if missing, depending on use_tls
func IrcDial(addr string, use_tls bool) (*Conn, error) {
	// trim colon off the end of addr, just in case
	addr = strings.TrimRight(addr, ":")
	// add the default port
	addr = addDefaultPort(addr, ircDefaultPort(use_tls))
	// establish TCP connection
	tcp_conn, err := netDial("tcp", addr, use_tls)
	if err != nil {
		return nil, err
	}
	// wrap in IRC line decoder/encoder
	irc_conn := *irc.NewConn(tcp_conn)
	return &Conn{irc_conn}, nil
}
