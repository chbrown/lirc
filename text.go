package lirc

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/sorcix/irc"
	"io"
	"os"
)

var (
	red     = color.New(color.FgRed).SprintFunc()
	green   = color.New(color.FgGreen).SprintFunc()
	yellow  = color.New(color.FgYellow).SprintFunc()
	blue    = color.New(color.FgBlue).SprintFunc()
	magenta = color.New(color.FgMagenta).SprintFunc()
	cyan    = color.New(color.FgCyan).SprintFunc()
	// with backgrounds:
	whiteOnBlack = color.New(color.FgBlack, color.BgWhite).SprintFunc()
	blackOnWhite = color.New(color.FgWhite, color.BgBlack).SprintFunc()
)

type TextListener struct {
	Listener
	w io.Writer
}

func NewTextListener() Listener {
	return &TextListener{w: os.Stdout}
}

func fmtPrefix(p *irc.Prefix) string {
	if p == nil {
		return "∅"
	}
	return fmt.Sprintf("%s!%s@%s", red(p.Name), magenta(p.User), p.Host)
}

func (listener *TextListener) printMessage(m *irc.Message, direction string) {
	prefix := fmtPrefix(m.Prefix)
	message := fmt.Sprintf("%s %s %s: %s", prefix, yellow(m.Command), cyan(m.Params), m.Trailing)
	fmt.Fprintf(listener.w, "%s %s %s\n", direction, timestamp(), message)
}

func (listener *TextListener) Incoming(m *irc.Message) {
	listener.printMessage(m, whiteOnBlack("<-"))
}
func (listener *TextListener) Outgoing(m *irc.Message) {
	listener.printMessage(m, blackOnWhite("->"))
}
