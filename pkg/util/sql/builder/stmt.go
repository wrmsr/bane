package builder

//

type Stmt interface {
	Node
	isStmt()
}

type stmt struct {
	node
}

func (s *stmt) isStmt() {}

//

type Select struct {
	stmt
	Items []*SelectItem
	From  Relation
	Where Expr
}

//

type SelectItem struct {
	node
	Expr       Expr
	Identifier *Identifier
}
