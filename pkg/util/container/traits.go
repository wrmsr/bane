package container

import (
	its "github.com/wrmsr/bane/pkg/util/iterators"
)

type Locked interface {
	isLocked()
}

type Mutable interface {
	isMutable()
}

type Decay[T any] interface {
	Decay() T
}

type Ordered[T any] interface {
	isOrdered()

	its.Iterable[T]
}

type Persistent[T any] interface {
	isPersistent()

	its.Iterable[T]
}

type Sorted[A, I any] interface {
	isSorted()

	its.Iterable[I]
	ReverseIterate() its.Iterator[I]

	IterateFrom(a A) its.Iterator[I]
	ReverseIterateFrom(a A) its.Iterator[I]
}
