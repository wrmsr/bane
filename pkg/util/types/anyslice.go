package types

//

type AnySlice interface {
	Iterable[any]
	Traversable[any]

	Len() int
	Get(int) any
}

type MutAnySlice interface {
	AnySlice

	Set(int, any)
}

//

type AnySliceImpl[T any] struct {
	s []T
}

type anySliceIterator[T any] struct {
	s []T
	c int
}

var _ Iterator[any] = &anySliceIterator[int]{}

func (i *anySliceIterator[T]) Iterate() Iterator[any] {
	return i
}

func (i *anySliceIterator[T]) HasNext() bool {
	return i.c < len(i.s)
}

func (i *anySliceIterator[T]) Next() any {
	if !i.HasNext() {
		panic(IteratorExhaustedError{})
	}
	v := i.s[i.c]
	i.c++
	return v
}

func (a AnySliceImpl[T]) Iterate() Iterator[any] {
	return &anySliceIterator[T]{s: a.s}
}

func (a AnySliceImpl[T]) ForEach(fn func(v any) bool) bool {
	for _, e := range a.s {
		if !fn(e) {
			return false
		}
	}
	return true
}

func AnySliceOf[T any](s []T) AnySliceImpl[T] {
	return AnySliceImpl[T]{s: s}
}

var _ AnySlice = AnySliceImpl[int]{}

func (a AnySliceImpl[T]) Len() int {
	return len(a.s)
}

func (a AnySliceImpl[T]) Get(i int) any {
	return a.s[i]
}

//

type MutAnySliceImpl[T any] struct {
	AnySliceImpl[T]
}

var _ MutAnySlice = MutAnySliceImpl[int]{}

func MutAnySliceOf[T any](s []T) MutAnySliceImpl[T] {
	return MutAnySliceImpl[T]{AnySliceImpl[T]{s: s}}
}

func (a MutAnySliceImpl[T]) Set(i int, e any) {
	a.s[i] = e.(T)
}
