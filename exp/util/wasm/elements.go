package wasm

import (
	"io"
	"strings"
)

//

type Element interface {
	isElement()

	Render(w io.Writer)
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

func RenderElement(p Element) string {
	var bs strings.Builder
	p.Render(&bs)
	return bs.String()
}

//

func (p Atom) Render(w io.Writer) {
	_, _ = io.WriteString(w, p.s)
}

func (p Quote) Render(w io.Writer) {
	_, _ = w.Write([]byte{'"'})
	_, _ = io.WriteString(w, p.s)
	_, _ = w.Write([]byte{'"'})
}

func (p List) Render(w io.Writer) {
	_, _ = w.Write([]byte{'('})
	for i, c := range p.ps {
		if i > 0 {
			_, _ = w.Write([]byte{' '})
		}
		c.Render(w)
	}
	_, _ = w.Write([]byte{')'})
}
