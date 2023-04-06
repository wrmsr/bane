package slog

import (
	"context"
	"time"

	"golang.org/x/exp/slog"
)

// attr.go

type Attr = slog.Attr

func String(key, value string) Attr             { return slog.String(key, value) }
func Int64(key string, value int64) Attr        { return slog.Int64(key, value) }
func Int(key string, value int) Attr            { return slog.Int(key, value) }
func Uint64(key string, v uint64) Attr          { return slog.Uint64(key, v) }
func Float64(key string, v float64) Attr        { return slog.Float64(key, v) }
func Bool(key string, v bool) Attr              { return slog.Bool(key, v) }
func Time(key string, v time.Time) Attr         { return slog.Time(key, v) }
func Duration(key string, v time.Duration) Attr { return slog.Duration(key, v) }
func Group(key string, as ...Attr) Attr         { return slog.Group(key, as...) }
func Any(key string, value any) Attr            { return slog.Any(key, value) }

// handler.go

type Handler = slog.Handler

// level.go

type Level = slog.Level
type LevelVar = slog.LevelVar

// logger.go

func Debug(msg string, args ...any)                         { slog.Debug(msg, args...) }
func DebugCtx(ctx context.Context, msg string, args ...any) { slog.DebugCtx(ctx, msg, args...) }

func Info(msg string, args ...any)                         { slog.Info(msg, args...) }
func InfoCtx(ctx context.Context, msg string, args ...any) { slog.InfoCtx(ctx, msg, args...) }

func Warn(msg string, args ...any)                         { slog.Warn(msg, args...) }
func WarnCtx(ctx context.Context, msg string, args ...any) { slog.WarnCtx(ctx, msg, args...) }

func Error(msg string, args ...any)                         { slog.Error(msg, args...) }
func ErrorCtx(ctx context.Context, msg string, args ...any) { slog.ErrorCtx(ctx, msg, args...) }

func Log(ctx context.Context, level Level, msg string, args ...any) {
	slog.Log(ctx, level, msg, args...)
}

func LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr) {
	slog.LogAttrs(ctx, level, msg, attrs...)
}

// record.go

type Record = slog.Record

// value.go

type Kind = slog.Kind
type Value = slog.Value
