package slog

import (
	"context"
	"fmt"
)

//

func Log(ctx context.Context, level Level, msg string, args ...any) {
	DefaultCtx(ctx).log(level, ActionNone, msg, args...)
}

func LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr) {
	l := DefaultCtx(ctx)
	log(
		l.DefaultContext(),
		l.DefaultHandler(),
		level,
		ActionNone,
		0,
		msg,
		func(r *Record) {
			r.AddAttrs(attrs...)
		},
	)
}

func Debug(msg string, args ...any) {
	Default().log(LevelDebug, ActionNone, msg, args...)
}

func DebugCtx(ctx context.Context, msg string, args ...any) {
	DefaultCtx(ctx).log(LevelDebug, ActionNone, msg, args...)
}

func Info(msg string, args ...any) {
	Default().log(LevelInfo, ActionNone, msg, args...)
}

func InfoCtx(ctx context.Context, msg string, args ...any) {
	DefaultCtx(ctx).log(LevelInfo, ActionNone, msg, args...)
}

func Warn(msg string, args ...any) {
	Default().log(LevelWarn, ActionNone, msg, args...)
}

func WarnCtx(ctx context.Context, msg string, args ...any) {
	DefaultCtx(ctx).log(LevelWarn, ActionNone, msg, args...)
}

func Error(msg string, args ...any) {
	Default().log(LevelError, ActionNone, msg, args...)
}

func ErrorCtx(ctx context.Context, msg string, args ...any) {
	DefaultCtx(ctx).log(LevelError, ActionNone, msg, args...)
}

//

func Panic(obj any, args ...any) {
	var msg string
	var ok bool
	if msg, ok = obj.(string); !ok {
		msg = fmt.Sprint(obj)
	}
	Default().log(LevelError, ActionPanic, msg, args...)
}

func Fatal(obj any, args ...any) {
	var msg string
	var ok bool
	if msg, ok = obj.(string); !ok {
		msg = fmt.Sprint(obj)
	}
	Default().log(LevelError, ActionFatal, msg, args...)
}

func IfError(err error, args ...any) {
	if err != nil {
		Default().log(LevelError, ActionNone, ErrorMsg, append(args, E(err))...)
	}
}

func IfPanic(err error, args ...any) {
	if err != nil {
		Default().log(LevelError, ActionPanic, ErrorMsg, append(args, E(err))...)
	}
}

func IfFatal(err error, args ...any) {
	if err != nil {
		Default().log(LevelError, ActionFatal, ErrorMsg, append(args, E(err))...)
	}
}

func OrError(fn func() error, args ...any) {
	if err := fn(); err != nil {
		Default().log(LevelError, ActionNone, ErrorMsg, append(args, E(err))...)
	}
}

func OrPanic(fn func() error, args ...any) {
	if err := fn(); err != nil {
		Default().log(LevelError, ActionPanic, ErrorMsg, append(args, E(err))...)
	}
}

func OrFatal(fn func() error, args ...any) {
	if err := fn(); err != nil {
		Default().log(LevelError, ActionFatal, ErrorMsg, append(args, E(err))...)
	}
}

//

func Print(v ...any) {
	Default().log(LevelInfo, ActionNone, fmt.Sprint(v...))
}

func Printf(format string, v ...any) {
	Default().log(LevelInfo, ActionNone, fmt.Sprintf(format, v...))
}

func Println(v ...any) {
	Default().log(LevelInfo, ActionNone, fmt.Sprintln(v...))
}

func Fatalf(format string, v ...any) {
	Default().log(LevelError, ActionFatal, fmt.Sprintf(format, v...))
}

func Fatalln(v ...any) {
	Default().log(LevelError, ActionFatal, fmt.Sprintln(v...))
}

func Panicf(format string, v ...any) {
	Default().log(LevelError, ActionPanic, fmt.Sprintf(format, v...))
}

func Panicln(v ...any) {
	Default().log(LevelError, ActionPanic, fmt.Sprintln(v...))
}
