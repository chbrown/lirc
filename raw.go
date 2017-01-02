package lirc

import (
	"fmt"
	"github.com/sorcix/irc"
	"io"
	"os"
	"time"
)

type RawListener struct {
	Listener
	w io.Writer
}

func NewRawListener() Listener {
	return &RawListener{w: os.Stdout}
}

func timestamp() string {
	return time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
}

func (listener *RawListener) Incoming(m *irc.Message) {
	fmt.Fprintf(listener.w, "<- %s %s\n", timestamp(), m.String())
}
func (listener *RawListener) Outgoing(m *irc.Message) {
	fmt.Fprintf(listener.w, "-> %s %s\n", timestamp(), m.String())
}
