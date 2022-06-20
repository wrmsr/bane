package containers

import its "github.com/wrmsr/bane/pkg/utils/iterators"

type List[T any] interface {
	Len() int
	Contains(T) bool

	its.CanIterate[T]
}
