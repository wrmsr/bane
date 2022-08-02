package container

import (
	"reflect"

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

	bt.Iterable[T]
}

type Persistent[T any] interface {
	isPersistent()

	bt.Iterable[T]
}

type Sorted[A, I any] interface {
	isSorted()

	bt.Iterable[I]
	ReverseIterate() bt.Iterator[I]

	IterateFrom(a A) bt.Iterator[I]
	ReverseIterateFrom(a A) bt.Iterator[I]
}

//

type List[T any] interface {
	Len() int
	Get(i int) T

	bt.Iterable[T]
	bt.Traversable[T]
}

type MutList[T any] interface {
	List[T]
	Mutable

	Put(i int, v T)
	Append(v T)
	Delete(i int)
}

//

type Set[T any] interface {
	Len() int
	Contains(v T) bool

	bt.Iterable[T]
	bt.Traversable[T]
}

type MutSet[T any] interface {
	Set[T]
	Mutable

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

	bt.Iterable[bt.Kv[K, V]]
	bt.Traversable[bt.Kv[K, V]]
}

type MutMap[K, V any] interface {
	Map[K, V]
	Mutable

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

type SyncMap[K, V any] interface {
	Map[K, V]
	Sync
}

type SyncMutMap[K, V any] interface {
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

	bt.Iterable[T]
	bt.Traversable[T]
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
	Mutable

	Decay() T
}
