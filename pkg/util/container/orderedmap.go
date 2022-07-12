package container

import (
	its "github.com/wrmsr/bane/pkg/util/iterators"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type orderedMapImpl[K comparable, V any] struct {
	s []bt.Kv[K, V]
	m map[K]int
}

func newOrderedMapImpl[K comparable, V any](it its.Iterable[bt.Kv[K, V]]) orderedMapImpl[K, V] {
	m := orderedMapImpl[K, V]{
		m: make(map[K]int),
	}
	if it != nil {
		for it := it.Iterate(); it.HasNext(); {
			c := it.Next()
			m.put(c.K, c.V)
		}
	}
	return m
}

func NewOrderedMap[K comparable, V any](it its.Iterable[bt.Kv[K, V]]) OrderedMap[K, V] {
	return newOrderedMapImpl[K, V](it)
}

var _ OrderedMap[int, any] = orderedMapImpl[int, any]{}

func (m orderedMapImpl[K, V]) isOrdered() {}

func (m orderedMapImpl[K, V]) Len() int {
	return len(m.s)
}

func (m orderedMapImpl[K, V]) Contains(k K) bool {
	_, ok := m.m[k]
	return ok
}

func (m orderedMapImpl[K, V]) Get(k K) V {
	i, ok := m.m[k]
	if !ok {
		return bt.Zero[V]()
	}
	return m.s[i].V
}

func (m orderedMapImpl[K, V]) TryGet(k K) (V, bool) {
	i, ok := m.m[k]
	if !ok {
		return bt.Zero[V](), false
	}
	return m.s[i].V, true
}

func (m orderedMapImpl[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	return its.OfSlice(m.s).Iterate()
}

func (m orderedMapImpl[K, V]) ForEach(fn func(bt.Kv[K, V]) bool) bool {
	for _, kv := range m.s {
		if !fn(kv) {
			return false
		}
	}
	return true
}

var _ its.AnyIterable = orderedMapImpl[int, string]{}

func (m orderedMapImpl[K, V]) AnyIterate() its.Iterator[any] {
	return its.AsAny[bt.Kv[K, V]](m).Iterate()
}

func (m *orderedMapImpl[K, V]) put(k K, v V) {
	i, ok := m.m[k]
	if ok {
		m.s[i].V = v
		return
	}
	m.m[k] = len(m.s)
	m.s = append(m.s, bt.KvOf(k, v))
}

func (m *orderedMapImpl[K, V]) delete(k K) {
	i, ok := m.m[k]
	if !ok {
		return
	}
	m.s = slices.DeleteAt(m.s, i)
	delete(m.m, k)
}

func (m *orderedMapImpl[K, V]) default_(k K, v V) bool {
	_, ok := m.m[k]
	if !ok {
		return false
	}
	m.m[k] = len(m.s)
	m.s = append(m.s, bt.KvOf(k, v))
	return true
}

func (m *orderedMapImpl[K, V]) filter(fn func(kv bt.Kv[K, V]) bool) {
	for i := 0; i < len(m.s); {
		kv := m.s[i]
		if !fn(kv) {
			m.s = slices.DeleteAt(m.s, i)
			delete(m.m, kv.K)
		} else {
			i++
		}
	}
}

//

type mutOrderedMapImpl[K comparable, V any] struct {
	orderedMapImpl[K, V]
}

func NewMutOrderedMap[K comparable, V any](it its.Iterable[bt.Kv[K, V]]) MutOrderedMap[K, V] {
	return &mutOrderedMapImpl[K, V]{
		orderedMapImpl: newOrderedMapImpl[K, V](it),
	}
}

var _ MutMap[int, any] = &mutOrderedMapImpl[int, any]{}

func (m *mutOrderedMapImpl[K, V]) isMutable()       {}
func (m *mutOrderedMapImpl[K, V]) Decay() Map[K, V] { return m.orderedMapImpl }

func (m *mutOrderedMapImpl[K, V]) Put(k K, v V) {
	m.put(k, v)
}

func (m *mutOrderedMapImpl[K, V]) Delete(k K) {
	m.delete(k)
}

func (m *mutOrderedMapImpl[K, V]) Default(k K, v V) bool {
	return m.default_(k, v)
}
