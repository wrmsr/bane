package slog

import (
	"context"
	"os"
	"runtime"
	"time"

	"golang.org/x/exp/slog"
)

var IgnorePC = false

func log(
	ctx context.Context,
	h Handler,
	level Level,
	action Action,
	pco int,
	msg string,
	prepRec func(*Record),
) {
	if !h.Enabled(ctx, level) {
		return
	}

	var pc uintptr
	if !IgnorePC {
		var pcs [1]uintptr
		runtime.Callers(3+pco, pcs[:])
		pc = pcs[0]
	}

	r := slog.NewRecord(time.Now(), level, msg, pc)
	if prepRec != nil {
		prepRec(&r)
	}
	_ = h.Handle(ctx, r)

	switch action {
	case ActionPanic:
		panic(r)
	case ActionFatal:
		os.Exit(1)
	}
}
