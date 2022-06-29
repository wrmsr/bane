package log

import (
	"fmt"
)

//

type Field int8

const (
	InvalidField Field = iota
	TimeField
	LevelField
	MessageField
	ArgsField
	CallerField
)

func (f Field) String() string {
	switch f {
	case InvalidField:
		return "invalid"
	case TimeField:
		return "time"
	case LevelField:
		return "level"
	case MessageField:
		return "message"
	case ArgsField:
		return "args"
	case CallerField:
		return "caller"
	}
	panic("unreachable")
}

func ParseField(s string) (Field, error) {
	switch s {
	case "time":
		return TimeField, nil
	case "level":
		return LevelField, nil
	case "message":
		return MessageField, nil
	case "args":
		return ArgsField, nil
	case "caller":
		return CallerField, nil
	}
	return InvalidField, fmt.Errorf("unknown field: %s", s)
}

//

func ExtractField(line Line, field Field) (any, error) {
	switch field {
	case TimeField:
		return line.Time.Format(defaultTimeFormat), nil
	case LevelField:
		return line.Level, nil
	case MessageField:
		return line.Message, nil
	case ArgsField:
		return line.Args, nil
	}
	return "", fmt.Errorf("unknoown field: %v", field)
}
