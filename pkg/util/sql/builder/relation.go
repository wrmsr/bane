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
	Name  Identifier
	Alias opt.Optional[Identifier]
}

func NewTableAs(name Identifier, alias opt.Optional[Identifier]) Table {
	return Table{
		Name:  check.NotZero(name),
		Alias: alias,
	}
}

func NewTable(name Identifier) Table {
	return NewTableAs(name, opt.None[Identifier]())
}
