/*
See also containers/mapbuilder.go
*/
package maps

import bt "github.com/wrmsr/bane/core/types"

type Builder[K comparable, V any] struct {
	m map[K]V
}

func NewBuilder[K comparable, V any](m map[K]V) *Builder[K, V] {
	if m == nil {
		m = make(map[K]V)
	}
	return &Builder[K, V]{m: m}
}

func (b *Builder[K, V]) Put(k K, v V) *Builder[K, V] {
	b.m[k] = v
	return b
}

func (b *Builder[K, V]) Update(m map[K]V) *Builder[K, V] {
	if m != nil {
		for k, v := range m {
			b.m[k] = v
		}
	}
	return b
}

func (b *Builder[K, V]) UpdateKv(it bt.Iterable[bt.Kv[K, V]]) *Builder[K, V] {
	if it != nil {
		for it := it.Iterate(); it.HasNext(); {
			kv := it.Next()
			b.m[kv.K] = kv.V
		}
	}
	return b
}

func (b *Builder[K, V]) Filter(fn func(k K, v V) bool) *Builder[K, V] {
	for k, v := range b.m {
		if !fn(k, v) {
			delete(b.m, k)
		}
	}
	return b
}

func (b *Builder[K, V]) FilterKv(fn func(kv bt.Kv[K, V]) bool) *Builder[K, V] {
	b.Filter(func(k K, v V) bool { return fn(bt.KvOf(k, v)) })
	return b
}

func (b *Builder[K, V]) FilterKeys(fn func(k K) bool) *Builder[K, V] {
	b.Filter(func(k K, v V) bool { return fn(k) })
	return b
}

func (b *Builder[K, V]) FilterValues(fn func(v V) bool) *Builder[K, V] {
	b.Filter(func(k K, v V) bool { return fn(v) })
	return b
}

func (b *Builder[K, V]) Delete(k K) *Builder[K, V] {
	delete(b.m, k)
	return b
}

func (b *Builder[K, V]) Default(k K, v V) *Builder[K, V] {
	if _, ok := b.m[k]; !ok {
		b.m[k] = v
	}
	return b
}

func (b *Builder[K, V]) Build() map[K]V {
	r := b.m
	b.m = nil
	return r
}
