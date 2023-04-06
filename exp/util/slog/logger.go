package slog

import "context"

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
