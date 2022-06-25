package gogen

import (
	"github.com/wrmsr/bane/pkg/util/check"
	opt "github.com/wrmsr/bane/pkg/util/optional"
)

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

type Import struct {
	node
	Spec  string
	Alias opt.Optional[Ident]
}

func NewImportAs(spec string, alias opt.Optional[Ident]) Import {
	return Import{
		Spec:  check.NotZero(spec),
		Alias: alias,
	}
}

func NewImport(spec string) Import {
	return NewImportAs(spec, opt.None[Ident]())
}

//

type Imports struct {
	node
	Imports []Import
}

func NewImports(imports []Import) Imports {
	return Imports{
		Imports: imports,
	}
}

//

type Param struct {
	node
	Name Ident
	Type Type
}

func NewParam(name Ident, type_ Type) Param {
	return Param{
		Name: name,
		Type: type_,
	}
}

//

type Func struct {
	decl
	Name   Ident
	Params []Param
	Type   opt.Optional[Type]
	Body   Block
}

func NewFunc(name Ident, params []Param, type_ opt.Optional[Type], body Block) Func {
	return Func{
		Name:   name,
		Params: params,
		Type:   type_,
		Body:   body,
	}
}

//

type StmtDecl struct {
	decl
	Stmt Stmt
}

func NewStmtDecl(stmt Stmt) StmtDecl {
	return StmtDecl{
		Stmt: stmt,
	}
}
