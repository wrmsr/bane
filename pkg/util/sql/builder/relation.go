package builder

import (
	"github.com/wrmsr/bane/pkg/util/check"
	opt "github.com/wrmsr/bane/pkg/util/optional"
)

//

type Relation interface {
	Node
	isRelation()
}

type relation struct {
	node
}

func (r relation) isRelation() {}

//

type Table struct {
	relation
	Name  Ident
	Alias opt.Optional[Ident]
}

func NewTableAs(name Ident, alias opt.Optional[Ident]) Table {
	return Table{
		Name:  check.NotZero(name),
		Alias: alias,
	}
}

func NewTable(name Ident) Table {
	return NewTableAs(name, opt.None[Ident]())
}
