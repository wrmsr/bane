package funcs

func Recurse[F, T any](fn func(func(F) T, F) T, f F) T {
	var rec func(F) T
	rec = func(f F) T {
		return fn(rec, f)
	}
	return rec(f)
}

func Not(fn func() bool) func() bool {
	return func() bool {
		return !fn()
	}
}

func Not1[T any](fn func(t T) bool) func(t T) bool {
	return func(t T) bool {
		return !fn(t)
	}
}

func Not2[T, U any](fn func(t T, u U) bool) func(t T, u U) bool {
	return func(t T, u U) bool {
		return !fn(t, u)
	}
}

func Returning[R any](fn func(), r R) func() R {
	return func() R {
		fn()
		return r
	}
}

func Returning1[T, R any](fn func(T), r R) func(T) R {
	return func(t T) R {
		fn(t)
		return r
	}
}

func Drop1[T, R any](fn func() R) func(T) R {
	return func(T) R { return fn() }
}

func Drop2[T, U, R any](fn func() R) func(T, U) R {
	return func(t T, u U) R { return fn() }
}

func Drop2x1[T, U, R any](fn func(U) R) func(T, U) R {
	return func(t T, u U) R { return fn(u) }
}

func DropR2x1[T, U, R any](fn func(T) R) func(T, U) R {
	return func(t T, u U) R { return fn(t) }
}
