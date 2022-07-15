package telnet

import (
	"io"

	iou "github.com/wrmsr/bane/pkg/util/io"
)

type Writer struct {
	w io.Writer
}

var _ io.Writer = &Writer{}

var _iacSlice = []byte{IAC}

func (w *Writer) Write(p []byte) (int, error) {
	n := 0

	for i := 0; i < len(p); i++ {
		if p[i] == IAC {
			if err := iou.WriteAdd(w.w, p[:i+1], &n); err != nil {
				return n, err
			}

			if err := iou.WriteAdd(w.w, _iacSlice, &n); err != nil {
				return n, err
			}

			p = p[i+1:]
			i = -1
		}
	}

	if err := iou.WriteAdd(w.w, p, &n); err != nil {
		return n, err
	}

	return n, nil
}
