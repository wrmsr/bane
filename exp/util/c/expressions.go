package c

//

type Expression interface {
	Node
	isExpression()
}

type expression struct {
	node
}

func (e expression) isExpression() {}

//

type Call struct {
	Fn   Expression
	Args []Expression

	expression
}
