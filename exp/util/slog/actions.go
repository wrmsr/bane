package slog

import "fmt"

//

type Action int8

const (
	ActionNone Action = iota
	ActionPanic
	ActionFatal
)

func (a Action) String() string {
	switch a {
	case ActionNone:
		return "none"
	case ActionPanic:
		return "panic"
	case ActionFatal:
		return "fatal"
	}
	panic(a)
}

//

const ErrorMsg = "error"

//

type LogPanic struct {
	*Record
}

var _ error = LogPanic{}

func (l LogPanic) Error() string {
	return fmt.Sprintf("%+v", l.Record)
}
