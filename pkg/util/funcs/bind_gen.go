package funcs

func Bind0x0x0(fn func()) func() {
	return func() {
		fn()
	}
}

func BindR0x0x0(fn func()) func() {
	return func() {
		fn()
	}
}

func Bind1x0x0[B0 any](fn func(B0), b0 B0) func() {
	return func() {
		fn(b0)
	}
}

func BindR0x1x0[B0 any](fn func(B0), b0 B0) func() {
	return func() {
		fn(b0)
	}
}

func Bind2x0x0[B0, B1 any](fn func(B0, B1), b0 B0, b1 B1) func() {
	return func() {
		fn(b0, b1)
	}
}

func BindR0x2x0[B0, B1 any](fn func(B0, B1), b0 B0, b1 B1) func() {
	return func() {
		fn(b0, b1)
	}
}

func Bind3x0x0[B0, B1, B2 any](fn func(B0, B1, B2), b0 B0, b1 B1, b2 B2) func() {
	return func() {
		fn(b0, b1, b2)
	}
}

func BindR0x3x0[B0, B1, B2 any](fn func(B0, B1, B2), b0 B0, b1 B1, b2 B2) func() {
	return func() {
		fn(b0, b1, b2)
	}
}

func Bind0x1x0[A0 any](fn func(A0)) func(A0) {
	return func(a0 A0) {
		fn(a0)
	}
}

func BindR1x0x0[A0 any](fn func(A0)) func(A0) {
	return func(a0 A0) {
		fn(a0)
	}
}

func Bind1x1x0[B0, A0 any](fn func(B0, A0), b0 B0) func(A0) {
	return func(a0 A0) {
		fn(b0, a0)
	}
}

func BindR1x1x0[A0, B0 any](fn func(A0, B0), b0 B0) func(A0) {
	return func(a0 A0) {
		fn(a0, b0)
	}
}

func Bind2x1x0[B0, B1, A0 any](fn func(B0, B1, A0), b0 B0, b1 B1) func(A0) {
	return func(a0 A0) {
		fn(b0, b1, a0)
	}
}

func BindR1x2x0[A0, B0, B1 any](fn func(A0, B0, B1), b0 B0, b1 B1) func(A0) {
	return func(a0 A0) {
		fn(a0, b0, b1)
	}
}

func Bind3x1x0[B0, B1, B2, A0 any](fn func(B0, B1, B2, A0), b0 B0, b1 B1, b2 B2) func(A0) {
	return func(a0 A0) {
		fn(b0, b1, b2, a0)
	}
}

func BindR1x3x0[A0, B0, B1, B2 any](fn func(A0, B0, B1, B2), b0 B0, b1 B1, b2 B2) func(A0) {
	return func(a0 A0) {
		fn(a0, b0, b1, b2)
	}
}

func Bind0x2x0[A0, A1 any](fn func(A0, A1)) func(A0, A1) {
	return func(a0 A0, a1 A1) {
		fn(a0, a1)
	}
}

func BindR2x0x0[A0, A1 any](fn func(A0, A1)) func(A0, A1) {
	return func(a0 A0, a1 A1) {
		fn(a0, a1)
	}
}

func Bind1x2x0[B0, A0, A1 any](fn func(B0, A0, A1), b0 B0) func(A0, A1) {
	return func(a0 A0, a1 A1) {
		fn(b0, a0, a1)
	}
}

func BindR2x1x0[A0, A1, B0 any](fn func(A0, A1, B0), b0 B0) func(A0, A1) {
	return func(a0 A0, a1 A1) {
		fn(a0, a1, b0)
	}
}

func Bind2x2x0[B0, B1, A0, A1 any](fn func(B0, B1, A0, A1), b0 B0, b1 B1) func(A0, A1) {
	return func(a0 A0, a1 A1) {
		fn(b0, b1, a0, a1)
	}
}

func BindR2x2x0[A0, A1, B0, B1 any](fn func(A0, A1, B0, B1), b0 B0, b1 B1) func(A0, A1) {
	return func(a0 A0, a1 A1) {
		fn(a0, a1, b0, b1)
	}
}

