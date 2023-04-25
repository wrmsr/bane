//go:build !nodev

package test

import "github.com/wrmsr/bane/core/def"

//

var _ = def.Inline(Inlined0, Inlined1)

func Inlined0(x, y int) int {
	return x + y
}

func Inlined1(x, y int) int {
	return x * y
}

var _ = def.WithInline(Bar)

func Bar(x, y int) int {
	//var foo0 = 10
	//foo1, foo2 := 20, 30
	//_ = foo0 + foo1 + foo2

	//Inlined0(x, y)

	return Inlined1(Inlined0(x, y), y)
}

//

type InlineThing struct {
	x int
}

var _ = def.Inline(InlineThing.Frob)

func (t InlineThing) Frob(y int) int {
	return t.x + y
}

var _ = def.WithInline(Baz)

func Baz(t InlineThing, y int) int {
	return t.Frob(y) + 1
}

// ==>

//func _def_inl_Bar(x, y int) int {
//	var __def_inl_0 int
//	__def_inl_1 := x
//	__def_inl_2 := y
//	{
//		{
//			__def_inl_0 = __def_inl_1 + __def_inl_2
//			goto __def_inl_3
//		}
//	__def_inl_3:
//	}
//	var __def_inl_4 int
//	__def_inl_5 := __def_inl_0
//	__def_inl_6 := y
//	{
//		{
//			__def_inl_4 = __def_inl_5 * __def_inl_6
//			goto __def_inl_7
//		}
//	__def_inl_7:
//	}
//
//	return __def_inl_4
//}
