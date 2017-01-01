package lirc

import (
	"github.com/sorcix/irc"
)

type Transformer interface {
	Incoming(*irc.Message) *irc.Message
	Outgoing(*irc.Message) *irc.Message
}
