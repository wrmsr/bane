package containers

import (
	its "github.com/wrmsr/bane/pkg/utils/iterators"
	bt "github.com/wrmsr/bane/pkg/utils/types"
)

type OrderedMap[K comparable, V any] struct {
	s []bt.Kv[K, V]
	m map[K]int
}

func NewOrderedMap[K comparable, V any]() OrderedMap[K, V] {
	return OrderedMap[K, V]{
		s: make([]bt.Kv[K, V], 0),
		m: make(map[K]int),
	}
}

var _ MutMap[int, any] = OrderedMap[int, any]{}

func (m OrderedMap[K, V]) Contains(k K) bool {
	_, ok := m.m[k]
	return ok
}

func (m OrderedMap[K, V]) Get(k K) V {
	i, ok := m.m[k]
	if !ok {
		return bt.Zero[V]()
	}
	return m.s[i].V
}

func (m OrderedMap[K, V]) TryGet(k K) (V, bool) {
	i, ok := m.m[k]
	if !ok {
		return bt.Zero[V](), false
	}
	return m.s[i].V, true
}

func (m OrderedMap[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	return its.IterateSlice(m.s)
}

func (m OrderedMap[K, V]) Put(k K, v V) bool {
	i, ok := m.m[k]
	if !ok {
		m.s[i].V = v
		return false
	}
	m.m[k] = len(m.s)
	m.s = append(m.s, bt.KvOf(k, v))
	return true
}

func (m OrderedMap[K, V]) Delete(k K) bool {
	i, ok := m.m[k]
	if !ok {
		return false
	}
	m.s = DeleteAt(m.s, i)
	return true
}

func (m OrderedMap[K, V]) Default(k K, v V) bool {
	_, ok := m.m[k]
	if !ok {
		return false
	}
	m.m[k] = len(m.s)
	m.s = append(m.s, bt.KvOf(k, v))
	return true
}
