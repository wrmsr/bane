package slog

import (
	"io"
	old "log"
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

func InstallOldAdapter(o *old.Logger, l Logger, level Level) {
	o.SetOutput(oldLogAdapter{l: l, level: level})
}
