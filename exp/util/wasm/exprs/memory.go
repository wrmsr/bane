package exprs

//

type SetLocal struct {
	expr
	N string
	V Expr
}

var _ Expr = SetLocal{}
