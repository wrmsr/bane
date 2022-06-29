package gogen

import (
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

type Func struct {
	decl
	Receiver opt.Optional[Param]
	Name     opt.Optional[Ident]
	Params   []Param
	Type     opt.Optional[Type]
	Body     opt.Optional[Block]
}

func NewFunc(
	receiver opt.Optional[Param],
	name opt.Optional[Ident],
	params []Param,
	type_ opt.Optional[Type],
	body opt.Optional[Block],
) Func {
	return Func{
		Receiver: receiver,
		Name:     name,
		Params:   params,
		Type:     type_,
		Body:     body,
	}
}

//

type Import struct {
	decl
	Spec  string
	Alias opt.Optional[Ident]
}

func NewImportAs(spec string, alias opt.Optional[Ident]) Import {
	return Import{
		Spec:  spec,
		Alias: alias,
	}
}

func NewImport(spec string) Import {
	return NewImportAs(spec, opt.None[Ident]())
}

//

type Imports struct {
	decl
	Imports []Import
}

func NewImports(imports ...Import) Imports {
	return Imports{
		Imports: imports,
	}
}

//

type Package struct {
	decl
	Name Ident
}

func NewPackage(name Ident) Package {
	return Package{
		Name: name,
	}
}

//

type Param struct {
	node
	Name opt.Optional[Ident]
	Type Type
}

func NewParam(name opt.Optional[Ident], type_ Type) Param {
	return Param{
		Name: name,
		Type: type_,
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

//

type Struct struct {
	decl
	Name   Ident
	Fields []StructField
}

func NewStruct(name Ident, fields ...StructField) Struct {
	return Struct{
		Name:   name,
		Fields: fields,
	}
}

//

type StructField struct {
	node
	Name Ident
	Type Type
}

func NewStructField(name Ident, type_ Type) StructField {
	return StructField{
		Name: name,
		Type: type_,
	}
}
