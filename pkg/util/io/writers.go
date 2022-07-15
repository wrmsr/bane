package io

import "io"

func WriteAdd(w io.Writer, p []byte, n *int) error {
	c, err := w.Write(p)
	*n += c
	return err
}
