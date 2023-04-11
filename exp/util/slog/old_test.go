package slog

import (
	"io"
	old "log"
	"os"
	"testing"

	stdslog "golang.org/x/exp/slog"
)

type oldLogAdapter struct {
	l     Logger
	level Level
}

var _ io.Writer = oldLogAdapter{}

func (a oldLogAdapter) Write(p []byte) (n int, err error) {
	p2 := p
	if len(p2) > 0 && p2[len(p2)-1] == '\n' {
		p2 = p2[:len(p2)-1]
	}
	log(
		a.l.DefaultContext(),
		a.l.DefaultHandler(),
		a.level,
		ActionNone,
		2,
		string(p2),
		nil,
	)
	return len(p), nil
}

func TestOld(t *testing.T) {
	stdslog.SetDefault(stdslog.New(
		//slog.NewTextHandler(os.Stderr),
		//slog.NewJSONHandler(os.Stderr),
		stdslog.HandlerOptions{
			AddSource: true,
		}.NewJSONHandler(os.Stderr),
	))

	old.SetFlags(0)
	old.SetOutput(oldLogAdapter{})
	old.Println("hi")
}
