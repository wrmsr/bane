package container

import (
	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func Keys[K, V any](m Map[K, V]) its.Iterable[K] {
	return its.Keys[K, V](m)
}

func Values[K, V any](m Map[K, V]) its.Iterable[V] {
	return its.Values[K, V](m)
}

func MapKeys[KF, KT comparable, V any](fn func(k KF) KT, m Map[KF, V]) Map[KT, V] {
	r := NewMutMap[KT, V](nil)
	m.ForEach(func(kv bt.Kv[KF, V]) bool {
		r.Put(fn(kv.K), kv.V)
		return true
	})
	return r
}

func MapValues[K comparable, VF, VT any](fn func(v VF) VT, m Map[K, VF]) Map[K, VT] {
	r := NewMutMap[K, VT](nil)
	m.ForEach(func(kv bt.Kv[K, VF]) bool {
		r.Put(kv.K, fn(kv.V))
		return true
	})
	return r
}
