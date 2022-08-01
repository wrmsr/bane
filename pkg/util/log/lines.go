package log

import (
	"time"

	"github.com/wrmsr/bane/pkg/util/check"
	rtu "github.com/wrmsr/bane/pkg/util/runtime"
)

type Line struct {
	Time  time.Time
	Level Level

	Message string
	Args    []Arg

	StackOffset int
}

func (l Line) AddStackOffset(n int) Line {
	l.StackOffset += n
	return l
}

func (l Line) GetStackFrame(ofs int) rtu.StackFrame {
	so := l.StackOffset + ofs + 1
	for _, a := range l.Args {
		if a, ok := a.(StackOffset); ok {
			so += a.Num
		}
	}
	return rtu.GetStackFrame(so)
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
	s, err := l.f.FormatLine(line.AddStackOffset(1))
	if err != nil {
		return err
	}
	return l.w.Write(s)
}
