package io

import (
	"strings"

	stru "github.com/wrmsr/bane/pkg/util/strings"
)

type IndentWriter struct {
	w DiscardStringWriter
	s string
	l int
	b bool
}

func NewIndentWriter(w DiscardStringWriter, indent string) *IndentWriter {
	return &IndentWriter{
		w: w,
		s: indent,
	}
}

var _ DiscardStringWriter = &IndentWriter{}

func (w *IndentWriter) WriteString(s string) {
	indent := strings.Repeat(w.s, w.l)
	if strings.ContainsRune(s, '\n') {
		s = stru.Dedent(s)
	}

	for s != "" {
		if !w.b {
			w.w.WriteString(indent)
			w.b = true
		}

		n := strings.IndexRune(s, '\n')
		if n < 0 {
			w.w.WriteString(s)
			break
		}

		w.w.WriteString(s[:n+1])
		w.b = false
		if n == len(s)-1 {
			break
		}

		s = s[n+1:]
	}
}

func (w *IndentWriter) Indent() {
	w.l++
}

func (w *IndentWriter) Dedent() {
	if w.l < 1 {
		panic("negative indent")
	}
	w.l--
}
