package gogen

//

type Decl interface {
	Node
	isDecl()
}

type decl struct {
	node
}

func (e decl) isDecl() {}

//

type Func struct {
	decl
	Name Ident
	Body Block
}
