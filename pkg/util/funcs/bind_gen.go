package funcs

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

func Bind3x1x1[B0, B1, B2, A0, R0 any](fn func(B0, B1, B2, A0) R0, b0 B0, b1 B1, b2 B2) func(A0) R0 {
	return func(a0 A0) R0 {
		return fn(b0, b1, b2, a0)
	}
}

func Bind1x2x1[B0, A0, A1, R0 any](fn func(B0, A0, A1) R0, b0 B0) func(A0, A1) R0 {
	return func(a0 A0, a1 A1) R0 {
		return fn(b0, a0, a1)
	}
}

func Bind2x2x1[B0, B1, A0, A1, R0 any](fn func(B0, B1, A0, A1) R0, b0 B0, b1 B1) func(A0, A1) R0 {
	return func(a0 A0, a1 A1) R0 {
		return fn(b0, b1, a0, a1)
	}
}

func Bind3x2x1[B0, B1, B2, A0, A1, R0 any](fn func(B0, B1, B2, A0, A1) R0, b0 B0, b1 B1, b2 B2) func(A0, A1) R0 {
	return func(a0 A0, a1 A1) R0 {
		return fn(b0, b1, b2, a0, a1)
	}
}

func Bind1x3x1[B0, A0, A1, A2, R0 any](fn func(B0, A0, A1, A2) R0, b0 B0) func(A0, A1, A2) R0 {
	return func(a0 A0, a1 A1, a2 A2) R0 {
		return fn(b0, a0, a1, a2)
	}
}

func Bind2x3x1[B0, B1, A0, A1, A2, R0 any](fn func(B0, B1, A0, A1, A2) R0, b0 B0, b1 B1) func(A0, A1, A2) R0 {
	return func(a0 A0, a1 A1, a2 A2) R0 {
		return fn(b0, b1, a0, a1, a2)
	}
}

func Bind3x3x1[B0, B1, B2, A0, A1, A2, R0 any](fn func(B0, B1, B2, A0, A1, A2) R0, b0 B0, b1 B1, b2 B2) func(A0, A1, A2) R0 {
	return func(a0 A0, a1 A1, a2 A2) R0 {
		return fn(b0, b1, b2, a0, a1, a2)
	}
}

func Bind1x1x2[B0, A0, R0, R1 any](fn func(B0, A0) (R0, R1), b0 B0) func(A0) (R0, R1) {
	return func(a0 A0) (R0, R1) {
		return fn(b0, a0)
	}
}

func Bind2x1x2[B0, B1, A0, R0, R1 any](fn func(B0, B1, A0) (R0, R1), b0 B0, b1 B1) func(A0) (R0, R1) {
	return func(a0 A0) (R0, R1) {
		return fn(b0, b1, a0)
	}
}

func Bind3x1x2[B0, B1, B2, A0, R0, R1 any](fn func(B0, B1, B2, A0) (R0, R1), b0 B0, b1 B1, b2 B2) func(A0) (R0, R1) {
	return func(a0 A0) (R0, R1) {
		return fn(b0, b1, b2, a0)
	}
}

func Bind1x2x2[B0, A0, A1, R0, R1 any](fn func(B0, A0, A1) (R0, R1), b0 B0) func(A0, A1) (R0, R1) {
	return func(a0 A0, a1 A1) (R0, R1) {
		return fn(b0, a0, a1)
	}
}

func Bind2x2x2[B0, B1, A0, A1, R0, R1 any](fn func(B0, B1, A0, A1) (R0, R1), b0 B0, b1 B1) func(A0, A1) (R0, R1) {
	return func(a0 A0, a1 A1) (R0, R1) {
		return fn(b0, b1, a0, a1)
	}
}

func Bind3x2x2[B0, B1, B2, A0, A1, R0, R1 any](fn func(B0, B1, B2, A0, A1) (R0, R1), b0 B0, b1 B1, b2 B2) func(A0, A1) (R0, R1) {
	return func(a0 A0, a1 A1) (R0, R1) {
		return fn(b0, b1, b2, a0, a1)
	}
}

