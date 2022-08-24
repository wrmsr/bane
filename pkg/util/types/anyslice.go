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

//

type AnySliceIterator struct {
	s AnySlice
	c int
}

var _ Iterator[any] = &AnySliceIterator{}

func (i *AnySliceIterator) Iterate() Iterator[any] {
	return i
}

func (i *AnySliceIterator) HasNext() bool {
	return i.c < i.s.Len()
}

func (i *AnySliceIterator) Next() any {
	if !i.HasNext() {
		panic(IteratorExhaustedError{})
	}
	v := i.s.Get(i.c)
	i.c++
	return v
}

func IterateAnySlice(s AnySlice) *AnySliceIterator {
	return &AnySliceIterator{s: s}
}

func (s AnySliceImpl[T]) Iterate() Iterator[any] {
	return IterateAnySlice(s)
}

//

func (s AnySliceImpl[T]) ForEach(fn func(v any) bool) bool {
	for _, e := range s.s {
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

func (s AnySliceImpl[T]) Len() int {
	return len(s.s)
}

func (s AnySliceImpl[T]) Get(i int) any {
	return s.s[i]
}

//

type MutAnySliceImpl[T any] struct {
	s AnySliceImpl[T]
}

func MutAnySliceOf[T any](s []T) MutAnySliceImpl[T] {
	return MutAnySliceImpl[T]{AnySliceImpl[T]{s: s}}
}

var _ MutAnySlice = MutAnySliceImpl[int]{}

func (s MutAnySliceImpl[T]) Len() int                         { return s.s.Len() }
func (s MutAnySliceImpl[T]) Get(i int) any                    { return s.s.Get(i) }
func (s MutAnySliceImpl[T]) Iterate() Iterator[any]           { return s.s.Iterate() }
func (s MutAnySliceImpl[T]) ForEach(fn func(v any) bool) bool { return s.s.ForEach(fn) }

func (s MutAnySliceImpl[T]) Set(i int, e any) {
	s.s.s[i] = e.(T)
}
