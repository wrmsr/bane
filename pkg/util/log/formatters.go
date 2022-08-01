package log

import (
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	iou "github.com/wrmsr/bane/pkg/util/io"
	ju "github.com/wrmsr/bane/pkg/util/json"
	rtu "github.com/wrmsr/bane/pkg/util/runtime"
	stru "github.com/wrmsr/bane/pkg/util/strings"
	tiu "github.com/wrmsr/bane/pkg/util/time"
)

//

const defaultTimeFormat = tiu.RFC3339NZ

//

type Formatter interface {
	FormatLine(line Line) (string, error)
}

//

func WriteStackFrame(b io.StringWriter, sf rtu.StackFrame) (n int) {
	_ = iou.WriteStringAdd(b, " ", &n)
	_, fn, ok := stru.LastCut(sf.File, "/")
	if ok {
		_ = iou.WriteStringAdd(b, fn, &n)
		_ = iou.WriteStringAdd(b, ":", &n)
		_ = iou.WriteStringAdd(b, strconv.Itoa(sf.Line), &n)
	}
	_ = iou.WriteStringAdd(b, ":", &n)
	_ = iou.WriteStringAdd(b, sf.ParsedName().Obj, &n)
	return
}

func FormatStackFrame(sf rtu.StackFrame) string {
	var b strings.Builder
	WriteStackFrame(&b, sf)
	return b.String()
}

//

type TextFormatter struct {
	prefix string
	suffix string

	noTime bool

	callerWidth int
}

func NewTextFormatter() TextFormatter {
	return TextFormatter{
		suffix: "\n",

		callerWidth: 36,
	}
}

var _ Formatter = TextFormatter{}

func (f TextFormatter) FormatLine(line Line) (string, error) {
	var b strings.Builder

	if f.prefix != "" {
		b.WriteString(f.prefix)
	}

	if !f.noTime {
		b.WriteString(line.Time.Format(defaultTimeFormat))
		b.WriteString(" ")
	}

	b.WriteString(line.Level.String())

	sf := line.GetStackFrame(1)
	sfl := WriteStackFrame(&b, sf)
	for i := sfl; i < f.callerWidth; i++ {
		b.WriteRune(' ')
	}

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

	if f.suffix != "" {
		b.WriteString(f.suffix)
	}

	return b.String(), nil
}

//

type JsonFormatter struct{}

var _ Formatter = JsonFormatter{}

type jsonLine struct {
	Time  time.Time `json:"time"`
	Level Level     `json:"level"`
	At    string    `json:"at"`

	Message string `json:"message,omitempty"`

	Args map[string]any `json:"args,omitempty"`
}

func (f JsonFormatter) FormatLine(line Line) (string, error) {
	jl := jsonLine{
		Time:  line.Time,
		Level: line.Level,

		Message: line.Message,
	}

	sf := line.GetStackFrame(1)
	jl.At = FormatStackFrame(sf)

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
