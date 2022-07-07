package container

import (
	"reflect"

	its "github.com/wrmsr/bane/pkg/util/iterators"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Map[K, V any] interface {
	Len() int
	Contains(k K) bool
	Get(k K) V
	TryGet(k K) (V, bool)

	its.Iterable[bt.Kv[K, V]]
	its.Traversable[bt.Kv[K, V]]

	rfl.TypeArgsReflector
}

type MutMap[K, V any] interface {
	Map[K, V]
	Mutable

	Put(k K, v V)
	Delete(k K)
	Default(k K, v V) bool
}

//

type mapImpl[K comparable, V any] struct {
	m map[K]V
}

func newMapImpl[K comparable, V any](it its.Iterable[bt.Kv[K, V]]) mapImpl[K, V] {
	m := make(map[K]V)
	if it != nil {
		for it := it.Iterate(); it.HasNext(); {
			c := it.Next()
			m[c.K] = c.V
		}
	}
	return mapImpl[K, V]{
		m: m,
	}
}

func NewMap[K comparable, V any](it its.Iterable[bt.Kv[K, V]]) Map[K, V] {
	if it, ok := it.(Map[K, V]); ok {
		return it
	}
	return newMapImpl[K, V](it)
}

var _ Map[int, any] = mapImpl[int, any]{}

func (m mapImpl[K, V]) ReflectTypeArgs() []reflect.Type {
	return []reflect.Type{rfl.TypeOf[K](), rfl.TypeOf[V]()}
}

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
	return its.OfMap[K, V](m.m).Iterate()
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

func NewMutMap[K comparable, V any](it its.Iterable[bt.Kv[K, V]]) MutMap[K, V] {
	return mutMapImpl[K, V]{newMapImpl[K, V](it)}
}

func WrapMap[K comparable, V any](m map[K]V) MutMap[K, V] {
	return mutMapImpl[K, V]{mapImpl[K, V]{m}}
}

var _ MutMap[int, any] = mutMapImpl[int, any]{}

func (m mutMapImpl[K, V]) isMutable() {}

func (m mutMapImpl[K, V]) Put(k K, v V) {
	m.m[k] = v
}

func (m mutMapImpl[K, V]) Delete(k K) {
	delete(m.m, k)
}

func (m mutMapImpl[K, V]) Default(k K, v V) bool {
	if m.Contains(k) {
		return false
	}
	m.m[k] = v
	return true
}
