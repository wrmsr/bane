package slog

import (
	"context"

	stdslog "golang.org/x/exp/slog"
)

type stdHandler struct {
	h Handler
}

func AsStdHandler(h Handler) stdslog.Handler {
	return stdHandler{h: h}
}

var _ stdslog.Handler = stdHandler{}

func (h stdHandler) Enabled(ctx context.Context, level stdslog.Level) bool {
	return h.h.Enabled(ctx, level)
}

func (h stdHandler) Handle(ctx context.Context, record stdslog.Record) error {
	return h.h.Handle(ctx, record)
}

func (h stdHandler) WithAttrs(attrs []stdslog.Attr) stdslog.Handler {
	return stdHandler{h.h.WithAttrs(attrs)}
}

func (h stdHandler) WithGroup(name string) stdslog.Handler {
	return stdHandler{h.h.WithGroup(name)}
}
