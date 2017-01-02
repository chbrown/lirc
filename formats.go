package lirc

import (
	"time"
	"github.com/sorcix/irc"
)

func toActor(p *irc.Prefix) *Actor {
	if p == nil {
		return nil
	}
	return &Actor{
		Name: p.Name,
		User: p.User,
		Host: p.Host,
	}
}

func ToChannelAction(m *irc.Message) *ChannelAction {
	var action_type ActionType
	var channel string
	var message string
	switch m.Command {
	case irc.PRIVMSG:
		action_type = ActionType_MESSAGE
		channel = m.Params[0]
		message = m.Trailing
	case irc.JOIN:
		action_type = ActionType_JOIN
		channel = m.Trailing
		// JOIN has no message
		message = ""
	case irc.QUIT:
		action_type = ActionType_QUIT
		// apparently QUIT does not specify the channel
		channel = ""
		message = m.Trailing
	case irc.NICK:
		action_type = ActionType_NICK
		channel = ""
		message = m.Trailing
	}
	return &ChannelAction{
		Timestamp: time.Now().Unix(),
		Actor: toActor(m.Prefix),
		Name: channel,
		Type: action_type,
		Message: message,
	}
}
