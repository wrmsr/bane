package context

import (
	"context"
	"time"
)

func WithTimeout(ctx *context.Context, to time.Duration) func() {
	if to == time.Duration(0) {
		return func() {}
	}
	tctx, cancel := context.WithTimeout(*ctx, to)
	*ctx = tctx
	return cancel
}
