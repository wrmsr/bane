package iterators

import bt "github.com/wrmsr/bane/pkg/util/types"

func Keys[K, V any](it Iterable[bt.Kv[K, V]]) Iterable[K] {
	return Map(it, func(kv bt.Kv[K, V]) K { return kv.K })
}

func Values[K, V any](it Iterable[bt.Kv[K, V]]) Iterable[V] {
	return Map(it, func(kv bt.Kv[K, V]) V { return kv.V })
}

func MapKeys[KF, KT, V any](it Iterable[bt.Kv[KF, V]], fn func(k KF) KT) Iterable[bt.Kv[KT, V]] {
	return Map(it, func(kv bt.Kv[KF, V]) bt.Kv[KT, V] {
		return bt.KvOf(fn(kv.K), kv.V)
	})
}

func MapValues[K, VF, VT any](it Iterable[bt.Kv[K, VF]], fn func(v VF) VT) Iterable[bt.Kv[K, VT]] {
	return Map(it, func(kv bt.Kv[K, VF]) bt.Kv[K, VT] {
		return bt.KvOf(kv.K, fn(kv.V))
	})
}