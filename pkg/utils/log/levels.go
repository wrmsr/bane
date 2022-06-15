package log

import "fmt"

type Level int8

const (
	DebugLevel Level = iota - 1
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case PanicLevel:
		return "panic"
	case FatalLevel:
		return "fatal"
	}
	panic("unreachable")
}

func ParseLevel(s string) (Level, error) {
	switch s {
	case "debug":
		return DebugLevel, nil
	case "info":
		return InfoLevel, nil
	case "warn":
		return WarnLevel, nil
	case "error":
		return ErrorLevel, nil
	case "panic":
		return PanicLevel, nil
	case "fatal":
		return FatalLevel, nil
	}
	return InfoLevel, fmt.Errorf("unknown level: %s", s)
}
