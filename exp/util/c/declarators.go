package c

//

type Declarator interface {
	Node
	isDeclarator()
}

type declarator struct{}

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
