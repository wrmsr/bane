package builder

import (
	"github.com/wrmsr/bane/core/check"
	bt "github.com/wrmsr/bane/core/types"
)

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

type Select struct {
	stmt
	Items []SelectItem
	From  Relation
	Where Expr
}

//

type SelectItem struct {
	node
	Expr  Expr
	Label bt.Optional[Ident]
}

func NewSelectItemAs(expr Expr, label bt.Optional[Ident]) SelectItem {
	return SelectItem{
		Expr:  check.NotNil(expr).(Expr),
		Label: label,
	}
}

func NewSelectItem(expr Expr) SelectItem {
	return NewSelectItemAs(expr, bt.None[Ident]())
}
