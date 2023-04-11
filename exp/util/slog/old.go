package slog

import (
	"io"
	old "log"
	"reflect"
)

func isSlogDefaultHandler(h Handler) bool {
	ty := reflect.ValueOf(h).Elem().Type()
	return ty.String() == "slog.defaultHandler"
}

type oldLogAdapter struct {
	l  Logger
	lv Level
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
		a.lv,
		ActionNone,
		2,
		string(p2),
		nil,
	)
	return len(p), nil
}

func InstallOldAdapter(o *old.Logger, l Logger, level Level) {
	o.SetOutput(oldLogAdapter{l: l, lv: level})
}
