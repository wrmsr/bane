package log

import (
	"fmt"

	bat "github.com/wrmsr/bane/pkg/utils/atomic"
)

//

func newDefaultLogger() Logger {
	return NewLogger(
		NewLineLogger(
			TextFormatter{},
			StdWriter{},
		),
	)
}

var _global = bat.NewLazy(newDefaultLogger)

func global() Logger {
	return _global.Get()
}

//

func Log(lvl Level, msg string, args ...Arg) { global().Log(lvl, msg, args...) }

func Debug(msg string, args ...Arg) { global().Debug(msg, args...) }
func Info(msg string, args ...Arg)  { global().Info(msg, args...) }
func Warn(msg string, args ...Arg)  { global().Warn(msg, args...) }
func Error(msg string, args ...Arg) { global().Error(msg, args...) }

func Panic(msg any, args ...Arg) {
	if msg, ok := msg.(string); ok {
		global().Panic(msg, args...)
		return
	}
	global().Panic(fmt.Sprint(msg), args...)
}

func Fatal(msg any, args ...Arg) {
	if msg, ok := msg.(string); ok {
		global().Fatal(msg, args...)
		return
	}
	global().Fatal(fmt.Sprint(msg), args...)
}

func OrError(fn func() error, args ...Arg) { global().OrError(fn, args...) }
func OrPanic(fn func() error, args ...Arg) { global().OrError(fn, args...) }
func OrFatal(fn func() error, args ...Arg) { global().OrError(fn, args...) }

//

func Print(v ...any)                 { Info(fmt.Sprint(v...)) }
func Printf(format string, v ...any) { Info(fmt.Sprintf(format, v...)) }
func Println(v ...any)               { Info(fmt.Sprintln(v...)) }

func Fatalf(format string, v ...any) { Fatal(fmt.Sprintf(format, v...)) }
func Fatalln(v ...any)               { Fatal(fmt.Sprintln(v...)) }

func Panicf(format string, v ...any) { Panic(fmt.Sprintf(format, v...)) }
func Panicln(v ...any)               { Panic(fmt.Sprintln(v...)) }
