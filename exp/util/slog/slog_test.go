package slog

import (
	"context"
	"net"
	"os"
	"testing"

	"golang.org/x/exp/slog"
)

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
