package container

import (
	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type StdMap[K comparable, V any] struct {
	m map[K]V
}

func NewStdMap[K comparable, V any](it bt.Iterable[bt.Kv[K, V]]) StdMap[K, V] {
	m := make(map[K]V)
	if it != nil {
		for it := it.Iterate(); it.HasNext(); {
			c := it.Next()
			m[c.K] = c.V
		}
	}
	return StdMap[K, V]{
		m: m,
	}
}

var _ Map[int, any] = StdMap[int, any]{}

func (m StdMap[K, V]) Len() int {
	return len(m.m)
}

func (m StdMap[K, V]) Contains(k K) bool {
	_, ok := m.m[k]
	return ok
}

func (m StdMap[K, V]) Get(k K) V {
	return m.m[k]
}

func (m StdMap[K, V]) TryGet(k K) (V, bool) {
	v, ok := m.m[k]
	return v, ok
}

func (m StdMap[K, V]) Iterate() bt.Iterator[bt.Kv[K, V]] {
	return its.OfMap[K, V](m.m).Iterate()
}

func (m StdMap[K, V]) ForEach(fn func(bt.Kv[K, V]) bool) bool {
	for k, v := range m.m {
		if !fn(bt.KvOf(k, v)) {
			return false
		}
	}
	return true
}

//

type MutStdMap[K comparable, V any] struct {
	m StdMap[K, V]
}

func NewMutStdMap[K comparable, V any](it bt.Iterable[bt.Kv[K, V]]) MutStdMap[K, V] {
	return MutStdMap[K, V]{m: NewStdMap[K, V](it)}
}

func WrapMap[K comparable, V any](m map[K]V) MutMap[K, V] {
	return MutStdMap[K, V]{StdMap[K, V]{m}}
}

var _ MutMap[int, any] = MutStdMap[int, any]{}

func (m MutStdMap[K, V]) isMut() {}

func (m MutStdMap[K, V]) Len() int                               { return m.m.Len() }
func (m MutStdMap[K, V]) Contains(k K) bool                      { return m.m.Contains(k) }
func (m MutStdMap[K, V]) Get(k K) V                              { return m.m.Get(k) }
func (m MutStdMap[K, V]) TryGet(k K) (V, bool)                   { return m.m.TryGet(k) }
func (m MutStdMap[K, V]) Iterate() bt.Iterator[bt.Kv[K, V]]      { return m.m.Iterate() }
func (m MutStdMap[K, V]) ForEach(fn func(bt.Kv[K, V]) bool) bool { return m.m.ForEach(fn) }

func (m MutStdMap[K, V]) Put(k K, v V) {
	m.m.m[k] = v
}

func (m MutStdMap[K, V]) Delete(k K) {
	delete(m.m.m, k)
}

func (m MutStdMap[K, V]) Default(k K, v V) bool {
	if m.Contains(k) {
		return false
	}
	m.m.m[k] = v
	return true
}

var _ Decay[Map[int, string]] = MutStdMap[int, string]{}

func (m MutStdMap[K, V]) Decay() Map[K, V] { return m.m }
