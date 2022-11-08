package c

//

type Declaration interface {
	Node
	isDeclaration()
}

type declaration struct {
	node
}

func (d declaration) isDeclaration() {}

//

type DeclarationSpecifier struct {
	node
	S any // TypeSpecifier | StorageClassSpecifier | TypeQualifier | FunctionSpecifier | AlignmentSpecifier
}

//

type DeclarationSpecifiers struct {
	node
	S []DeclarationSpecifier
}

//

type ParameterDeclaration struct {
	declaration
	S DeclarationSpecifiers
	D Declarator
}

type ParameterList struct {
	node
	Ps []ParameterDeclaration
	Va bool
}

//

type FunctionDeclaration struct {
	declaration
	Ds DeclarationSpecifiers
	D  Declarator
	// FIXME: Dl DeclarationList
	B CompoundStatement
}
