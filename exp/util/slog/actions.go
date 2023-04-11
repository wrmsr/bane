package slog

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
