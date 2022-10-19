package parsing

import (
	"fmt"
	"io"

	wr "github.com/wrmsr/bane/exp/util/wasm/rendering"
	iou "github.com/wrmsr/bane/pkg/util/io"
)

//

type Element interface {
	isElement()

	wr.Renderer
	fmt.Stringer
}

type element struct{}

func (p element) isElement() {}

type Atom struct {
	element
	S string
}

var _ Element = Atom{}

type Quote struct {
	element
	S string
}

var _ Element = Quote{}

type List struct {
	element
	Ps []Element
}

var _ Element = List{}

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

func (e Atom) String() string  { return wr.RenderString(e) }
func (e Quote) String() string { return wr.RenderString(e) }
func (e List) String() string  { return wr.RenderString(e) }
