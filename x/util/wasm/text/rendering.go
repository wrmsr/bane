package text

import (
	"io"

	iou "github.com/wrmsr/bane/pkg/util/io"
	wr "github.com/wrmsr/bane/x/util/wasm/rendering"
)

//

func (e Atom) Render(w io.Writer) {
	iou.Dw(w).
		String(e.S)
}

func (e Quote) Render(w io.Writer) {
	iou.Dw(w).
		Byte('"').
		String(e.S).
		Byte('"')
}

func (e List) Render(w io.Writer) {
	d := iou.Dw(w).
		Byte('(')
	for i, c := range e.Ps {
		if i > 0 {
			d.Byte(' ')
		}
		c.Render(w)
	}
	d.Byte(')')
}

//

func (e Atom) String() string  { return wr.RenderString(e) }
func (e Quote) String() string { return wr.RenderString(e) }
func (e List) String() string  { return wr.RenderString(e) }
