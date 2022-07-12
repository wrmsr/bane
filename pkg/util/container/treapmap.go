package container

import (
	"github.com/wrmsr/bane/pkg/util/container/treap"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	rndu "github.com/wrmsr/bane/pkg/util/rand"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

type treapMapComparer[K, V any] func(kv1, kv2 bt.Kv[K, V]) int

func (c treapMapComparer[K, V]) Compare(kv1, kv2 bt.Kv[K, V]) int {
	return c(kv1, kv2)
}

type TreapMap[K, V any] struct {
	n *treap.TreapNode[bt.Kv[K, V]]
	c treap.Comparer[bt.Kv[K, V]]
}

func NewTreapMap[K, V any](cmp treap.Comparer[K]) TreapMap[K, V] {
	return TreapMap[K, V]{
		c: treapMapComparer[K, V](func(v1, v2 bt.Kv[K, V]) int {
			return cmp.Compare(v1.K, v2.K)
		}),
	}
}

var _ PersistentMap[int, string] = TreapMap[int, string]{}

func (m TreapMap[K, V]) isPersistent() {}

func (m TreapMap[K, V]) Len() int {
	// TODO: memo
	result := 0
	m.n.ForEach(func(_ bt.Kv[K, V]) bool {
		result++
		return true
	})
	return result
}

func (m TreapMap[K, V]) Contains(k K) bool {
	_, ok := m.TryGet(k)
	return ok
}

func (m TreapMap[K, V]) Get(k K) V {
	if v, ok := m.TryGet(k); ok {
		return v
	}
	var z V
	return z
}

func (m TreapMap[K, V]) TryGet(k K) (V, bool) {
	n := m.n.Find(bt.KvOf[K, V](k, bt.Zero[V]()), m.c)
	if n == nil {
		var z V
		return z, false
	}
	return n.Value.V, true
}

func (m TreapMap[K, V]) Iterate() its.Iterator[bt.Kv[K, V]] {
	panic("implement me")
}

func (m TreapMap[K, V]) ForEach(fn func(kv bt.Kv[K, V]) bool) bool {
	return m.n.ForEach(func(kv bt.Kv[K, V]) bool {
		return fn(kv)
	})
}

func (m TreapMap[K, V]) With(k K, v V) PersistentMap[K, V] {
	node := &treap.TreapNode[bt.Kv[K, V]]{
		bt.KvOf(k, v),
		int(rndu.FastUint32()),
		nil,
		nil,
	}
	n := m.n.Union(node, m.c, true)
	return TreapMap[K, V]{n, m.c}
}

func (m TreapMap[K, V]) Without(k K) PersistentMap[K, V] {
	n := m.n.Delete(bt.KvOf[K, V](k, bt.Zero[V]()), m.c)
	return TreapMap[K, V]{n, m.c}
}

func (m TreapMap[K, V]) Default(k K, v V) PersistentMap[K, V] {
	if _, ok := m.TryGet(k); ok {
		return m
	}
	return m.With(k, v)
}
