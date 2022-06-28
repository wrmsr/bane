package container

import its "github.com/wrmsr/bane/pkg/util/iterators"

type Mutable interface {
	isMutable()
}

type Ordered interface {
	isOrdered()
}

type Persistent interface {
	isPersistent()
}

type Sorted[A, I any] interface {
	isSorted()

	its.Iterable[I]
	ReverseIterate() its.Iterator[I]

	IterateFrom(a A) its.Iterator[I]
	ReverseIterateFrom(a A) its.Iterator[I]
}