func Bind1x3x2[B0, A0, A1, A2, R0, R1 any](fn func(B0, A0, A1, A2) (R0, R1), b0 B0) func(A0, A1, A2) (R0, R1) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1) {
		return fn(b0, a0, a1, a2)
	}
}

func Bind2x3x2[B0, B1, A0, A1, A2, R0, R1 any](fn func(B0, B1, A0, A1, A2) (R0, R1), b0 B0, b1 B1) func(A0, A1, A2) (R0, R1) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1) {
		return fn(b0, b1, a0, a1, a2)
	}
}

func Bind3x3x2[B0, B1, B2, A0, A1, A2, R0, R1 any](fn func(B0, B1, B2, A0, A1, A2) (R0, R1), b0 B0, b1 B1, b2 B2) func(A0, A1, A2) (R0, R1) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1) {
		return fn(b0, b1, b2, a0, a1, a2)
	}
}

func Bind1x1x3[B0, A0, R0, R1, R2 any](fn func(B0, A0) (R0, R1, R2), b0 B0) func(A0) (R0, R1, R2) {
	return func(a0 A0) (R0, R1, R2) {
		return fn(b0, a0)
	}
}

func Bind2x1x3[B0, B1, A0, R0, R1, R2 any](fn func(B0, B1, A0) (R0, R1, R2), b0 B0, b1 B1) func(A0) (R0, R1, R2) {
	return func(a0 A0) (R0, R1, R2) {
		return fn(b0, b1, a0)
	}
}

func Bind3x1x3[B0, B1, B2, A0, R0, R1, R2 any](fn func(B0, B1, B2, A0) (R0, R1, R2), b0 B0, b1 B1, b2 B2) func(A0) (R0, R1, R2) {
	return func(a0 A0) (R0, R1, R2) {
		return fn(b0, b1, b2, a0)
	}
}

func Bind1x2x3[B0, A0, A1, R0, R1, R2 any](fn func(B0, A0, A1) (R0, R1, R2), b0 B0) func(A0, A1) (R0, R1, R2) {
	return func(a0 A0, a1 A1) (R0, R1, R2) {
		return fn(b0, a0, a1)
	}
}

func Bind2x2x3[B0, B1, A0, A1, R0, R1, R2 any](fn func(B0, B1, A0, A1) (R0, R1, R2), b0 B0, b1 B1) func(A0, A1) (R0, R1, R2) {
	return func(a0 A0, a1 A1) (R0, R1, R2) {
		return fn(b0, b1, a0, a1)
	}
}

func Bind3x2x3[B0, B1, B2, A0, A1, R0, R1, R2 any](fn func(B0, B1, B2, A0, A1) (R0, R1, R2), b0 B0, b1 B1, b2 B2) func(A0, A1) (R0, R1, R2) {
	return func(a0 A0, a1 A1) (R0, R1, R2) {
		return fn(b0, b1, b2, a0, a1)
	}
}

func Bind1x3x3[B0, A0, A1, A2, R0, R1, R2 any](fn func(B0, A0, A1, A2) (R0, R1, R2), b0 B0) func(A0, A1, A2) (R0, R1, R2) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1, R2) {
		return fn(b0, a0, a1, a2)
	}
}

func Bind2x3x3[B0, B1, A0, A1, A2, R0, R1, R2 any](fn func(B0, B1, A0, A1, A2) (R0, R1, R2), b0 B0, b1 B1) func(A0, A1, A2) (R0, R1, R2) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1, R2) {
		return fn(b0, b1, a0, a1, a2)
	}
}

func Bind3x3x3[B0, B1, B2, A0, A1, A2, R0, R1, R2 any](fn func(B0, B1, B2, A0, A1, A2) (R0, R1, R2), b0 B0, b1 B1, b2 B2) func(A0, A1, A2) (R0, R1, R2) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1, R2) {
		return fn(b0, b1, b2, a0, a1, a2)
	}
}
