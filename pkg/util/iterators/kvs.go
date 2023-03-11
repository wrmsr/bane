package iterators

import (
	"reflect"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

func Keys[K, V any](it bt.Iterable[bt.Kv[K, V]]) bt.Iterable[K] {
	return Map(it, func(kv bt.Kv[K, V]) K { return kv.K })
}

func Values[K, V any](it bt.Iterable[bt.Kv[K, V]]) bt.Iterable[V] {
	return Map(it, func(kv bt.Kv[K, V]) V { return kv.V })
}

func MapKeys[KF, KT, V any](it bt.Iterable[bt.Kv[KF, V]], fn func(k KF) KT) bt.Iterable[bt.Kv[KT, V]] {
	return Map(it, func(kv bt.Kv[KF, V]) bt.Kv[KT, V] {
		return bt.KvOf(fn(kv.K), kv.V)
	})
}

func MapValues[K, VF, VT any](it bt.Iterable[bt.Kv[K, VF]], fn func(v VF) VT) bt.Iterable[bt.Kv[K, VT]] {
	return Map(it, func(kv bt.Kv[K, VF]) bt.Kv[K, VT] {
		return bt.KvOf(kv.K, fn(kv.V))
	})
}

func FilterKeys[K, V any](it bt.Iterable[bt.Kv[K, V]], fn func(k K) bool) bt.Iterable[bt.Kv[K, V]] {
	return Filter(it, func(kv bt.Kv[K, V]) bool {
		return fn(kv.K)
	})
}

func FilterValues[K, V any](it bt.Iterable[bt.Kv[K, V]], fn func(v V) bool) bt.Iterable[bt.Kv[K, V]] {
	return Filter(it, func(kv bt.Kv[K, V]) bool {
		return fn(kv.V)
	})
}

func MakeKvs[K, V any](it bt.Iterable[V], fn func(v V) K) bt.Iterable[bt.Kv[K, V]] {
	return Map(it, func(v V) bt.Kv[K, V] {
		return bt.KvOf(fn(v), v)
	})
}

func StubKvs[K any](it bt.Iterable[K]) bt.Iterable[bt.Kv[K, struct{}]] {
	return Map[K, bt.Kv[K, struct{}]](it, func(k K) bt.Kv[K, struct{}] { return bt.KvOf(k, struct{}{}) })
}

func MakeTypeKvs[V any](it bt.Iterable[V]) bt.Iterable[bt.Kv[reflect.Type, V]] {
	return MakeKvs(it, func(v V) reflect.Type { return reflect.TypeOf(v) })
}
