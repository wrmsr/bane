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

const TimeKey = slog.TimeKey
const LevelKey = slog.LevelKey
const MessageKey = slog.MessageKey
const SourceKey = slog.SourceKey

type Level = slog.Level
type LevelVar = slog.LevelVar

const LevelDebug = slog.LevelDebug
const LevelInfo = slog.LevelInfo
const LevelWarn = slog.LevelWarn
const LevelError = slog.LevelError

type Leveler = slog.Leveler

// record.go

type Record = slog.Record

// value.go

type Kind = slog.Kind
type Value = slog.Value
