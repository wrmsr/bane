package log

import (
	"os"
	"time"

	"github.com/wrmsr/bane/pkg/util/check"
)

//

type Arg struct {
	Name  string
	Value any
}

func A(name string, value any) Arg {
	return Arg{Name: name, Value: value}
}

func Err(err error) Arg {
	return Arg{Name: "error", Value: err}
}

func Recovery(r any) Arg {
	return Arg{Name: "recovery", Value: r}
}

//

type BaseLogger interface {
	Log(lvl Level, msg string, args ...Arg)
}

type baseLoggerImpl struct {
	ll LineLogger

	clock func() time.Time
}

var _ BaseLogger = baseLoggerImpl{}

func (l baseLoggerImpl) now() time.Time {
	if l.clock != nil {
		return l.clock()
	}
	return time.Now()
}

func (l baseLoggerImpl) Log(lvl Level, msg string, args ...Arg) {
	line := Line{
		Time:  l.now(),
		Level: lvl,

		Message: msg,
		Args:    args,
	}

	err := l.ll.LogLine(line)
	if err != nil {
		panic(err)
	}

	switch lvl {
	case PanicLevel:
		panic(line)
	case FatalLevel:
		os.Exit(1)
	}
}

//

type Logger interface {
	BaseLogger

	Debug(msg string, args ...Arg)
	Info(msg string, args ...Arg)
	Warn(msg string, args ...Arg)
	Error(msg string, args ...Arg)
	Panic(msg string, args ...Arg)
	Fatal(msg string, args ...Arg)

	IfError(err error, args ...Arg)
	IfPanic(err error, args ...Arg)
	IfFatal(err error, args ...Arg)

	OrError(fn func() error, args ...Arg)
	OrPanic(fn func() error, args ...Arg)
	OrFatal(fn func() error, args ...Arg)
}

type loggerImpl struct {
	BaseLogger
}

func NewLogger(ll LineLogger) Logger {
	return loggerImpl{
		BaseLogger: baseLoggerImpl{
			ll: check.NotNil(ll).(LineLogger),
		},
	}
}

var _ Logger = loggerImpl{}

func (l loggerImpl) Debug(msg string, args ...Arg) { l.Log(DebugLevel, msg, args...) }
func (l loggerImpl) Info(msg string, args ...Arg)  { l.Log(InfoLevel, msg, args...) }
func (l loggerImpl) Warn(msg string, args ...Arg)  { l.Log(WarnLevel, msg, args...) }
func (l loggerImpl) Error(msg string, args ...Arg) { l.Log(ErrorLevel, msg, args...) }
func (l loggerImpl) Panic(msg string, args ...Arg) { l.Log(PanicLevel, msg, args...) }
func (l loggerImpl) Fatal(msg string, args ...Arg) { l.Log(FatalLevel, msg, args...) }

func (l loggerImpl) IfError(err error, args ...Arg) {
	if err != nil {
		l.Log(ErrorLevel, "error", append(args, Err(err))...)
	}
}

func (l loggerImpl) IfPanic(err error, args ...Arg) {
	if err != nil {
		l.Log(PanicLevel, "error", append(args, Err(err))...)
	}
}

func (l loggerImpl) IfFatal(err error, args ...Arg) {
	if err != nil {
		l.Log(FatalLevel, "error", append(args, Err(err))...)
	}
}

func (l loggerImpl) OrError(fn func() error, args ...Arg) {
	if err := fn(); err != nil {
		l.Log(ErrorLevel, "error", append(args, Err(err))...)
	}
}

func (l loggerImpl) OrPanic(fn func() error, args ...Arg) {
	if err := fn(); err != nil {
		l.Log(PanicLevel, "error", append(args, Err(err))...)
	}
}

func (l loggerImpl) OrFatal(fn func() error, args ...Arg) {
	if err := fn(); err != nil {
		l.Log(FatalLevel, "error", append(args, Err(err))...)
	}
}
