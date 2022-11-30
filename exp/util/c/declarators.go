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

type InitDeclarator struct {
	declarator
	D Declarator
	I Expression
}

//

type TypeQualifierList struct {
	node
	S []TypeQualifier
}

type PointerLevel struct {
	Q  *TypeQualifierList
	Ca bool
}

type Pointer struct {
	node
	S []PointerLevel
}

type PointerDeclarator struct {
	declarator
	P Pointer
	D Declarator
}

//

type ParameterListDeclarator struct {
	declarator
	D  Declarator
	Ps ParameterList
}
