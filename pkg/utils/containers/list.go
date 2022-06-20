package containers

import its "github.com/wrmsr/bane/pkg/utils/iterators"

//

type List[T any] interface {
	Len() int
	Contains(T) bool
	Get(i int) T
	TryGet(i int) (T, bool)

	its.Iterable[T]
	its.Traversable[T]
}

type MutList[T any] interface {
	List[T]
}
