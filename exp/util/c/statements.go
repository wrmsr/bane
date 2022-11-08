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

//

type BlockItem = any // Declaration | Statement

type CompoundStatement struct {
	statement
	S []BlockItem
}

//

type JumpStatement interface {
	Statement
	isJumpStatement()
}

type jumpStatement struct {
	statement
}

func (s jumpStatement) isJumpStatement() {}

//

type ReturnStatement struct {
	jumpStatement
	V Expression
}