func Bind3x2x0[B0, B1, B2, A0, A1 any](fn func(B0, B1, B2, A0, A1), b0 B0, b1 B1, b2 B2) func(A0, A1) {
	return func(a0 A0, a1 A1) {
		fn(b0, b1, b2, a0, a1)
	}
}

func BindR2x3x0[A0, A1, B0, B1, B2 any](fn func(A0, A1, B0, B1, B2), b0 B0, b1 B1, b2 B2) func(A0, A1) {
	return func(a0 A0, a1 A1) {
		fn(a0, a1, b0, b1, b2)
	}
}

func Bind0x3x0[A0, A1, A2 any](fn func(A0, A1, A2)) func(A0, A1, A2) {
	return func(a0 A0, a1 A1, a2 A2) {
		fn(a0, a1, a2)
	}
}

func BindR3x0x0[A0, A1, A2 any](fn func(A0, A1, A2)) func(A0, A1, A2) {
	return func(a0 A0, a1 A1, a2 A2) {
		fn(a0, a1, a2)
	}
}

func Bind1x3x0[B0, A0, A1, A2 any](fn func(B0, A0, A1, A2), b0 B0) func(A0, A1, A2) {
	return func(a0 A0, a1 A1, a2 A2) {
		fn(b0, a0, a1, a2)
	}
}

func BindR3x1x0[A0, A1, A2, B0 any](fn func(A0, A1, A2, B0), b0 B0) func(A0, A1, A2) {
	return func(a0 A0, a1 A1, a2 A2) {
		fn(a0, a1, a2, b0)
	}
}

func Bind2x3x0[B0, B1, A0, A1, A2 any](fn func(B0, B1, A0, A1, A2), b0 B0, b1 B1) func(A0, A1, A2) {
	return func(a0 A0, a1 A1, a2 A2) {
		fn(b0, b1, a0, a1, a2)
	}
}

func BindR3x2x0[A0, A1, A2, B0, B1 any](fn func(A0, A1, A2, B0, B1), b0 B0, b1 B1) func(A0, A1, A2) {
	return func(a0 A0, a1 A1, a2 A2) {
		fn(a0, a1, a2, b0, b1)
	}
}

func Bind3x3x0[B0, B1, B2, A0, A1, A2 any](fn func(B0, B1, B2, A0, A1, A2), b0 B0, b1 B1, b2 B2) func(A0, A1, A2) {
	return func(a0 A0, a1 A1, a2 A2) {
		fn(b0, b1, b2, a0, a1, a2)
	}
}

func BindR3x3x0[A0, A1, A2, B0, B1, B2 any](fn func(A0, A1, A2, B0, B1, B2), b0 B0, b1 B1, b2 B2) func(A0, A1, A2) {
	return func(a0 A0, a1 A1, a2 A2) {
		fn(a0, a1, a2, b0, b1, b2)
	}
}

func Bind0x0x1[R0 any](fn func() R0) func() R0 {
	return func() R0 {
		return fn()
	}
}

func BindR0x0x1[R0 any](fn func() R0) func() R0 {
	return func() R0 {
		return fn()
	}
}

func Bind1x0x1[B0, R0 any](fn func(B0) R0, b0 B0) func() R0 {
	return func() R0 {
		return fn(b0)
	}
}

func BindR0x1x1[B0, R0 any](fn func(B0) R0, b0 B0) func() R0 {
	return func() R0 {
		return fn(b0)
	}
}

func Bind2x0x1[B0, B1, R0 any](fn func(B0, B1) R0, b0 B0, b1 B1) func() R0 {
	return func() R0 {
		return fn(b0, b1)
	}
}

func BindR0x2x1[B0, B1, R0 any](fn func(B0, B1) R0, b0 B0, b1 B1) func() R0 {
	return func() R0 {
		return fn(b0, b1)
	}
}

func Bind3x0x1[B0, B1, B2, R0 any](fn func(B0, B1, B2) R0, b0 B0, b1 B1, b2 B2) func() R0 {
	return func() R0 {
		return fn(b0, b1, b2)
	}
}

