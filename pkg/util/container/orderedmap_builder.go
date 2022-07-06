package container

import (
	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

type OrderedMapBuilder[K comparable, V any] struct {
	m *orderedMapImpl[K, V]
}

func NewOrderedMapBuilder[K comparable, V any]() *OrderedMapBuilder[K, V] {
	return &OrderedMapBuilder[K, V]{
		m: &orderedMapImpl[K, V]{
			m: make(map[K]int),
		},
	}
}

func (b *OrderedMapBuilder[K, V]) Put(k K, v V) *OrderedMapBuilder[K, V] {
	b.m.put(k, v)
	return b
}

func (b *OrderedMapBuilder[K, V]) Update(it its.Iterable[bt.Kv[K, V]]) *OrderedMapBuilder[K, V] {
	if it != nil {
		for it := it.Iterate(); it.HasNext(); {
			kv := it.Next()
			b.m.put(kv.K, kv.V)
		}
	}
	return b
}

func (b *OrderedMapBuilder[K, V]) Filter(fn func(kv bt.Kv[K, V]) bool) *OrderedMapBuilder[K, V] {
	b.m.filter(fn)
	return b
}

func (b *OrderedMapBuilder[K, V]) FilterKeys(fn func(k K) bool) *OrderedMapBuilder[K, V] {
	b.m.filter(func(kv bt.Kv[K, V]) bool { return fn(kv.K) })
	return b
}

func (b *OrderedMapBuilder[K, V]) FilterValues(fn func(v V) bool) *OrderedMapBuilder[K, V] {
	b.m.filter(func(kv bt.Kv[K, V]) bool { return fn(kv.V) })
	return b
}

func (b *OrderedMapBuilder[K, V]) Delete(k K) *OrderedMapBuilder[K, V] {
	b.m.delete(k)
	return b
}

func (b *OrderedMapBuilder[K, V]) Default(k K, v V) *OrderedMapBuilder[K, V] {
	b.m.default_(k, v)
	return b
}

func (b *OrderedMapBuilder[K, V]) Build() OrderedMap[K, V] {
	r := b.m
	b.m = nil
	return r
}
