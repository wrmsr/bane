package c

//

type ExpressionStatement struct {
	statement
	Expression Expression
}

//

type Expression interface {
	Node
	isExpression()
}

type expression struct {
	node
}

func (e expression) isExpression() {}
