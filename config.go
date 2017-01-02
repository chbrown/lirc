package lirc

import (
	"flag"
	"fmt"
	"strings"
)

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

type Config struct {
	Addr     string
	Tls      bool
	Nickname string
	Username string
	Realname string
	Channels []string
	Outputs  []string
}

func ParseFlags() *Config {
	addr_raw := flag.String("addr", "irc.freenode.net", "server in 'hostname:port' format")
	tls := flag.Bool("tls", false, "connect to <addr> with SSL/TLS")
	nickname := flag.String("nick", "lirc", "nickname")
	username := flag.String("user", "anonymous", "username")
	realname := flag.String("real", "Anonymous", "realname")
	out := flag.String("out", "text", "list of outputs, comma-separated")
	flag.Parse()

	// trim colon off the end of addr, just in case
	addr := strings.TrimRight(*addr_raw, ":")
	// add the default port
	addr = addDefaultPort(addr, ircDefaultPort(*tls))

	rest := flag.Args()
	// add a preceding # to the channel names that lack one
	channels := make([]string, len(rest))
	for i, channel := range rest {
		channels[i] = addDefaultPrefix(channel)
	}

	// split out on commas
	outputs := strings.Split(*out, ",")

	return &Config{
		addr,
		*tls,
		*nickname,
		*username,
		*realname,
		channels,
		outputs,
	}
}
