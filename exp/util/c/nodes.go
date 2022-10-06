package c

//

type Node interface {
	isNode()
}

type node struct{}

func (n node) isNode() {}

//

type CompilationUnit struct {
	node
}

//

type FunctionDef struct {
	node
}

//

type Decl interface {
	Node
	isDecl()
}

type decl struct {
	node
}

func (d decl) isDecl() {}

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

type ExprStmt struct {
	Expr Expr

	stmt
}

//

type Expr interface {
	Node
	isExpr()
}

type expr struct {
	node
}

func (e expr) isExpr() {}
