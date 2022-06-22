package funcs

func Const[T any](v T) func() T {
	return func() T { return v }
}
