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

type loggerWriter struct {
	l  Logger
	lv Level
}

var _ io.Writer = loggerWriter{}

func (a loggerWriter) Write(p []byte) (n int, err error) {
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

func InstallOldWriter(o *old.Logger, l Logger, level Level) {
	if isSlogDefaultHandler(l.Handler()) {
		h := NewFormatterHandler(
			nil, // FIXME:
			old.Writer().Write,
			LevelInfo,
		)
		l = Logger{h: &h}
		panic("FIXME")
	}
	o.SetOutput(loggerWriter{l: l, lv: level})
}
