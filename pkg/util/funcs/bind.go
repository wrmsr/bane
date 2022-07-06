package funcs

//go:generate go run github.com/wrmsr/bane/pkg/util/funcs/dev/genbind 2 2 2

/*
func Bind1x1x1[B0, A0, R0 any](fn func(B0, A0) R0, b0 B0) func(A0) R0 {
	return func(a0 A0) R0 {
		return fn(b0, a0)
	}
}

func Bind2x1x1[B0, B1, A0, R0 any](fn func(B0, B1, A0) R0, b0 B0, b1 B1) func(A0) R0 {
	return func(a0 A0) R0 {
		return fn(b0, b1, a0)
	}
}

func Bind1x2x1[B0, A0, A1, R0 any](fn func(B0, A0, A1) R0, b0 B0) func(A0, A1) R0 {
	return func(a0 A0, a1 A1) R0 {
		return fn(b0, a0, a1)
	}
}

func Bind1x1x2[B0, A0, R0, R1 any](fn func(B0, A0) (R0, R1), b0 B0) func(A0) (R0, R1) {
	return func(a0 A0) (R0, R1) {
		return fn(b0, a0)
	}
}
*/
