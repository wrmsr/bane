package container

import (
	"reflect"

	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Sync interface {
	isSync()
}

type Mutable interface {
	isMutable()
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

//

type List[T any] interface {
	Len() int
	Get(i int) T

	its.Iterable[T]
	its.Traversable[T]
}

type MutList[T any] interface {
	List[T]
	Mutable
	//Decay[List[T]]

	Append(v T)
	Delete(i int)
}

//

type Set[T any] interface {
	Len() int
	Contains(v T) bool

	its.Iterable[T]
	its.Traversable[T]
}

type MutSet[T any] interface {
	Set[T]
	Mutable
	//Decay[Set[T]]

	Add(v T)
	TryAdd(v T) bool
	Remove(v T)
	TryRemove(v T) bool
}

//

type Map[K, V any] interface {
	Len() int
	Contains(k K) bool
	Get(k K) V
	TryGet(k K) (V, bool)

	its.Iterable[bt.Kv[K, V]]
	its.Traversable[bt.Kv[K, V]]
}

type MutMap[K, V any] interface {
	Map[K, V]
	Mutable
	//Decay[Map[K, V]]

	Put(k K, v V)
	Delete(k K)
	Default(k K, v V) bool
}

//

type OrderedMap[K, V any] interface {
	Map[K, V]
	Ordered[bt.Kv[K, V]]
}

type MutOrderedMap[K, V any] interface {
	OrderedMap[K, V]
	MutMap[K, V]
}

//

type SyncedMap[K, V any] interface {
	Map[K, V]
	Sync
}

type SyncedMutMap[K, V any] interface {
	MutMap[K, V]
	Sync
}

//

type BiMap[K, V comparable] interface {
	Map[K, V]

	Invert() BiMap[V, K]
}

type MutBiMap[K, V comparable] interface {
	MutMap[K, V]
	BiMap[K, V]

	MutInvert() MutBiMap[V, K]
}

//

type TypeMap[T any] interface {
	Len() int
	Contains(ty reflect.Type) bool
	Get(ty reflect.Type) T
	TryGet(ty reflect.Type) (T, bool)

	its.Iterable[T]
	its.Traversable[T]
}

type MutTypeMap[T any] interface {
	TypeMap[T]
	Mutable

	Put(v T)
	Delete(ty reflect.Type)
	Default(v T) bool
}

//

type PersistentMap[K, V any] interface {
	Map[K, V]
	Persistent[bt.Kv[K, V]]

	With(k K, v V) PersistentMap[K, V]
	Without(k K) PersistentMap[K, V]
	Default(k K, v V) PersistentMap[K, V]
}

//

type Decay[T any] interface {
	Decay() T
}
