package slog

import (
	"context"

	"golang.org/x/exp/slog"
)

//

type Logger interface {
	Handler() Handler

	With(args ...any) Logger
	WithGroup(name string) Logger

	Enabled(ctx context.Context, level Level) bool

	Log(ctx context.Context, level Level, msg string, args ...any)
	LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr)

	Debug(msg string, args ...any)
	DebugCtx(ctx context.Context, msg string, args ...any)

	Info(msg string, args ...any)
	InfoCtx(ctx context.Context, msg string, args ...any)

	Warn(msg string, args ...any)
	WarnCtx(ctx context.Context, msg string, args ...any)

	Error(msg string, args ...any)
	ErrorCtx(ctx context.Context, msg string, args ...any)
}

//

type slogLogger struct {
	*slog.Logger
}

var _ Logger = slogLogger{}

func (l slogLogger) With(args ...any) Logger {
	return slogLogger{l.Logger.With(args...)}
}

func (l slogLogger) WithGroup(name string) Logger {
	return slogLogger{l.Logger.WithGroup(name)}
}

//

func Default() Logger {
	// FIXME: atomic global, save alloc
	return slogLogger{slog.Default()}
}

//

type DefaultLogger struct {
	Logger Logger
}

func (l DefaultLogger) get() Logger {
	if l.Logger != nil {
		return l.Logger
	}
	return Default()
}

var _ Logger = DefaultLogger{}

func (l DefaultLogger) Handler() Handler {
	return l.get().Handler()
}

func (l DefaultLogger) With(args ...any) Logger {
	return DefaultLogger{l.get().With(args...)}
}

func (l DefaultLogger) WithGroup(name string) Logger {
	return DefaultLogger{l.get().WithGroup(name)}
}

func (l DefaultLogger) Enabled(ctx context.Context, level Level) bool {
	return l.get().Enabled(ctx, level)
}

func (l DefaultLogger) Log(ctx context.Context, level Level, msg string, args ...any) {
	l.get().Log(ctx, level, msg, args...)
}

func (l DefaultLogger) LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr) {
	l.get().LogAttrs(ctx, level, msg, attrs...)
}

func (l DefaultLogger) Debug(msg string, args ...any) {
	l.get().Debug(msg, args...)
}

func (l DefaultLogger) DebugCtx(ctx context.Context, msg string, args ...any) {
	l.get().DebugCtx(ctx, msg, args...)
}

func (l DefaultLogger) Info(msg string, args ...any) {
	l.get().Info(msg, args...)
}

func (l DefaultLogger) InfoCtx(ctx context.Context, msg string, args ...any) {
	l.get().InfoCtx(ctx, msg, args...)
}

func (l DefaultLogger) Warn(msg string, args ...any) {
	l.get().Warn(msg, args...)
}

func (l DefaultLogger) WarnCtx(ctx context.Context, msg string, args ...any) {
	l.get().WarnCtx(ctx, msg, args...)
}

func (l DefaultLogger) Error(msg string, args ...any) {
	l.get().Error(msg, args...)
}

func (l DefaultLogger) ErrorCtx(ctx context.Context, msg string, args ...any) {
	l.get().ErrorCtx(ctx, msg, args...)
}
