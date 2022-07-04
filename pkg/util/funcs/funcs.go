package funcs

func Const[T any](v T) func() T {
	return func() T { return v }
}

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
