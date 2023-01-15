package c

import bt "github.com/wrmsr/bane/pkg/util/types"

//

type OptDeclarator = bt.Optional[Declarator]

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
	D OptDeclarator
}

//

type ParameterListDeclarator struct {
	declarator
	D  Declarator
	Ps ParameterList
}

//

type StructDeclarator struct {
	declarator
	D Declarator
	// C Expression
}
