package containers

import (
	its "github.com/wrmsr/bane/pkg/utils/iterators"
	bt "github.com/wrmsr/bane/pkg/utils/types"
)

type OrMap[K comparable, V any] struct {
	s []bt.Kv[K, V]
	m map[K]int
}

func NewOrMap[K comparable, V any]() OrMap[K, V] {
	return OrMap[K, V]{
		s: make([]bt.Kv[K, V], 0),
		m: make(map[K]int),
	}
}

var _ MutMap[int, any] = OrMap[int, any]{}

func (m OrMap[K, V]) Contains(k K) bool {
	_, ok := m.m[k]
	return ok
}

func (m OrMap[K, V]) Get(k K) V {
	i, ok := m.m[k]
	if !ok {
		return bt.Zero[V]()
	}
	return m.s[i].V
}

func (m OrMap[K, V]) TryGet(k K) (V, bool) {
	i, ok := m.m[k]
	if !ok {
		return bt.Zero[V](), false
	}
	return m.s[i].V, true
}

func (m OrMap[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	return its.IterateSlice(m.s)
}

func (m OrMap[K, V]) Put(k K, v V) bool {
	i, ok := m.m[k]
	if !ok {
		m.s[i].V = v
		return false
	}
	m.m[k] = len(m.s)
	m.s = append(m.s, bt.KvOf(k, v))
	return true
}

func (m OrMap[K, V]) Delete(k K) bool {
	i, ok := m.m[k]
	if !ok {
		return false
	}
	m.s = DeleteAt(m.s, i)
	return true
}

func (m OrMap[K, V]) Default(k K, v V) bool {
	_, ok := m.m[k]
	if !ok {
		return false
	}
	m.m[k] = len(m.s)
	m.s = append(m.s, bt.KvOf(k, v))
	return true
}