func BindR0x3x1[B0, B1, B2, R0 any](fn func(B0, B1, B2) R0, b0 B0, b1 B1, b2 B2) func() R0 {
	return func() R0 {
		return fn(b0, b1, b2)
	}
}

func Bind0x1x1[A0, R0 any](fn func(A0) R0) func(A0) R0 {
	return func(a0 A0) R0 {
		return fn(a0)
	}
}

func BindR1x0x1[A0, R0 any](fn func(A0) R0) func(A0) R0 {
	return func(a0 A0) R0 {
		return fn(a0)
	}
}

func Bind1x1x1[B0, A0, R0 any](fn func(B0, A0) R0, b0 B0) func(A0) R0 {
	return func(a0 A0) R0 {
		return fn(b0, a0)
	}
}

func BindR1x1x1[A0, B0, R0 any](fn func(A0, B0) R0, b0 B0) func(A0) R0 {
	return func(a0 A0) R0 {
		return fn(a0, b0)
	}
}

func Bind2x1x1[B0, B1, A0, R0 any](fn func(B0, B1, A0) R0, b0 B0, b1 B1) func(A0) R0 {
	return func(a0 A0) R0 {
		return fn(b0, b1, a0)
	}
}

func BindR1x2x1[A0, B0, B1, R0 any](fn func(A0, B0, B1) R0, b0 B0, b1 B1) func(A0) R0 {
	return func(a0 A0) R0 {
		return fn(a0, b0, b1)
	}
}

func Bind3x1x1[B0, B1, B2, A0, R0 any](fn func(B0, B1, B2, A0) R0, b0 B0, b1 B1, b2 B2) func(A0) R0 {
	return func(a0 A0) R0 {
		return fn(b0, b1, b2, a0)
	}
}

func BindR1x3x1[A0, B0, B1, B2, R0 any](fn func(A0, B0, B1, B2) R0, b0 B0, b1 B1, b2 B2) func(A0) R0 {
	return func(a0 A0) R0 {
		return fn(a0, b0, b1, b2)
	}
}

func Bind0x2x1[A0, A1, R0 any](fn func(A0, A1) R0) func(A0, A1) R0 {
	return func(a0 A0, a1 A1) R0 {
		return fn(a0, a1)
	}
}

func BindR2x0x1[A0, A1, R0 any](fn func(A0, A1) R0) func(A0, A1) R0 {
	return func(a0 A0, a1 A1) R0 {
		return fn(a0, a1)
	}
}

func Bind1x2x1[B0, A0, A1, R0 any](fn func(B0, A0, A1) R0, b0 B0) func(A0, A1) R0 {
	return func(a0 A0, a1 A1) R0 {
		return fn(b0, a0, a1)
	}
}

func BindR2x1x1[A0, A1, B0, R0 any](fn func(A0, A1, B0) R0, b0 B0) func(A0, A1) R0 {
	return func(a0 A0, a1 A1) R0 {
		return fn(a0, a1, b0)
	}
}

func Bind2x2x1[B0, B1, A0, A1, R0 any](fn func(B0, B1, A0, A1) R0, b0 B0, b1 B1) func(A0, A1) R0 {
	return func(a0 A0, a1 A1) R0 {
		return fn(b0, b1, a0, a1)
	}
}

func BindR2x2x1[A0, A1, B0, B1, R0 any](fn func(A0, A1, B0, B1) R0, b0 B0, b1 B1) func(A0, A1) R0 {
	return func(a0 A0, a1 A1) R0 {
		return fn(a0, a1, b0, b1)
	}
}

func Bind3x2x1[B0, B1, B2, A0, A1, R0 any](fn func(B0, B1, B2, A0, A1) R0, b0 B0, b1 B1, b2 B2) func(A0, A1) R0 {
	return func(a0 A0, a1 A1) R0 {
		return fn(b0, b1, b2, a0, a1)
	}
}

func BindR2x3x1[A0, A1, B0, B1, B2, R0 any](fn func(A0, A1, B0, B1, B2) R0, b0 B0, b1 B1, b2 B2) func(A0, A1) R0 {
	return func(a0 A0, a1 A1) R0 {
		return fn(a0, a1, b0, b1, b2)
	}
}

