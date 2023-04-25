package types

type Traversable[T any] interface {
	ForEach(fn func(v T) bool) bool
}
