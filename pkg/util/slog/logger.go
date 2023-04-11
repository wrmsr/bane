package slog

import (
	"context"
	"fmt"
)

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

//

func (l Logger) Panic(obj any, args ...any) {
	var msg string
	var ok bool
	if msg, ok = obj.(string); !ok {
		msg = fmt.Sprint(obj)
	}
	l.log(LevelError, ActionPanic, msg, args...)
}

func (l Logger) Fatal(obj any, args ...any) {
	var msg string
	var ok bool
	if msg, ok = obj.(string); !ok {
		msg = fmt.Sprint(obj)
	}
	l.log(LevelError, ActionFatal, msg, args...)
}

func (l Logger) IfError(err error, args ...any) {
	if err != nil {
		l.log(LevelError, ActionNone, ErrorMsg, append(args, E(err))...)
	}
}

func (l Logger) IfPanic(err error, args ...any) {
	if err != nil {
		l.log(LevelError, ActionPanic, ErrorMsg, append(args, E(err))...)
	}
}

func (l Logger) IfFatal(err error, args ...any) {
	if err != nil {
		l.log(LevelError, ActionFatal, ErrorMsg, append(args, E(err))...)
	}
}

func (l Logger) OrError(fn func() error, args ...any) {
	if err := fn(); err != nil {
		l.log(LevelError, ActionNone, ErrorMsg, append(args, E(err))...)
	}
}

func (l Logger) OrPanic(fn func() error, args ...any) {
	if err := fn(); err != nil {
		l.log(LevelError, ActionPanic, ErrorMsg, append(args, E(err))...)
	}
}

func (l Logger) OrFatal(fn func() error, args ...any) {
	if err := fn(); err != nil {
		l.log(LevelError, ActionFatal, ErrorMsg, append(args, E(err))...)
	}
}
