package wasm

import (
	"fmt"
	"io"

	iou "github.com/wrmsr/bane/pkg/util/io"
)

//

type Expr interface {
	isExpr()

	Renderer
	fmt.Stringer
}

type expr struct{}

func (e expr) isExpr() {}

//

type Block struct {
	expr
	id string
	s  []Expr
}

var _ Expr = Block{}

//

type Const struct {
	expr
	s  string
	ty Type
}

var _ Expr = Const{}

//

type Nop struct {
	expr
}

var _ Expr = Nop{}

//

type SetLocal struct {
	expr
	n string
	v Expr
}

var _ Expr = SetLocal{}

//

func (e Block) Render(w io.Writer) {
	d := iou.Dw(w).
		String("(block ")
	if e.id != "" {
		d.
			String(e.id).
			Byte(' ')
	}
	for i, c := range e.s {
		if i > 0 {
			d.Byte(' ')
		}
		c.Render(w)
	}
	d.Byte(')')
}

func (e Const) Render(w io.Writer) {
	d := iou.Dw(w).
		String("(const ")
	e.ty.Render(w)
	d.
		Byte(' ').
		String(e.s).
		Byte(')')
}

func (e Nop) Render(w io.Writer) {
	iou.Dw(w).String("(nop)")
}

func (e SetLocal) Render(w io.Writer) {
	d := iou.Dw(w).
		String("(set_local ").
		String(e.n).
		Byte(' ')
	e.v.Render(w)
	d.Byte(')')
}

func (e Block) String() string    { return RenderString(e) }
func (e Const) String() string    { return RenderString(e) }
func (e Nop) String() string      { return RenderString(e) }
func (e SetLocal) String() string { return RenderString(e) }