func Bind0x3x1[A0, A1, A2, R0 any](fn func(A0, A1, A2) R0) func(A0, A1, A2) R0 {
	return func(a0 A0, a1 A1, a2 A2) R0 {
		return fn(a0, a1, a2)
	}
}

func BindR3x0x1[A0, A1, A2, R0 any](fn func(A0, A1, A2) R0) func(A0, A1, A2) R0 {
	return func(a0 A0, a1 A1, a2 A2) R0 {
		return fn(a0, a1, a2)
	}
}

func Bind1x3x1[B0, A0, A1, A2, R0 any](fn func(B0, A0, A1, A2) R0, b0 B0) func(A0, A1, A2) R0 {
	return func(a0 A0, a1 A1, a2 A2) R0 {
		return fn(b0, a0, a1, a2)
	}
}

func BindR3x1x1[A0, A1, A2, B0, R0 any](fn func(A0, A1, A2, B0) R0, b0 B0) func(A0, A1, A2) R0 {
	return func(a0 A0, a1 A1, a2 A2) R0 {
		return fn(a0, a1, a2, b0)
	}
}

func Bind2x3x1[B0, B1, A0, A1, A2, R0 any](fn func(B0, B1, A0, A1, A2) R0, b0 B0, b1 B1) func(A0, A1, A2) R0 {
	return func(a0 A0, a1 A1, a2 A2) R0 {
		return fn(b0, b1, a0, a1, a2)
	}
}

func BindR3x2x1[A0, A1, A2, B0, B1, R0 any](fn func(A0, A1, A2, B0, B1) R0, b0 B0, b1 B1) func(A0, A1, A2) R0 {
	return func(a0 A0, a1 A1, a2 A2) R0 {
		return fn(a0, a1, a2, b0, b1)
	}
}

func Bind3x3x1[B0, B1, B2, A0, A1, A2, R0 any](fn func(B0, B1, B2, A0, A1, A2) R0, b0 B0, b1 B1, b2 B2) func(A0, A1, A2) R0 {
	return func(a0 A0, a1 A1, a2 A2) R0 {
		return fn(b0, b1, b2, a0, a1, a2)
	}
}

func BindR3x3x1[A0, A1, A2, B0, B1, B2, R0 any](fn func(A0, A1, A2, B0, B1, B2) R0, b0 B0, b1 B1, b2 B2) func(A0, A1, A2) R0 {
	return func(a0 A0, a1 A1, a2 A2) R0 {
		return fn(a0, a1, a2, b0, b1, b2)
	}
}

func Bind0x0x2[R0, R1 any](fn func() (R0, R1)) func() (R0, R1) {
	return func() (R0, R1) {
		return fn()
	}
}

func BindR0x0x2[R0, R1 any](fn func() (R0, R1)) func() (R0, R1) {
	return func() (R0, R1) {
		return fn()
	}
}

func Bind1x0x2[B0, R0, R1 any](fn func(B0) (R0, R1), b0 B0) func() (R0, R1) {
	return func() (R0, R1) {
		return fn(b0)
	}
}

func BindR0x1x2[B0, R0, R1 any](fn func(B0) (R0, R1), b0 B0) func() (R0, R1) {
	return func() (R0, R1) {
		return fn(b0)
	}
}

func Bind2x0x2[B0, B1, R0, R1 any](fn func(B0, B1) (R0, R1), b0 B0, b1 B1) func() (R0, R1) {
	return func() (R0, R1) {
		return fn(b0, b1)
	}
}

func BindR0x2x2[B0, B1, R0, R1 any](fn func(B0, B1) (R0, R1), b0 B0, b1 B1) func() (R0, R1) {
	return func() (R0, R1) {
		return fn(b0, b1)
	}
}

func Bind3x0x2[B0, B1, B2, R0, R1 any](fn func(B0, B1, B2) (R0, R1), b0 B0, b1 B1, b2 B2) func() (R0, R1) {
	return func() (R0, R1) {
		return fn(b0, b1, b2)
	}
}

func BindR0x3x2[B0, B1, B2, R0, R1 any](fn func(B0, B1, B2) (R0, R1), b0 B0, b1 B1, b2 B2) func() (R0, R1) {
	return func() (R0, R1) {
		return fn(b0, b1, b2)
	}
}

