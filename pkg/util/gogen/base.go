package gogen

//

type Node interface {
	isNode()
}

type node struct {
}

func (n node) isNode() {}

//

type Type struct {
	node
	Name Ident
}

func NewType(name Ident) Type {
	return Type{
		Name: name,
	}
}
