package jmespath

import (
	msh "github.com/wrmsr/bane/pkg/util/marshal"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

var _ = msh.Register(rfl.TypeOf[Node](),
	msh.SetImpl{Impl: rfl.TypeOf[And](), Tag: "and"},
	msh.SetImpl{Impl: rfl.TypeOf[Call](), Tag: "call"},
	msh.SetImpl{Impl: rfl.TypeOf[Cmp](), Tag: "cmp"},
	msh.SetImpl{Impl: rfl.TypeOf[CreateArray](), Tag: "create_array"},
	msh.SetImpl{Impl: rfl.TypeOf[CreateObject](), Tag: "create_object"},
	msh.SetImpl{Impl: rfl.TypeOf[Current](), Tag: "current"},
	msh.SetImpl{Impl: rfl.TypeOf[ExprRef](), Tag: "expr_ref"},
	msh.SetImpl{Impl: rfl.TypeOf[FlattenArray](), Tag: "flatten_array"},
	msh.SetImpl{Impl: rfl.TypeOf[FlattenObject](), Tag: "flatten_object"},
	msh.SetImpl{Impl: rfl.TypeOf[Index](), Tag: "index"},
	msh.SetImpl{Impl: rfl.TypeOf[JsonLiteral](), Tag: "json_literal"},
	msh.SetImpl{Impl: rfl.TypeOf[NameTarget](), Tag: "name_target"},
	msh.SetImpl{Impl: rfl.TypeOf[Negate](), Tag: "negate"},
	msh.SetImpl{Impl: rfl.TypeOf[NumberTarget](), Tag: "number_target"},
	msh.SetImpl{Impl: rfl.TypeOf[Or](), Tag: "or"},
	msh.SetImpl{Impl: rfl.TypeOf[Parameter](), Tag: "parameter"},
	msh.SetImpl{Impl: rfl.TypeOf[Project](), Tag: "project"},
	msh.SetImpl{Impl: rfl.TypeOf[Property](), Tag: "property"},
	msh.SetImpl{Impl: rfl.TypeOf[Selection](), Tag: "selection"},
	msh.SetImpl{Impl: rfl.TypeOf[Sequence](), Tag: "sequence"},
	msh.SetImpl{Impl: rfl.TypeOf[Slice](), Tag: "slice"},
	msh.SetImpl{Impl: rfl.TypeOf[String](), Tag: "string"},
)
