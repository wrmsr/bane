package test

//

func Inline(fns ...any) any {
	return nil
}

func WithInlining(fn any) any {
	return nil
}

//

func Inlined0(x, y int) int {
	return x + y
}

func Inlined1(x, y int) int {
	return x * y
}

func Bar(x, y int) int {
	return Inlined1(Inlined0(x, y), y)
}

//

func _def_Bar(x, y int) int {
	_def_tmp_0 := x + y
	_def_tmp_1 := _def_tmp_0 * y
	return _def_tmp_1
}
