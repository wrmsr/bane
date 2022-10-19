package wasm

import (
	"fmt"
	"io"

	iou "github.com/wrmsr/bane/pkg/util/io"
)

//

type Element interface {
	isElement()

	Renderer
	fmt.Stringer
}

type element struct{}

func (p element) isElement() {}

type Atom struct {
	element
	s string
}

var _ Element = Atom{}

type Quote struct {
	element
	s string
}

var _ Element = Quote{}

type List struct {
	element
	ps []Element
}

var _ Element = List{}

//

func (e Atom) Render(w io.Writer) {
	iou.Dw(w).
		String(e.s)
}

func (e Quote) Render(w io.Writer) {
	iou.Dw(w).
		Byte('"').
		String(e.s).
		Byte('"')
}

func (e List) Render(w io.Writer) {
	d := iou.Dw(w).
		Byte('(')
	for i, c := range e.ps {
		if i > 0 {
			d.Byte(' ')
		}
		c.Render(w)
	}
	d.Byte(')')
}

func (e Atom) String() string  { return RenderString(e) }
func (e Quote) String() string { return RenderString(e) }
func (e List) String() string  { return RenderString(e) }
