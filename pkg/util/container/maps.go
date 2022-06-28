package container

import its "github.com/wrmsr/bane/pkg/util/iterators"

func Keys[K, V any](m Map[K, V]) its.Iterable[K] {
	return its.Keys[K, V](m)
}

func Values[K, V any](m Map[K, V]) its.Iterable[V] {
	return its.Values[K, V](m)
}
