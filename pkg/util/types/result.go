package types

//

type Result[T any] struct {
	Val T
	Err error
}

func (r Result[T]) GetVal() T     { return r.Val }
func (r Result[T]) GetErr() error { return r.Err }

func (r Result[T]) Unpack() (T, error) { return r.Val, r.Err }

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
