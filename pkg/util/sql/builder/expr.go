package builder

//

type Expr interface {
	Node
	isExpr()
}

type expr struct {
	node
}

func (e *expr) isExpr() {}

//

type Literal struct {
	expr
	String string
}

//

type Identifier struct {
	expr
	Name string
}
