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

type DeclarationSpecifier interface {
	isDeclarationSpecifier()
}

func (t typeSpecifier) isDeclarationSpecifier()         {}
func (s StorageClassSpecifier) isDeclarationSpecifier() {}
func (q TypeQualifier) isDeclarationSpecifier()         {}

// func (s FunctionSpecifier) isDeclarationSpecifier()  {}
// func (s AlignmentSpecifier) isDeclarationSpecifier() {}

type DeclarationSpecifiers struct {
	node
	S []DeclarationSpecifier
}

type DeclaratorsDeclaration struct {
	declaration
	S  DeclarationSpecifiers
	Ds []Declarator
}

//

type ParameterDeclaration struct {
	declaration
	S DeclarationSpecifiers
	D OptDeclarator
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

//

type StructDeclaration struct {
	declaration
	Sqs []SpecifierQualifier
	Sds []StructDeclarator
}
