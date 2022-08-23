package types

type AnySlice interface {
	Len() int
	Get(int) any
}

type AnySliceImpl[T any] struct {
	s []T
}

func AnySliceOf[T any](s []T) AnySliceImpl[T] {
	return AnySliceImpl[T]{s: s}
}

var _ AnySlice = AnySliceImpl[int]{}

func (v AnySliceImpl[T]) Len() int {
	return len(v.s)
}

func (v AnySliceImpl[T]) Get(i int) any {
	return v.s[i]
}

func (v AnySliceImpl[T]) Set(i int, e any) {
	v.s[i] = e.(T)
}
