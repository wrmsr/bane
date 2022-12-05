package exprs

import "github.com/wrmsr/bane/exp/util/wasm/text"

//

type Expr interface {
	isExpr()

	Text() text.Element
}

type expr struct{}

func (e expr) isExpr() {}
