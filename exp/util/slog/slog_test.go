package slog

import (
	"context"
	"net"
	"os"
	"testing"

	"golang.org/x/exp/slog"
)

//

type Handler = slog.Handler
type Level = slog.Level
type Attr = slog.Attr

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

type DefaultLogger struct {
	l *slog.Logger
}

//

type colorHandler struct {
	level slog.Level
}

func (h *colorHandler) clone() *colorHandler {
	return &colorHandler{
		level: h.level,
	}
}

var _ slog.Handler = &colorHandler{}

func (h *colorHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.level
}

func (h *colorHandler) Handle(ctx context.Context, record slog.Record) error {
	panic("implement me")
}

func (h *colorHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	panic("implement me")
}

func (h *colorHandler) WithGroup(name string) slog.Handler {
	panic("implement me")
}

//

func TestSlog(t *testing.T) {
	slog.SetDefault(slog.New(
		//slog.NewTextHandler(os.Stderr),
		slog.NewJSONHandler(os.Stderr),
	))

	slog.Info("hello", "name", "Al")
	slog.Error("oops", "err", net.ErrClosed, "status", 500)

	slog.LogAttrs(nil, slog.LevelError, "oops",
		slog.Any("err", net.ErrClosed), slog.Int("status", 500))
}
