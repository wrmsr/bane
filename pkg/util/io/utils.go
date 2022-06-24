package io

import (
	"fmt"
	"io"
)

//

type StringWriter interface {
	WriteString(string) (int, error)
}

//

func Discard(n int, err error) {}

//

func WriteString(w io.Writer, s string) (int, error) {
	return w.Write([]byte(s))
}

func WriteStringX(w io.Writer, s string) {
	Discard(WriteString(w, s))
}

func FprintfX(w io.Writer, format string, a ...any) {
	Discard(fmt.Fprintf(w, format, a...))
}

//

type DiscardStringWriter interface {
	WriteString(string)
}

type discardStringWriter struct {
	w StringWriter
}

func (w discardStringWriter) WriteString(s string) {
	Discard(w.w.WriteString(s))
}

func NewDiscardStringWriter(w StringWriter) DiscardStringWriter {
	return discardStringWriter{w}
}
