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
