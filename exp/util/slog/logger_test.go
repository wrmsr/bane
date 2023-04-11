/*
TODO:
  - stack stability
  - buffer pool?
  - 'old' bridge
  - fast pc?
  - Error/Panic/Fatal
  - OrFatal
*/
package slog_test

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"testing"

	stdslog "golang.org/x/exp/slog"

	"github.com/wrmsr/bane/exp/util/slog"
)

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
	stdslog.SetDefault(stdslog.New(
		//slog.NewTextHandler(os.Stderr),
		//slog.NewJSONHandler(os.Stderr),
		stdslog.HandlerOptions{
			AddSource: true,
		}.NewJSONHandler(os.Stderr),
	))

	slog.Info("hello", "name", "Al")
	slog.Default().Info("hello", "name", "Al")
	slog.Error("oops", "err", net.ErrClosed, "status", 500)

	slog.LogAttrs(nil, slog.LevelError, "oops", slog.Any("err", net.ErrClosed), slog.Int("status", 500))

	var l slog.Logger
	l.Info("hi")

	l.LogAttrs(slog.LevelError, "oops", slog.Any("err", net.ErrClosed), slog.Int("status", 500))

	slog.Info("hi")

	//slog.LogAttrs(nil, )

	l.OrError(func() error {
		return errors.New("oops")
	})

	slog.OrError(func() error {
		return errors.New("oops")
	})
}

func TestAllocs(t *testing.T) {
	stdslog.SetDefault(stdslog.New(slog.NullHandler{}))
	var l slog.Logger

	n := testing.AllocsPerRun(1_000, func() {
		slog.Info("hi", "foo", 420)
		l.Info("hi", "foo", 421)
	})

	fmt.Println(n)
}
