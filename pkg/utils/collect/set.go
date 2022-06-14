package collect

type Set[T comparable] interface {
	Contains(T) bool
}

type MutSet[T comparable] interface {
	Set[T]
}
