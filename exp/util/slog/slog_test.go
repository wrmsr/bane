/*
TODO:
  - stack stability
  - buffer pool?
  - 'old' bridge
  - fast pc?
  - Error/Panic/Fatal
  - OrFatal
*/
package slog

import (
	"context"
	"net"
	"os"
	"testing"

	"golang.org/x/exp/slog"
)

//

type colorHandler struct {
	level Level
}

func (h *colorHandler) clone() *colorHandler {
	return &colorHandler{
		level: h.level,
	}
}

var _ Handler = &colorHandler{}

func (h *colorHandler) Enabled(ctx context.Context, level Level) bool {
	return level >= h.level
}

func (h *colorHandler) Handle(ctx context.Context, record Record) error {
	panic("implement me")
}

func (h *colorHandler) WithAttrs(attrs []Attr) Handler {
	panic("implement me")
}

func (h *colorHandler) WithGroup(name string) Handler {
	panic("implement me")
}

//

func TestSlog(t *testing.T) {
	slog.SetDefault(slog.New(
		//slog.NewTextHandler(os.Stderr),
		//slog.NewJSONHandler(os.Stderr),
		slog.HandlerOptions{
			AddSource: true,
		}.NewJSONHandler(os.Stderr),
	))

	slog.Info("hello", "name", "Al")
	slog.Default().Info("hello", "name", "Al")
	slog.Error("oops", "err", net.ErrClosed, "status", 500)

	slog.LogAttrs(nil, slog.LevelError, "oops", slog.Any("err", net.ErrClosed), slog.Int("status", 500))

	l := slogLogger{slog.Default()}
	l.Info("hi")

	var dl DefaultLogger
	dl.Info("hi2")
}