func Bind0x1x2[A0, R0, R1 any](fn func(A0) (R0, R1)) func(A0) (R0, R1) {
	return func(a0 A0) (R0, R1) {
		return fn(a0)
	}
}

func BindR1x0x2[A0, R0, R1 any](fn func(A0) (R0, R1)) func(A0) (R0, R1) {
	return func(a0 A0) (R0, R1) {
		return fn(a0)
	}
}

func Bind1x1x2[B0, A0, R0, R1 any](fn func(B0, A0) (R0, R1), b0 B0) func(A0) (R0, R1) {
	return func(a0 A0) (R0, R1) {
		return fn(b0, a0)
	}
}

func BindR1x1x2[A0, B0, R0, R1 any](fn func(A0, B0) (R0, R1), b0 B0) func(A0) (R0, R1) {
	return func(a0 A0) (R0, R1) {
		return fn(a0, b0)
	}
}

func Bind2x1x2[B0, B1, A0, R0, R1 any](fn func(B0, B1, A0) (R0, R1), b0 B0, b1 B1) func(A0) (R0, R1) {
	return func(a0 A0) (R0, R1) {
		return fn(b0, b1, a0)
	}
}

func BindR1x2x2[A0, B0, B1, R0, R1 any](fn func(A0, B0, B1) (R0, R1), b0 B0, b1 B1) func(A0) (R0, R1) {
	return func(a0 A0) (R0, R1) {
		return fn(a0, b0, b1)
	}
}

func Bind3x1x2[B0, B1, B2, A0, R0, R1 any](fn func(B0, B1, B2, A0) (R0, R1), b0 B0, b1 B1, b2 B2) func(A0) (R0, R1) {
	return func(a0 A0) (R0, R1) {
		return fn(b0, b1, b2, a0)
	}
}

func BindR1x3x2[A0, B0, B1, B2, R0, R1 any](fn func(A0, B0, B1, B2) (R0, R1), b0 B0, b1 B1, b2 B2) func(A0) (R0, R1) {
	return func(a0 A0) (R0, R1) {
		return fn(a0, b0, b1, b2)
	}
}

func Bind0x2x2[A0, A1, R0, R1 any](fn func(A0, A1) (R0, R1)) func(A0, A1) (R0, R1) {
	return func(a0 A0, a1 A1) (R0, R1) {
		return fn(a0, a1)
	}
}

func BindR2x0x2[A0, A1, R0, R1 any](fn func(A0, A1) (R0, R1)) func(A0, A1) (R0, R1) {
	return func(a0 A0, a1 A1) (R0, R1) {
		return fn(a0, a1)
	}
}

func Bind1x2x2[B0, A0, A1, R0, R1 any](fn func(B0, A0, A1) (R0, R1), b0 B0) func(A0, A1) (R0, R1) {
	return func(a0 A0, a1 A1) (R0, R1) {
		return fn(b0, a0, a1)
	}
}

func BindR2x1x2[A0, A1, B0, R0, R1 any](fn func(A0, A1, B0) (R0, R1), b0 B0) func(A0, A1) (R0, R1) {
	return func(a0 A0, a1 A1) (R0, R1) {
		return fn(a0, a1, b0)
	}
}

func Bind2x2x2[B0, B1, A0, A1, R0, R1 any](fn func(B0, B1, A0, A1) (R0, R1), b0 B0, b1 B1) func(A0, A1) (R0, R1) {
	return func(a0 A0, a1 A1) (R0, R1) {
		return fn(b0, b1, a0, a1)
	}
}

func BindR2x2x2[A0, A1, B0, B1, R0, R1 any](fn func(A0, A1, B0, B1) (R0, R1), b0 B0, b1 B1) func(A0, A1) (R0, R1) {
	return func(a0 A0, a1 A1) (R0, R1) {
		return fn(a0, a1, b0, b1)
	}
}

func Bind3x2x2[B0, B1, B2, A0, A1, R0, R1 any](fn func(B0, B1, B2, A0, A1) (R0, R1), b0 B0, b1 B1, b2 B2) func(A0, A1) (R0, R1) {
	return func(a0 A0, a1 A1) (R0, R1) {
		return fn(b0, b1, b2, a0, a1)
	}
}

