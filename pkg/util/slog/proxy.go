package slog

import (
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

const (
	TimeKey    = slog.TimeKey
	LevelKey   = slog.LevelKey
	MessageKey = slog.MessageKey
	SourceKey  = slog.SourceKey
)

type Level = slog.Level
type LevelVar = slog.LevelVar

const (
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
)

type Leveler = slog.Leveler

// record.go

type Record = slog.Record

func NewRecord(t time.Time, level Level, msg string, pc uintptr) Record {
	return slog.NewRecord(t, level, msg, pc)
}

// value.go

type Kind = slog.Kind
type Value = slog.Value

const (
	KindAny       = slog.KindAny
	KindBool      = slog.KindBool
	KindDuration  = slog.KindDuration
	KindFloat64   = slog.KindFloat64
	KindInt64     = slog.KindInt64
	KindString    = slog.KindString
	KindTime      = slog.KindTime
	KindUint64    = slog.KindUint64
	KindGroup     = slog.KindGroup
	KindLogValuer = slog.KindLogValuer
)
