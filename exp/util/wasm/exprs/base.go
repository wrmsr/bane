package exprs

import (
	"github.com/wrmsr/bane/exp/util/wasm/text"
)

//

type Expr interface {
	isExpr()

	Text() text.Element
}

type expr struct{}

func (e expr) isExpr() {}

////
//
//type VarExpr interface {
//	Expr
//
//	isVarExpr()
//
//	Var() int
//}
//
//type varExpr struct {
//	expr
//	V int
//}
//
//func (e varExpr) isVarExpr() {}
//
//func (e varExpr) Var() int { return e.V }
