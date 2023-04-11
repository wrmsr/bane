package slog

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"

	"golang.org/x/exp/slog"
)

//

type Action int8

const (
	ActionNone Action = iota
	ActionPanic
	ActionFatal
)

func (a Action) String() string {
	switch a {
	case ActionNone:
		return "none"
	case ActionPanic:
		return "panic"
	case ActionFatal:
		return "fatal"
	}
	panic(a)
}

//

type NullHandler struct{}

var _ Handler = NullHandler{}

func (h NullHandler) Enabled(ctx context.Context, level slog.Level) bool   { return false }
func (h NullHandler) Handle(ctx context.Context, record slog.Record) error { return nil }

func (h NullHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return h }
func (h NullHandler) WithGroup(name string) slog.Handler       { return h }

//

const BadKey = "!BADKEY"

func ArgsToAttr(args []any) (Attr, []any) {
	switch x := args[0].(type) {
	case string:
		if len(args) == 1 {
			return String(BadKey, x), nil
		}
		a := Any(x, args[1])
		a.Value = a.Value.Resolve()
		return a, args[2:]

	case Attr:
		x.Value = x.Value.Resolve()
		return x, args[1:]

	default:
		return Any(BadKey, x), args[1:]
	}
}

//

func DefaultHandler() Handler {
	// FIXME: cache
	return slog.Default().Handler()
}

//

type Logger struct {
	h   Handler
	ctx context.Context
}

func NewLogger(h Handler) Logger {
	return Logger{
		h: h,
	}
}

func Default() Logger {
	return NewLogger(DefaultHandler())
}

func DefaultCtx(ctx context.Context) Logger {
	return NewLogger(DefaultHandler()).WithContext(ctx)
}

func (l Logger) Handler() Handler {
	return l.h
}

func (l Logger) DefaultHandler() Handler {
	if h := l.h; h != nil {
		return h
	}
	if h := DefaultHandler(); h != nil {
		return h
	}
	return NullHandler{}
}

func (l Logger) Context() context.Context {
	return l.ctx
}

func (l Logger) DefaultContext() context.Context {
	if l.ctx != nil {
		return l.ctx
	}
	return context.Background()
}

//

func (l Logger) With(args ...any) Logger {
	var (
		attr  Attr
		attrs []Attr
	)
	for len(args) > 0 {
		attr, args = ArgsToAttr(args)
		attrs = append(attrs, attr)
	}
	l.h = l.h.WithAttrs(attrs)
	return l
}

func (l Logger) WithGroup(name string) Logger {
	l.h = l.h.WithGroup(name)
	return l
}

func (l Logger) WithContext(ctx context.Context) Logger {
	l.ctx = ctx
	return l
}

func (l Logger) Enabled(level Level) bool {
	return l.h.Enabled(l.DefaultContext(), level)
}

//

type LogPanic struct {
	*Record
}

var _ error = LogPanic{}

func (l LogPanic) Error() string {
	return fmt.Sprintf("%+v", l.Record)
}

//

var IgnorePC = false

func log(
	ctx context.Context,
	h Handler,
	level Level,
	action Action,
	pco int,
	msg string,
	prepRec func(*Record),
) {
	if !h.Enabled(ctx, level) {
		return
	}

	var pc uintptr
	if !IgnorePC {
		var pcs [1]uintptr
		runtime.Callers(3+pco, pcs[:])
		pc = pcs[0]
	}

	r := slog.NewRecord(time.Now(), level, msg, pc)
	prepRec(&r)
	_ = h.Handle(ctx, r)

	switch action {
	case ActionPanic:
		panic(r)
	case ActionFatal:
		os.Exit(1)
	}
}

//

func (l Logger) log(
	level Level,
	action Action,
	msg string,
	args ...any,
) {
	log(
		l.DefaultContext(),
		l.DefaultHandler(),
		level,
		action,
		1,
		msg,
		func(r *Record) {
			var attr Attr
			for len(args) > 0 {
				attr, args = ArgsToAttr(args)
				r.AddAttrs(attr)
			}
		},
	)
}

func (l Logger) Log(level Level, msg string, args ...any) {
	l.log(level, ActionNone, msg, args...)
}

func (l Logger) LogAttrs(level Level, msg string, attrs ...Attr) {
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

func (l Logger) Debug(msg string, args ...any) { l.log(LevelDebug, ActionNone, msg, args...) }
func (l Logger) Info(msg string, args ...any)  { l.log(LevelInfo, ActionNone, msg, args...) }
func (l Logger) Warn(msg string, args ...any)  { l.log(LevelWarn, ActionNone, msg, args...) }
func (l Logger) Error(msg string, args ...any) { l.log(LevelError, ActionNone, msg, args...) }
func (l Logger) Panic(msg string, args ...any) { l.log(LevelError, ActionPanic, msg, args...) }
func (l Logger) Fatal(msg string, args ...any) { l.log(LevelError, ActionFatal, msg, args...) }

func (l Logger) IfError(err error, args ...any) {}
func (l Logger) IfPanic(err error, args ...any) {}
func (l Logger) IfFatal(err error, args ...any) {}

func (l Logger) OrError(fn func() error, args ...any) {}
func (l Logger) OrPanic(fn func() error, args ...any) {}
func (l Logger) OrFatal(fn func() error, args ...any) {}
