package exprs

import (
	"fmt"

	wr "github.com/wrmsr/bane/exp/util/wasm/rendering"
)

type Expr interface {
	isExpr()

	wr.Renderer
	fmt.Stringer
}

type expr struct{}

func (e expr) isExpr() {}
