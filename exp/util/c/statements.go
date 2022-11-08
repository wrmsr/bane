package c

//

type Statement interface {
	Node
	isStatement()
}

type statement struct {
	node
}

func (s statement) isStatement() {}

//

type ExpressionStatement struct {
	statement
	Expression Expression
}
