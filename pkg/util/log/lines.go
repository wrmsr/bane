package log

import (
	"time"

	"github.com/wrmsr/bane/pkg/util/check"
)

type Line struct {
	Time  time.Time
	Level Level

	Message string
	Args    []Arg

	StackOffset int
}

type LineLogger interface {
	LogLine(line Line) error
}

type lineLoggerImpl struct {
	f Formatter
	w Writer
}

func NewLineLogger(
	f Formatter,
	w Writer,
) LineLogger {
	return lineLoggerImpl{
		f: check.NotNil(f).(Formatter),
		w: check.NotNil(w).(Writer),
	}
}

var _ LineLogger = lineLoggerImpl{}

func (l lineLoggerImpl) LogLine(line Line) error {
	s, err := l.f.FormatLine(line)
	if err != nil {
		return err
	}
	return l.w.Write(s)
}
