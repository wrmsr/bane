package slog

import (
	"context"
)

// logger.go

func Log(ctx context.Context, level Level, msg string, args ...any) {
	DefaultCtx(ctx).log(level, ActionNone, msg, args...)
}

func LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr) {
	//DefaultCtx(ctx).LogAttrs(level, msg, attrs...)
	panic("FIXME")
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
