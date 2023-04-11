package slog

import (
	"context"

	"golang.org/x/exp/slog"
)

//

func Default() Logger {
	return NewLogger(DefaultHandler())
}

func DefaultCtx(ctx context.Context) Logger {
	return NewLogger(DefaultHandler()).WithContext(ctx)
}

func DefaultHandler() Handler {
	return slog.Default().Handler()
}
