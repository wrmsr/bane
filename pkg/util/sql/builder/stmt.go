package builder

import (
	"github.com/wrmsr/bane/pkg/util/check"
	opt "github.com/wrmsr/bane/pkg/util/optional"
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
	Label opt.Optional[Ident]
}

func NewSelectItemAs(expr Expr, label opt.Optional[Ident]) SelectItem {
	return SelectItem{
		Expr:  check.NotNil(expr).(Expr),
		Label: label,
	}
}

func NewSelectItem(expr Expr) SelectItem {
	return NewSelectItemAs(expr, opt.None[Ident]())
}
