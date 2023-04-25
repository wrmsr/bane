package types

type BiFunc[F, T any] struct {
	t func(f F) T
	f func(t T) F
}

func BiFuncOf[F, T any](t func(F) T, f func(T) F) BiFunc[F, T] {
	return BiFunc[F, T]{t: t, f: f}
}

func (b BiFunc[F, T]) Call(f F) T {
	return b.t(f)
}

func (b BiFunc[F, T]) Flip() BiFunc[T, F] {
	return BiFunc[T, F]{t: b.f, f: b.t}
}

func (b BiFunc[F, T]) Unwrap() func(F) T {
	return b.t
}
