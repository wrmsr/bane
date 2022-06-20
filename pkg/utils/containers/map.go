package containers

import (
	its "github.com/wrmsr/bane/pkg/utils/iterators"
	bt "github.com/wrmsr/bane/pkg/utils/types"
)

//

type Map[K comparable, V any] interface {
	Len() int
	Contains(k K) bool
	Get(k K) V
	TryGet(k K) (V, bool)

	its.Iterable[bt.Kv[K, V]]
	its.Traversable[bt.Kv[K, V]]
}

type MutMap[K comparable, V any] interface {
	Map[K, V]
	Mutable

	Put(k K, v V) bool
	Delete(k K) bool
	Default(k K, v V) bool
}

//

type mapImpl[K comparable, V any] struct {
	m map[K]V
}

func newMapImpl[K comparable, V any]() mapImpl[K, V] {
	return mapImpl[K, V]{
		m: make(map[K]V),
	}
}

func NewMap[K comparable, V any]() Map[K, V] {
	return newMapImpl[K, V]()
}

var _ Map[int, any] = mapImpl[int, any]{}

func (m mapImpl[K, V]) Len() int {
	return len(m.m)
}

func (m mapImpl[K, V]) Contains(k K) bool {
	_, ok := m.m[k]
	return ok
}

func (m mapImpl[K, V]) Get(k K) V {
	return m.m[k]
}

func (m mapImpl[K, V]) TryGet(k K) (V, bool) {
	v, ok := m.m[k]
	return v, ok
}

func (m mapImpl[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	return its.Map[K, V](m.m).Iterate()
}

func (m mapImpl[K, V]) ForEach(fn func(bt.Kv[K, V]) bool) bool {
	for k, v := range m.m {
		if !fn(bt.KvOf(k, v)) {
			return false
		}
	}
	return true
}

//

type mutMapImpl[K comparable, V any] struct {
	mapImpl[K, V]
}

func NewMutMap[K comparable, V any]() MutMap[K, V] {
	return mutMapImpl[K, V]{newMapImpl[K, V]()}
}

var _ MutMap[int, any] = mutMapImpl[int, any]{}

func (m mutMapImpl[K, V]) isMutable() {}

func (m mutMapImpl[K, V]) Put(k K, v V) bool {
	r := m.Contains(k)
	m.m[k] = v
	return r
}

func (m mutMapImpl[K, V]) Delete(k K) bool {
	r := m.Contains(k)
	delete(m.m, k)
	return r
}

func (m mutMapImpl[K, V]) Default(k K, v V) bool {
	if m.Contains(k) {
		return false
	}
	m.m[k] = v
	return true
}
