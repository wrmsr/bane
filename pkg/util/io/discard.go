package io

import (
	"fmt"
	"io"
)

//

func Discard(n int, err error) {}

func DiscardLen(n int, err error) error { return err }

//

func WriteErr(w io.Writer, p []byte) error {
	return DiscardLen(w.Write(p))
}

func WriteDiscard(w io.Writer, p []byte) {
	Discard(w.Write(p))
}

//

func WriteString(w io.Writer, s string) (int, error) {
	return w.Write([]byte(s))
}

func WriteStringErr(w io.Writer, s string) error {
	return DiscardLen(w.Write([]byte(s)))
}

func WriteStringDiscard(w io.Writer, s string) {
	Discard(WriteString(w, s))
}

//

func FprintfErr(w io.Writer, format string, a ...any) error {
	return DiscardLen(fmt.Fprintf(w, format, a...))
}

func FprintfDiscard(w io.Writer, format string, a ...any) {
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
