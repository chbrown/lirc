package lirc

import (
	"os"
	"fmt"
	"encoding/json"
	"github.com/sorcix/irc"
)

type JsonListener struct {
	Listener
}

// no options
func NewJsonListener() *JsonListener {
	return &JsonListener{}
}

func (_ *JsonListener) Incoming(m *irc.Message) {
	switch m.Command {
    case irc.PRIVMSG, irc.JOIN, irc.QUIT, irc.NICK:
		action := ToChannelAction(m)
		json_bytes, err := json.Marshal(action)
		if err != nil {
			fmt.Println("json.Marshal error:", err)
		}
		os.Stdout.Write(json_bytes)
		os.Stdout.WriteString("\n")
	}
}
func (_ *JsonListener) Outgoing(m *irc.Message) {
}
