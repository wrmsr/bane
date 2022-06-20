package containers

import (
	its "github.com/wrmsr/bane/pkg/utils/iterators"
	bt "github.com/wrmsr/bane/pkg/utils/types"
)

type OrdMap[K comparable, V any] struct {
	s []bt.Kv[K, V]
	m map[K]int
}

func NewOrdMap[K comparable, V any]() OrdMap[K, V] {
	return OrdMap[K, V]{
		s: make([]bt.Kv[K, V], 0),
		m: make(map[K]int),
	}
}

var _ MutMap[int, any] = OrdMap[int, any]{}

func (m OrdMap[K, V]) Contains(k K) bool {
	_, ok := m.m[k]
	return ok
}

func (m OrdMap[K, V]) Get(k K) V {
	i, ok := m.m[k]
	if !ok {
		return bt.Zero[V]()
	}
	return m.s[i].V
}

func (m OrdMap[K, V]) TryGet(k K) (V, bool) {
	i, ok := m.m[k]
	if !ok {
		return bt.Zero[V](), false
	}
	return m.s[i].V, true
}

func (m OrdMap[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	return its.Slice(m.s)
}

func (m OrdMap[K, V]) Put(k K, v V) bool {
	i, ok := m.m[k]
	if !ok {
		m.s[i].V = v
		return false
	}
	m.m[k] = len(m.s)
	m.s = append(m.s, bt.KvOf(k, v))
	return true
}

func (m OrdMap[K, V]) Delete(k K) bool {
	i, ok := m.m[k]
	if !ok {
		return false
	}
	m.s = DeleteAt(m.s, i)
	return true
}

func (m OrdMap[K, V]) Default(k K, v V) bool {
	_, ok := m.m[k]
	if !ok {
		return false
	}
	m.m[k] = len(m.s)
	m.s = append(m.s, bt.KvOf(k, v))
	return true
}
