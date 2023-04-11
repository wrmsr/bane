package slog

import (
	"context"

	"golang.org/x/exp/slog"
)

//

type NullHandler struct{}

var _ Handler = NullHandler{}

func (h NullHandler) Enabled(ctx context.Context, level slog.Level) bool   { return false }
func (h NullHandler) Handle(ctx context.Context, record slog.Record) error { return nil }

func (h NullHandler) WithAttrs(attrs []slog.Attr) slog.Handler { return h }
func (h NullHandler) WithGroup(name string) slog.Handler       { return h }
