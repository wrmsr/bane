package iterators

type CanForEach[T any] interface {
	ForEach(fn func(v T) bool) bool
}

func ForEach[T any](it CanIterate[T], fn func(v T) bool) bool {
	for i := it.Iterate(); i.HasNext(); {
		if !fn(i.Next()) {
			return false
		}
	}
	return true
}
