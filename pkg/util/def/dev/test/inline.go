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
	var foo0 = 10
	foo1, foo2 := 20, 30
	_ = foo0 + foo1 + foo2

	Inlined0(x, y)
	return Inlined1(Inlined0(x, y), y)
}

//

//func _def_inl_Bar(x, y int) int {
//	_def_tmp_0 := x + y
//	_def_tmp_1 := _def_tmp_0 * y
//	return _def_tmp_1
//}
