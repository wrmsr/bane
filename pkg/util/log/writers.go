package log

import (
	"io"
)

//

type Writer interface {
	Write(s string) error
}

//

type IoWriter struct {
	w io.StringWriter
}

func NewIoWriter(w io.StringWriter) IoWriter {
	return IoWriter{w: w}
}

var _ Writer = IoWriter{}

func (w IoWriter) Write(s string) error {
	_, _ = w.w.WriteString(s)
	return nil
}
