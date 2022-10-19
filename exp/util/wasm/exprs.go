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
	id string
	s  []Expr
}

//

type SetLocal struct {
	expr
	n string
	v Expr
}

//

type Const struct {
	expr
	s  string
	ty string
}
