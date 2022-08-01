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
	if line.Message != "" {
		b.WriteString(" ")
		b.WriteString(line.Message)
	}
	for _, a := range line.Args {
		n := a.ArgName()
		if n == "" {
			continue
		}
		b.WriteString(" ")
		b.WriteString(n)
		b.WriteString("=")
		b.WriteString(fmt.Sprintf("%v", a.ArgValue()))
	}
	return b.String(), nil
}

//

type JsonFormatter struct{}

var _ Formatter = JsonFormatter{}

type jsonLine struct {
	Time    time.Time      `json:"time"`
	Level   Level          `json:"level"`
	Message string         `json:"message,omitempty"`
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
			n := a.ArgName()
			if n == "" {
				continue
			}
			m[n] = a.ArgValue()
		}
		jl.Args = m
	}

	return ju.MarshalString(jl)
}
