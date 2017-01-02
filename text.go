package lirc

import (
	"fmt"
	"log"
	"github.com/sorcix/irc"
)

type TextListener struct {
	Listener
}

// no options
func NewTextListener() *TextListener {
	return &TextListener{}
}

func fmtPrefix(p *irc.Prefix) string {
	if p == nil {
		return "∅"
	}
	return fmt.Sprintf("%s!%s@%s", p.Name, p.User, p.Host)
}

func fmtMessage(m *irc.Message) string {
	return fmt.Sprintf("%s %s %v: %s", fmtPrefix(m.Prefix), m.Command, m.Params, m.Trailing)
}

func (_ *TextListener) Incoming(m *irc.Message) {
	log.Printf("-> %s   \n", fmtMessage(m))
}
func (_ *TextListener) Outgoing(m *irc.Message) {
	log.Printf("   %s ->\n", fmtMessage(m))
}
