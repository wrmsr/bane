package container

import (
	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

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

type biMapImpl[K, V comparable] struct {
	Map[K, V]
	i Map[V, K]
}

func initBiMap[K, V comparable](it its.Iterable[bt.Kv[K, V]]) (map[K]V, map[V]K) {
	m := make(map[K]V)
	i := make(map[V]K)
	if it != nil {
		for it := it.Iterate(); it.HasNext(); {
			kv := it.Next()
			m[kv.K] = kv.V
			i[kv.V] = kv.K
		}
	}
	return m, i
}

func NewBiMap[K, V comparable](it its.Iterable[bt.Kv[K, V]]) BiMap[K, V] {
	m, i := initBiMap(it)
	return biMapImpl[K, V]{Map: mapImpl[K, V]{m: m}, i: mapImpl[V, K]{i}}
}

var _ BiMap[int, string] = biMapImpl[int, string]{}

func (m biMapImpl[K, V]) Invert() BiMap[V, K] {
	var r any = biMapImpl[V, K]{Map: m.i, i: m.Map}
	return r.(BiMap[V, K]) // FIXME: goland bug
}

//

type mutBiMapImpl[K, V comparable] struct {
	m MutMap[K, V]
	i MutMap[V, K]
}

func NewMutBiMap[K, V comparable](it its.Iterable[bt.Kv[K, V]]) MutBiMap[K, V] {
	m, i := initBiMap(it)
	return mutBiMapImpl[K, V]{m: mutMapImpl[K, V]{mapImpl[K, V]{m: m}}, i: mutMapImpl[V, K]{mapImpl[V, K]{i}}}
}

var _ MutBiMap[int, string] = mutBiMapImpl[int, string]{}

func (m mutBiMapImpl[K, V]) Len() int {
	return m.m.Len()
}

func (m mutBiMapImpl[K, V]) Contains(k K) bool {
	return m.m.Contains(k)
}

func (m mutBiMapImpl[K, V]) Get(k K) V {
	return m.m.Get(k)
}

func (m mutBiMapImpl[K, V]) TryGet(k K) (V, bool) {
	return m.m.TryGet(k)
}

func (m mutBiMapImpl[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	return m.m.Iterate()
}

func (m mutBiMapImpl[K, V]) ForEach(fn func(kv bt.Kv[K, V]) bool) bool {
	return m.m.ForEach(fn)
}

func (m mutBiMapImpl[K, V]) isMutable() {}

func (m mutBiMapImpl[K, V]) Put(k K, v V) {
	if o, ok := m.m.TryGet(k); ok {
		m.m.Delete(k)
		m.i.Delete(o)
	}
	if o, ok := m.i.TryGet(v); ok {
		m.m.Delete(o)
		m.i.Delete(v)
	}
	m.m.Put(k, v)
	m.i.Put(v, k)
}

func (m mutBiMapImpl[K, V]) Delete(k K) {
	if o, ok := m.TryGet(k); ok {
		m.m.Delete(k)
		m.i.Delete(o)
	}
}

func (m mutBiMapImpl[K, V]) Default(k K, v V) bool {
	if m.Contains(k) {
		return false
	}
	m.Put(k, v)
	return true
}

func (m mutBiMapImpl[K, V]) Invert() BiMap[V, K] {
	return m.MutInvert()
}

func (m mutBiMapImpl[K, V]) MutInvert() MutBiMap[V, K] {
	var r any = mutBiMapImpl[V, K]{m: m.i, i: m.m}
	return r.(MutBiMap[V, K]) // FIXME: goland bug
}
