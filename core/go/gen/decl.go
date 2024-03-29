package gen

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

func DeclOf(o any) Decl {
	if o, ok := o.(Decl); ok {
		return o
	}

	return StmtDeclOf(StmtOf(o))
}

//

type Decls []Decl

func (s *Decls) Append(os ...any) {
	for _, o := range os {
		if os, ok := o.(Decls); ok {
			*s = append(*s, os...)
		} else {
			*s = append(*s, DeclOf(o))
		}
	}
}

//

type Func struct {
	Receiver *Param
	Name     *Ident
	Params   []Param
	Type     Type
	Body     *Block

	decl
}

//

type Import struct {
	Spec  string
	Alias *Ident

	decl
}

//

type Imports struct {
	Imports []Import

	decl
}

//

type Package struct {
	Name Ident

	decl
}

//

type Param struct {
	Name *Ident
	Type Type

	node
}

//

type StmtDecl struct {
	Stmt Stmt

	decl
}

func StmtDeclOf(stmt Stmt) StmtDecl {
	return StmtDecl{
		Stmt: stmt,
	}
}

//

type Struct struct {
	Name   Ident
	Fields []StructField

	decl
}

//

type StructField struct {
	Name Ident
	Type Type

	node
}
