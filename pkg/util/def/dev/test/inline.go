//go:build !nodev

package test

import "github.com/wrmsr/bane/pkg/util/def"

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

// ==>

func _def_inl_Bar(x, y int) int {
	var __def_inl_0 int
	__def_inl_1 := x
	__def_inl_2 := y
	{
		__def_inl_0 = __def_inl_1 + __def_inl_2
	}
	var __def_inl_3 int
	__def_inl_4 := __def_inl_0
	__def_inl_5 := y
	{
		__def_inl_3 = __def_inl_4 * __def_inl_5
	}

	return __def_inl_3
}