func BindR2x3x2[A0, A1, B0, B1, B2, R0, R1 any](fn func(A0, A1, B0, B1, B2) (R0, R1), b0 B0, b1 B1, b2 B2) func(A0, A1) (R0, R1) {
	return func(a0 A0, a1 A1) (R0, R1) {
		return fn(a0, a1, b0, b1, b2)
	}
}

func Bind0x3x2[A0, A1, A2, R0, R1 any](fn func(A0, A1, A2) (R0, R1)) func(A0, A1, A2) (R0, R1) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1) {
		return fn(a0, a1, a2)
	}
}

func BindR3x0x2[A0, A1, A2, R0, R1 any](fn func(A0, A1, A2) (R0, R1)) func(A0, A1, A2) (R0, R1) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1) {
		return fn(a0, a1, a2)
	}
}

func Bind1x3x2[B0, A0, A1, A2, R0, R1 any](fn func(B0, A0, A1, A2) (R0, R1), b0 B0) func(A0, A1, A2) (R0, R1) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1) {
		return fn(b0, a0, a1, a2)
	}
}

func BindR3x1x2[A0, A1, A2, B0, R0, R1 any](fn func(A0, A1, A2, B0) (R0, R1), b0 B0) func(A0, A1, A2) (R0, R1) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1) {
		return fn(a0, a1, a2, b0)
	}
}

func Bind2x3x2[B0, B1, A0, A1, A2, R0, R1 any](fn func(B0, B1, A0, A1, A2) (R0, R1), b0 B0, b1 B1) func(A0, A1, A2) (R0, R1) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1) {
		return fn(b0, b1, a0, a1, a2)
	}
}

func BindR3x2x2[A0, A1, A2, B0, B1, R0, R1 any](fn func(A0, A1, A2, B0, B1) (R0, R1), b0 B0, b1 B1) func(A0, A1, A2) (R0, R1) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1) {
		return fn(a0, a1, a2, b0, b1)
	}
}

func Bind3x3x2[B0, B1, B2, A0, A1, A2, R0, R1 any](fn func(B0, B1, B2, A0, A1, A2) (R0, R1), b0 B0, b1 B1, b2 B2) func(A0, A1, A2) (R0, R1) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1) {
		return fn(b0, b1, b2, a0, a1, a2)
	}
}

func BindR3x3x2[A0, A1, A2, B0, B1, B2, R0, R1 any](fn func(A0, A1, A2, B0, B1, B2) (R0, R1), b0 B0, b1 B1, b2 B2) func(A0, A1, A2) (R0, R1) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1) {
		return fn(a0, a1, a2, b0, b1, b2)
	}
}

func Bind0x0x3[R0, R1, R2 any](fn func() (R0, R1, R2)) func() (R0, R1, R2) {
	return func() (R0, R1, R2) {
		return fn()
	}
}

func BindR0x0x3[R0, R1, R2 any](fn func() (R0, R1, R2)) func() (R0, R1, R2) {
	return func() (R0, R1, R2) {
		return fn()
	}
}

func Bind1x0x3[B0, R0, R1, R2 any](fn func(B0) (R0, R1, R2), b0 B0) func() (R0, R1, R2) {
	return func() (R0, R1, R2) {
		return fn(b0)
	}
}

func BindR0x1x3[B0, R0, R1, R2 any](fn func(B0) (R0, R1, R2), b0 B0) func() (R0, R1, R2) {
	return func() (R0, R1, R2) {
		return fn(b0)
	}
}

func Bind2x0x3[B0, B1, R0, R1, R2 any](fn func(B0, B1) (R0, R1, R2), b0 B0, b1 B1) func() (R0, R1, R2) {
	return func() (R0, R1, R2) {
		return fn(b0, b1)
	}
}

func BindR0x2x3[B0, B1, R0, R1, R2 any](fn func(B0, B1) (R0, R1, R2), b0 B0, b1 B1) func() (R0, R1, R2) {
	return func() (R0, R1, R2) {
		return fn(b0, b1)
	}
}

