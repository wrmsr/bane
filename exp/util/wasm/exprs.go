package wasm

//

type Expr interface {
	isExpr()
}

type expr struct{}

func (e expr) isExpr() {}

//

type Block struct {
	expr
}

//

type SetLocal struct {
	expr
}

//

type Const struct {
	expr
}
