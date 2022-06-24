package builder

import "github.com/wrmsr/bane/pkg/util/check"

//

type Relation interface {
	Node
	isRelation()
}

type relation struct {
	node
}

func (r *relation) isRelation() {}

//

type Table struct {
	relation
	Name *Identifier
}

func NewTable(name *Identifier) *Table {
	return &Table{
		Name: check.NotZero(name),
	}
}
