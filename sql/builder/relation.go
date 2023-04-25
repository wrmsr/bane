package builder

import (
	"github.com/wrmsr/bane/pkg/util/check"
	bt "github.com/wrmsr/bane/pkg/util/types"
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
	Alias bt.Optional[Ident]
}

func NewTableAs(name Ident, alias bt.Optional[Ident]) Table {
	return Table{
		Name:  check.NotZero(name),
		Alias: alias,
	}
}

func NewTable(name Ident) Table {
	return NewTableAs(name, bt.None[Ident]())
}
