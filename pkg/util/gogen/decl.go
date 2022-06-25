package gogen

import opt "github.com/wrmsr/bane/pkg/util/optional"

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
