package gogen

//

type Stmt interface {
	Node
	isStmt()
}

type stmt struct {
	node
}

func (s stmt) isStmt() {}

//

type Block struct {
	stmt
	Body []Stmt
}

//

type If struct {
	stmt
	Then Block
	Else Block
}
