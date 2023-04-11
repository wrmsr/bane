package slog

import (
	"runtime"
)

func RecordFrame(r Record) runtime.Frame {
	fs := runtime.CallersFrames([]uintptr{r.PC})
	f, _ := fs.Next()
	return f
}
