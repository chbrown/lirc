package lirc

import (
	"github.com/sorcix/irc"
)

type Listener interface {
	Incoming(*irc.Message)
	Outgoing(*irc.Message)
}

type Transformer interface {
	Incoming(*irc.Message) *irc.Message
	Outgoing(*irc.Message) *irc.Message
}
