package types

import (
	"fmt"
	"io"

	wr "github.com/wrmsr/bane/exp/util/wasm/rendering"
	iou "github.com/wrmsr/bane/pkg/util/io"
)

//

type Type interface {
	isType()

	wr.Renderer
	fmt.Stringer
}

type type_ struct{}

func (t type_) isType() {}

//

type PrimType struct {
	type_
	P Prim
}

var _ Type = PrimType{}

func (t PrimType) Render(w io.Writer) { iou.Dw(w).String("none") }

func (t PrimType) String() string { return wr.RenderString(t) }
