package containers

import (
	its "github.com/wrmsr/bane/pkg/utils/iterators"
	bt "github.com/wrmsr/bane/pkg/utils/types"
)

//

type OrdMap[K comparable, V any] interface {
	Map[K, V]
	Ordered
}

type MutOrdMap[K comparable, V any] interface {
	OrdMap[K, V]
	MutMap[K, V]
}

//

type ordMapImpl[K comparable, V any] struct {
	s []bt.Kv[K, V]
	m map[K]int
}

func newOrdMapImpl[K comparable, V any]() ordMapImpl[K, V] {
	return ordMapImpl[K, V]{
		s: make([]bt.Kv[K, V], 0),
		m: make(map[K]int),
	}
}

func NewOrdMap[K comparable, V any]() OrdMap[K, V] {
	return newOrdMapImpl[K, V]()
}

var _ OrdMap[int, any] = ordMapImpl[int, any]{}

func (m ordMapImpl[K, V]) isOrdered() {}

func (m ordMapImpl[K, V]) Len() int {
	return len(m.s)
}

func (m ordMapImpl[K, V]) Contains(k K) bool {
	_, ok := m.m[k]
	return ok
}

func (m ordMapImpl[K, V]) Get(k K) V {
	i, ok := m.m[k]
	if !ok {
		return bt.Zero[V]()
	}
	return m.s[i].V
}

func (m ordMapImpl[K, V]) TryGet(k K) (V, bool) {
	i, ok := m.m[k]
	if !ok {
		return bt.Zero[V](), false
	}
	return m.s[i].V, true
}

func (m ordMapImpl[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	return its.Slice(m.s).Iterate()
}

func (m ordMapImpl[K, V]) ForEach(fn func(bt.Kv[K, V]) bool) bool {
	for _, kv := range m.s {
		if !fn(kv) {
			return false
		}
	}
	return true
}

//

type mutOrdMapImpl[K comparable, V any] struct {
	ordMapImpl[K, V]
}

func NewMutOrdMap[K comparable, V any]() MutOrdMap[K, V] {
	return &mutOrdMapImpl[K, V]{
		ordMapImpl: newOrdMapImpl[K, V](),
	}
}

var _ MutMap[int, any] = &mutOrdMapImpl[int, any]{}

func (m *ordMapImpl[K, V]) isMutable() {}

func (m *ordMapImpl[K, V]) Put(k K, v V) bool {
	i, ok := m.m[k]
	if !ok {
		m.s[i].V = v
		return false
	}
	m.m[k] = len(m.s)
	m.s = append(m.s, bt.KvOf(k, v))
	return true
}

func (m *ordMapImpl[K, V]) Delete(k K) bool {
	i, ok := m.m[k]
	if !ok {
		return false
	}
	m.s = DeleteAt(m.s, i)
	return true
}

func (m *ordMapImpl[K, V]) Default(k K, v V) bool {
	_, ok := m.m[k]
	if !ok {
		return false
	}
	m.m[k] = len(m.s)
	m.s = append(m.s, bt.KvOf(k, v))
	return true
}
