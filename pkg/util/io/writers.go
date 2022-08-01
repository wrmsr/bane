package io

import "io"

func WriteAdd(w io.Writer, p []byte, n *int) error {
	c, err := w.Write(p)
	*n += c
	return err
}

func WriteStringAdd(w io.StringWriter, s string, n *int) error {
	c, err := w.WriteString(s)
	*n += c
	return err
}
