package c

//

type Declarator interface {
	Node
	isDeclarator()
}

type declarator struct {
	node
}

func (d declarator) isDeclarator() {}

//

type IdentifierDeclarator struct {
	declarator
	I Identifier
}

//

type ParameterListDeclarator struct {
	declarator
	D  Declarator
	Ps ParameterList
}
