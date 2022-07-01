package iterators

//

type Traversable[T any] interface {
	ForEach(fn func(v T) bool) bool
}

func ForEach[T any](it Iterable[T], fn func(v T) bool) bool {
	for it := it.Iterate(); it.HasNext(); {
		if !fn(it.Next()) {
			return false
		}
	}
	return true
}

func ForAll[T any](tv Traversable[T], fn func(v T)) {
	tv.ForEach(func(v T) bool {
		fn(v)
		return true
	})
}

func SeqForEach[T any](t Traversable[T]) []T {
	var s []T
	t.ForEach(func(v T) bool {
		s = append(s, v)
		return true
	})
	return s
}

//

type traversableIterable[T any] struct {
	Iterable[T]
}

func AsTraversable[T any](it Iterable[T]) Traversable[T] {
	return traversableIterable[T]{it}
}

var _ Traversable[any] = traversableIterable[any]{}

func (t traversableIterable[T]) ForEach(fn func(v T) bool) bool {
	return ForEach[T](t, fn)
}

//

type fnTraversable[T any] struct {
	fn func(fn func(v T) bool) bool
}

var _ Traversable[any] = fnTraversable[any]{}

func (f fnTraversable[T]) ForEach(fn func(v T) bool) bool {
	return f.fn(fn)
}
