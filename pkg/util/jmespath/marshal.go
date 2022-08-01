/*
 */
package jmespath

import (
	msh "github.com/wrmsr/bane/pkg/util/marshal"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

var _ = msh.Register(rfl.TypeOf[Node](),
	msh.SetImplOf[And](),
	msh.SetImplOf[Call](),
	msh.SetImplOf[Cmp](),
	msh.SetImplOf[CreateArray](),
	msh.SetImplOf[CreateObject](),
	msh.SetImplOf[Current](),
	msh.SetImplOf[ExprRef](),
	msh.SetImplOf[FlattenArray](),
	msh.SetImplOf[FlattenObject](),
	msh.SetImplOf[Index](),
	msh.SetImplOf[JsonLiteral](),
	msh.SetImplOf[Negate](),
	msh.SetImplOf[Or](),
	msh.SetImplOf[Parameter](),
	msh.SetImplOf[Project](),
	msh.SetImplOf[Property](),
	msh.SetImplOf[Selection](),
	msh.SetImplOf[Sequence](),
	msh.SetImplOf[Slice](),
	msh.SetImplOf[String](),
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
