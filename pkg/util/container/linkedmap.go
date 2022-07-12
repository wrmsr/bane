package container

import (
	its "github.com/wrmsr/bane/pkg/util/iterators"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type LinkedMap[K comparable, V any] struct {
	s []bt.Kv[K, V]
	m map[K]int
}

func NewLinkedMap[K comparable, V any](it its.Iterable[bt.Kv[K, V]]) LinkedMap[K, V] {
	m := LinkedMap[K, V]{
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

var _ OrderedMap[int, any] = LinkedMap[int, any]{}

func (m LinkedMap[K, V]) isOrdered() {}

func (m LinkedMap[K, V]) Len() int {
	return len(m.s)
}

func (m LinkedMap[K, V]) Contains(k K) bool {
	_, ok := m.m[k]
	return ok
}

func (m LinkedMap[K, V]) Get(k K) V {
	i, ok := m.m[k]
	if !ok {
		return bt.Zero[V]()
	}
	return m.s[i].V
}

func (m LinkedMap[K, V]) TryGet(k K) (V, bool) {
	i, ok := m.m[k]
	if !ok {
		return bt.Zero[V](), false
	}
	return m.s[i].V, true
}

func (m LinkedMap[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	return its.OfSlice(m.s).Iterate()
}

func (m LinkedMap[K, V]) ForEach(fn func(bt.Kv[K, V]) bool) bool {
	for _, kv := range m.s {
		if !fn(kv) {
			return false
		}
	}
	return true
}

var _ its.AnyIterable = LinkedMap[int, string]{}

func (m LinkedMap[K, V]) AnyIterate() its.Iterator[any] {
	return its.AsAny[bt.Kv[K, V]](m).Iterate()
}

func (m *LinkedMap[K, V]) put(k K, v V) {
	i, ok := m.m[k]
	if ok {
		m.s[i].V = v
		return
	}
	m.m[k] = len(m.s)
	m.s = append(m.s, bt.KvOf(k, v))
}

func (m *LinkedMap[K, V]) delete(k K) {
	i, ok := m.m[k]
	if !ok {
		return
	}
	m.s = slices.DeleteAt(m.s, i)
	delete(m.m, k)
}

func (m *LinkedMap[K, V]) default_(k K, v V) bool {
	_, ok := m.m[k]
	if !ok {
		return false
	}
	m.m[k] = len(m.s)
	m.s = append(m.s, bt.KvOf(k, v))
	return true
}

func (m *LinkedMap[K, V]) filter(fn func(kv bt.Kv[K, V]) bool) {
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

type MutLinkedMap[K comparable, V any] struct {
	m LinkedMap[K, V]
}

func NewMutLinkedMap[K comparable, V any](it its.Iterable[bt.Kv[K, V]]) *MutLinkedMap[K, V] {
	return &MutLinkedMap[K, V]{m: NewLinkedMap[K, V](it)}
}

var _ MutMap[int, any] = &MutLinkedMap[int, any]{}

func (m *MutLinkedMap[K, V]) isMutable() {}
func (m *MutLinkedMap[K, V]) isOrdered() {}

func (m *MutLinkedMap[K, V]) Len() int                                 { return m.m.Len() }
func (m *MutLinkedMap[K, V]) Contains(k K) bool                        { return m.m.Contains(k) }
func (m *MutLinkedMap[K, V]) Get(k K) V                                { return m.m.Get(k) }
func (m *MutLinkedMap[K, V]) TryGet(k K) (V, bool)                     { return m.m.TryGet(k) }
func (m *MutLinkedMap[K, V]) Iterate() its.Iterator[bt.Kv[K, V]]       { return m.m.Iterate() }
func (m *MutLinkedMap[K, V]) ForEach(fn func(v bt.Kv[K, V]) bool) bool { return m.m.ForEach(fn) }

func (m *MutLinkedMap[K, V]) Put(k K, v V) {
	m.m.put(k, v)
}

func (m *MutLinkedMap[K, V]) Delete(k K) {
	m.m.delete(k)
}

func (m *MutLinkedMap[K, V]) Default(k K, v V) bool {
	return m.m.default_(k, v)
}

var _ Decay[Map[int, string]] = &MutLinkedMap[int, string]{}

func (m *MutLinkedMap[K, V]) Decay() Map[K, V] { return m.m }
