package io

import (
	"fmt"
	"io"
	"strings"
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

func WriteByte(w io.Writer, b byte) (int, error) {
	return w.Write([]byte{b})
}

func WriteByteErr(w io.Writer, b byte) error {
	return DiscardLen(w.Write([]byte{b}))
}

func WriteByteDiscard(w io.Writer, b byte) {
	Discard(w.Write([]byte{b}))
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
	w io.StringWriter
}

func (w discardStringWriter) WriteString(s string) {
	Discard(w.w.WriteString(s))
}

func NewDiscardStringWriter(w io.StringWriter) DiscardStringWriter {
	return discardStringWriter{w}
}

//

func InvokeToString(fn func(writer DiscardStringWriter)) string {
	var sb strings.Builder
	fn(NewDiscardStringWriter(&sb))
	return sb.String()
}

//

type DiscardingWriter struct {
	io.Writer
}

func Dw(w io.Writer) DiscardingWriter {
	return DiscardingWriter{w}
}

func (d DiscardingWriter) Bytes(p []byte) DiscardingWriter {
	Discard(d.Write(p))
	return d
}

func (d DiscardingWriter) Byte(b byte) DiscardingWriter {
	Discard(d.Write([]byte{b}))
	return d
}

func (d DiscardingWriter) String(s string) DiscardingWriter {
	Discard(d.Write([]byte(s)))
	return d
}

func (d DiscardingWriter) Fprintf(format string, a ...any) DiscardingWriter {
	Discard(fmt.Fprintf(d.Writer, format, a...))
	return d
}
