package exprs

import (
	"fmt"
	"io"

	wr "github.com/wrmsr/bane/exp/util/wasm/rendering"
	wt "github.com/wrmsr/bane/exp/util/wasm/types"
	iou "github.com/wrmsr/bane/pkg/util/io"
)

//

type Expr interface {
	isExpr()

	wr.Renderer
	fmt.Stringer
}

type expr struct{}

func (e expr) isExpr() {}

//

type Block struct {
	expr
	Id string
	S  []Expr
}

var _ Expr = Block{}

//

type Const struct {
	expr
	S  string
	Ty wt.Type
}

var _ Expr = Const{}

//

type Nop struct {
	expr
}

var _ Expr = Nop{}

//

type Select struct {
	expr
	C Expr
	T Expr
	F Expr
}

var _ Expr = Select{}

//

type SetLocal struct {
	expr
	N string
	V Expr
}

var _ Expr = SetLocal{}

//

func (e Block) Render(w io.Writer) {
	d := iou.Dw(w).
		String("(block ")
	if e.Id != "" {
		d.
			String(e.Id).
			Byte(' ')
	}
	for i, c := range e.S {
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
	e.Ty.Render(w)
	d.
		Byte(' ').
		String(e.S).
		Byte(')')
}

func (e Nop) Render(w io.Writer) {
	iou.Dw(w).String("(nop)")
}

func (e Select) Render(w io.Writer) {
	d := iou.Dw(w).
		String("(select ")
	e.C.Render(w)
	d.Byte(' ')
	e.T.Render(w)
	if e.F != nil {
		d.Byte(' ')
		e.F.Render(w)
	}
	d.Byte(')')
}

func (e SetLocal) Render(w io.Writer) {
	d := iou.Dw(w).
		String("(set_local ").
		String(e.N).
		Byte(' ')
	e.V.Render(w)
	d.Byte(')')
}

func (e Block) String() string    { return wr.RenderString(e) }
func (e Const) String() string    { return wr.RenderString(e) }
func (e Nop) String() string      { return wr.RenderString(e) }
func (e Select) String() string   { return wr.RenderString(e) }
func (e SetLocal) String() string { return wr.RenderString(e) }
