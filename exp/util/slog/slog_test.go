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
	"fmt"
	"net"
	"os"
	"strconv"
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

	var l Logger
	l.Info("hi")
}

func TestAllocs(t *testing.T) {
	slog.SetDefault(slog.New(NullHandler{}))

	var l Logger
	n := testing.AllocsPerRun(1_000, func() {
		var a []any
		a = append(a, "foo", 420)
		for i := 0; i < 3; i++ {
			a = append(a, strconv.Itoa(i), i)
		}
		l.Info("hi", a...)
	})

	fmt.Println(n)
}
