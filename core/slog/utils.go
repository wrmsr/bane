package slog

import (
	"runtime"
)

func RecordFrame(r Record) runtime.Frame {
	fs := runtime.CallersFrames([]uintptr{r.PC})
	f, _ := fs.Next()
	return f
}

func ParseLevel(s string) (Level, error) {
	var l Level
	err := l.UnmarshalText([]byte(s))
	return l, err
}
