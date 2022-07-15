/*
 */
package jmespath

import (
	msh "github.com/wrmsr/bane/pkg/util/marshal"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

var _ = msh.Register(rfl.TypeOf[Node](),
	msh.SetImplOf[And]("and"),
	msh.SetImplOf[Call]("call"),
	msh.SetImplOf[Cmp]("cmp"),
	msh.SetImplOf[CreateArray]("create_array"),
	msh.SetImplOf[CreateObject]("create_object"),
	msh.SetImplOf[Current]("current"),
	msh.SetImplOf[ExprRef]("expr_ref"),
	msh.SetImplOf[FlattenArray]("flatten_array"),
	msh.SetImplOf[FlattenObject]("flatten_object"),
	msh.SetImplOf[Index]("index"),
	msh.SetImplOf[JsonLiteral]("json_literal"),
	msh.SetImplOf[NameTarget]("name_target"),
	msh.SetImplOf[Negate]("negate"),
	msh.SetImplOf[NumberTarget]("number_target"),
	msh.SetImplOf[Or]("or"),
	msh.SetImplOf[Parameter]("parameter"),
	msh.SetImplOf[Project]("project"),
	msh.SetImplOf[Property]("property"),
	msh.SetImplOf[Selection]("selection"),
	msh.SetImplOf[Sequence]("sequence"),
	msh.SetImplOf[Slice]("slice"),
	msh.SetImplOf[String]("string"),
)

var _ = msh.Register(rfl.TypeOf[CmpOp](),
	msh.SetEnumTypes(map[CmpOp]string{
		CmpEq: "=",
		CmpNe: "!=",
		CmpGt: ">",
		CmpGe: ">=",
		CmpLt: "<",
		CmpLe: "<=",
	}),
)

var _ = msh.Register(rfl.TypeOf[Target](),
	msh.SetImplOf[NumberTarget]("number"),
	msh.SetImplOf[NameTarget]("name"),
)
