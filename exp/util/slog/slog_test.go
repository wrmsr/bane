package slog

import (
	"context"
	"testing"

	"golang.org/x/exp/slog"
)

type colorHandler struct{}

var _ slog.Handler = &colorHandler{}

func (h colorHandler) Enabled(ctx context.Context, level slog.Level) bool {
	panic("implement me")
}

func (h colorHandler) Handle(ctx context.Context, record slog.Record) error {
	panic("implement me")
}

func (h colorHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	panic("implement me")
}

func (h colorHandler) WithGroup(name string) slog.Handler {
	panic("implement me")
}

func TestSlog(t *testing.T) {

}
