package types

//

type Lesser[T any] interface {
	Less(o T) bool
}

//

type LessImpl[T any] func(l, r T) bool

func CmpLessImpl[T any](cmp CmpImpl[T]) LessImpl[T] {
	return func(l, r T) bool {
		return cmp(l, r) == CmpLess
	}
}
