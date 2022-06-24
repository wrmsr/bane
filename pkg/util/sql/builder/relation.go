package builder

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
	Identifier *Identifier
}
