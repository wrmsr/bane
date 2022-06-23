package io

import (
	"fmt"
	"io"
)

func Discard(n int, err error) {}

func WriteString(w io.Writer, s string) (int, error) {
	return w.Write([]byte(s))
}

func WriteStringX(w io.Writer, s string) {
	Discard(WriteString(w, s))
}

func FprintfX(w io.Writer, format string, a ...any) {
	Discard(fmt.Fprintf(w, format, a...))
}