func Bind3x0x3[B0, B1, B2, R0, R1, R2 any](fn func(B0, B1, B2) (R0, R1, R2), b0 B0, b1 B1, b2 B2) func() (R0, R1, R2) {
	return func() (R0, R1, R2) {
		return fn(b0, b1, b2)
	}
}

func BindR0x3x3[B0, B1, B2, R0, R1, R2 any](fn func(B0, B1, B2) (R0, R1, R2), b0 B0, b1 B1, b2 B2) func() (R0, R1, R2) {
	return func() (R0, R1, R2) {
		return fn(b0, b1, b2)
	}
}

func Bind0x1x3[A0, R0, R1, R2 any](fn func(A0) (R0, R1, R2)) func(A0) (R0, R1, R2) {
	return func(a0 A0) (R0, R1, R2) {
		return fn(a0)
	}
}

func BindR1x0x3[A0, R0, R1, R2 any](fn func(A0) (R0, R1, R2)) func(A0) (R0, R1, R2) {
	return func(a0 A0) (R0, R1, R2) {
		return fn(a0)
	}
}

func Bind1x1x3[B0, A0, R0, R1, R2 any](fn func(B0, A0) (R0, R1, R2), b0 B0) func(A0) (R0, R1, R2) {
	return func(a0 A0) (R0, R1, R2) {
		return fn(b0, a0)
	}
}

func BindR1x1x3[A0, B0, R0, R1, R2 any](fn func(A0, B0) (R0, R1, R2), b0 B0) func(A0) (R0, R1, R2) {
	return func(a0 A0) (R0, R1, R2) {
		return fn(a0, b0)
	}
}

func Bind2x1x3[B0, B1, A0, R0, R1, R2 any](fn func(B0, B1, A0) (R0, R1, R2), b0 B0, b1 B1) func(A0) (R0, R1, R2) {
	return func(a0 A0) (R0, R1, R2) {
		return fn(b0, b1, a0)
	}
}

func BindR1x2x3[A0, B0, B1, R0, R1, R2 any](fn func(A0, B0, B1) (R0, R1, R2), b0 B0, b1 B1) func(A0) (R0, R1, R2) {
	return func(a0 A0) (R0, R1, R2) {
		return fn(a0, b0, b1)
	}
}

func Bind3x1x3[B0, B1, B2, A0, R0, R1, R2 any](fn func(B0, B1, B2, A0) (R0, R1, R2), b0 B0, b1 B1, b2 B2) func(A0) (R0, R1, R2) {
	return func(a0 A0) (R0, R1, R2) {
		return fn(b0, b1, b2, a0)
	}
}

func BindR1x3x3[A0, B0, B1, B2, R0, R1, R2 any](fn func(A0, B0, B1, B2) (R0, R1, R2), b0 B0, b1 B1, b2 B2) func(A0) (R0, R1, R2) {
	return func(a0 A0) (R0, R1, R2) {
		return fn(a0, b0, b1, b2)
	}
}

func Bind0x2x3[A0, A1, R0, R1, R2 any](fn func(A0, A1) (R0, R1, R2)) func(A0, A1) (R0, R1, R2) {
	return func(a0 A0, a1 A1) (R0, R1, R2) {
		return fn(a0, a1)
	}
}

func BindR2x0x3[A0, A1, R0, R1, R2 any](fn func(A0, A1) (R0, R1, R2)) func(A0, A1) (R0, R1, R2) {
	return func(a0 A0, a1 A1) (R0, R1, R2) {
		return fn(a0, a1)
	}
}

func Bind1x2x3[B0, A0, A1, R0, R1, R2 any](fn func(B0, A0, A1) (R0, R1, R2), b0 B0) func(A0, A1) (R0, R1, R2) {
	return func(a0 A0, a1 A1) (R0, R1, R2) {
		return fn(b0, a0, a1)
	}
}

func BindR2x1x3[A0, A1, B0, R0, R1, R2 any](fn func(A0, A1, B0) (R0, R1, R2), b0 B0) func(A0, A1) (R0, R1, R2) {
	return func(a0 A0, a1 A1) (R0, R1, R2) {
		return fn(a0, a1, b0)
	}
}

