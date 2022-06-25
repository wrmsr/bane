package gogen

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

//

type SelectItem struct {
	node
	Expr  Expr
	Label opt.Optional[Identifier]
}

func NewSelectItemAs(expr Expr, label opt.Optional[Identifier]) SelectItem {
	return SelectItem{
		Expr:  check.NotNil(expr).(Expr),
		Label: label,
	}
}

func NewSelectItem(expr Expr) SelectItem {
	return NewSelectItemAs(expr, opt.None[Identifier]())
}
