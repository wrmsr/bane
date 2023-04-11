package slog

import (
	"context"
)

//

type NullHandler struct{}

var _ Handler = NullHandler{}

func (h NullHandler) Enabled(ctx context.Context, level Level) bool   { return false }
func (h NullHandler) Handle(ctx context.Context, record Record) error { return nil }

func (h NullHandler) WithAttrs(attrs []Attr) Handler { return h }
func (h NullHandler) WithGroup(name string) Handler  { return h }

//

type Writer = func(b []byte) (int, error)

type Formatter interface {
	WithAttrs(as []Attr) Formatter
	WithGroup(name string) Formatter

	Format(r Record, w Writer) error
}

type FormatterHandler struct {
	f  Formatter
	w  Writer
	lv Level
}

func NewFormatterHandler(
	f Formatter,
	w Writer,
	lv Level,
) FormatterHandler {
	return FormatterHandler{
		f:  f,
		w:  w,
		lv: lv,
	}
}

func (h *FormatterHandler) clone() *FormatterHandler {
	c := *h
	return &c
}

var _ Handler = &FormatterHandler{}

func (h *FormatterHandler) Enabled(ctx context.Context, level Level) bool {
	return level >= h.lv
}

func (h *FormatterHandler) Handle(ctx context.Context, record Record) error {
	return h.f.Format(record, h.w)
}

func (h *FormatterHandler) WithAttrs(attrs []Attr) Handler {
	c := h.clone()
	c.f = h.f.WithAttrs(attrs)
	return c
}

func (h *FormatterHandler) WithGroup(name string) Handler {
	c := h.clone()
	c.f = h.f.WithGroup(name)
	return c
}
