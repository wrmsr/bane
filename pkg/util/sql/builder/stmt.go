package builder

import "github.com/wrmsr/bane/pkg/util/check"

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
	Expr  Expr
	Label *Identifier
}

func NewSelectItem(expr Expr, label *Identifier) *SelectItem {
	return &SelectItem{
		Expr:  check.NotNil(expr).(Expr),
		Label: label,
	}
}
