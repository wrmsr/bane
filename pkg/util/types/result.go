package types

type Result[T any] struct {
	Val T
	Err error
}

func Ok[T any](val T) Result[T] {
	return Result[T]{Val: val}
}

func Err[T any](err error) Result[T] {
	return Result[T]{Err: err}
}

func AsResult[T any](v T, err error) Result[T] {
	return Result[T]{v, err}
}

func AsResultFn[T any](fn func() (T, error)) func() Result[T] {
	return func() Result[T] { return AsResult(fn()) }
}
