package log

import (
	"fmt"
	"strings"
	"time"

	ju "github.com/wrmsr/bane/pkg/util/json"
	tiu "github.com/wrmsr/bane/pkg/util/time"
)

//

const defaultTimeFormat = tiu.RFC3339NZ

//

type Formatter interface {
	FormatLine(line Line) (string, error)
}

//

type TextFormatter struct{}

var _ Formatter = TextFormatter{}

func (f TextFormatter) FormatLine(line Line) (string, error) {
	var b strings.Builder
	b.WriteString(line.Time.Format(defaultTimeFormat))
	b.WriteString(" ")
	b.WriteString(line.Level.String())
	b.WriteString(" ")
	b.WriteString(line.Message)
	for _, a := range line.Args {
		b.WriteString(" ")
		b.WriteString(a.Name)
		b.WriteString("=")
		b.WriteString(fmt.Sprintf("%v", a.Value))
	}
	return b.String(), nil
}

//

type JsonFormatter struct{}

var _ Formatter = JsonFormatter{}

type jsonLine struct {
	Time    time.Time      `json:"time"`
	Level   Level          `json:"level"`
	Message string         `json:"message"`
	Args    map[string]any `json:"args,omitempty"`
}

func (f JsonFormatter) FormatLine(line Line) (string, error) {
	jl := jsonLine{
		Time:    line.Time,
		Level:   line.Level,
		Message: line.Message,
	}

	// FIXME: ordmap
	if len(line.Args) > 0 {
		m := make(map[string]any, len(line.Args))
		for _, a := range line.Args {
			m[a.Name] = a.Value
		}
		jl.Args = m
	}

	return ju.MarshalString(jl)
}
