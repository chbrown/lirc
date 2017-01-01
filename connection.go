package lirc

import (
	"fmt"
	"strings"
	"net"
	"crypto/tls"
	"github.com/sorcix/irc"
)

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

// create a new irc.Conn, adding the default IRC port to addr if missing, depending on use_tls
func IrcDial(addr string, use_tls bool) (*irc.Conn, error) {
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
	return irc.NewConn(tcp_conn), nil
}
