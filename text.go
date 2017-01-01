package lirc

import (
	"fmt"
	"log"
	"github.com/sorcix/irc"
)

type TextTransformer struct {
	Transformer
}

// no options
func NewTextTransformer() *TextTransformer {
	return &TextTransformer{}
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

func (_ *TextTransformer) Incoming(m *irc.Message) *irc.Message {
	log.Printf("-> %s   \n", fmtMessage(m))
	return m
}
func (_ *TextTransformer) Outgoing(m *irc.Message) *irc.Message {
	log.Printf("   %s ->\n", fmtMessage(m))
	return m
}