func Bind2x2x3[B0, B1, A0, A1, R0, R1, R2 any](fn func(B0, B1, A0, A1) (R0, R1, R2), b0 B0, b1 B1) func(A0, A1) (R0, R1, R2) {
	return func(a0 A0, a1 A1) (R0, R1, R2) {
		return fn(b0, b1, a0, a1)
	}
}

func BindR2x2x3[A0, A1, B0, B1, R0, R1, R2 any](fn func(A0, A1, B0, B1) (R0, R1, R2), b0 B0, b1 B1) func(A0, A1) (R0, R1, R2) {
	return func(a0 A0, a1 A1) (R0, R1, R2) {
		return fn(a0, a1, b0, b1)
	}
}

func Bind3x2x3[B0, B1, B2, A0, A1, R0, R1, R2 any](fn func(B0, B1, B2, A0, A1) (R0, R1, R2), b0 B0, b1 B1, b2 B2) func(A0, A1) (R0, R1, R2) {
	return func(a0 A0, a1 A1) (R0, R1, R2) {
		return fn(b0, b1, b2, a0, a1)
	}
}

func BindR2x3x3[A0, A1, B0, B1, B2, R0, R1, R2 any](fn func(A0, A1, B0, B1, B2) (R0, R1, R2), b0 B0, b1 B1, b2 B2) func(A0, A1) (R0, R1, R2) {
	return func(a0 A0, a1 A1) (R0, R1, R2) {
		return fn(a0, a1, b0, b1, b2)
	}
}

func Bind0x3x3[A0, A1, A2, R0, R1, R2 any](fn func(A0, A1, A2) (R0, R1, R2)) func(A0, A1, A2) (R0, R1, R2) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1, R2) {
		return fn(a0, a1, a2)
	}
}

func BindR3x0x3[A0, A1, A2, R0, R1, R2 any](fn func(A0, A1, A2) (R0, R1, R2)) func(A0, A1, A2) (R0, R1, R2) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1, R2) {
		return fn(a0, a1, a2)
	}
}

func Bind1x3x3[B0, A0, A1, A2, R0, R1, R2 any](fn func(B0, A0, A1, A2) (R0, R1, R2), b0 B0) func(A0, A1, A2) (R0, R1, R2) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1, R2) {
		return fn(b0, a0, a1, a2)
	}
}

func BindR3x1x3[A0, A1, A2, B0, R0, R1, R2 any](fn func(A0, A1, A2, B0) (R0, R1, R2), b0 B0) func(A0, A1, A2) (R0, R1, R2) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1, R2) {
		return fn(a0, a1, a2, b0)
	}
}

func Bind2x3x3[B0, B1, A0, A1, A2, R0, R1, R2 any](fn func(B0, B1, A0, A1, A2) (R0, R1, R2), b0 B0, b1 B1) func(A0, A1, A2) (R0, R1, R2) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1, R2) {
		return fn(b0, b1, a0, a1, a2)
	}
}

func BindR3x2x3[A0, A1, A2, B0, B1, R0, R1, R2 any](fn func(A0, A1, A2, B0, B1) (R0, R1, R2), b0 B0, b1 B1) func(A0, A1, A2) (R0, R1, R2) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1, R2) {
		return fn(a0, a1, a2, b0, b1)
	}
}

func Bind3x3x3[B0, B1, B2, A0, A1, A2, R0, R1, R2 any](fn func(B0, B1, B2, A0, A1, A2) (R0, R1, R2), b0 B0, b1 B1, b2 B2) func(A0, A1, A2) (R0, R1, R2) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1, R2) {
		return fn(b0, b1, b2, a0, a1, a2)
	}
}

func BindR3x3x3[A0, A1, A2, B0, B1, B2, R0, R1, R2 any](fn func(A0, A1, A2, B0, B1, B2) (R0, R1, R2), b0 B0, b1 B1, b2 B2) func(A0, A1, A2) (R0, R1, R2) {
	return func(a0 A0, a1 A1, a2 A2) (R0, R1, R2) {
		return fn(a0, a1, a2, b0, b1, b2)
	}
}
