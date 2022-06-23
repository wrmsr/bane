package iterators

import bt "github.com/wrmsr/bane/pkg/util/types"

//

func Seq[T any](it Iterable[T]) []T {
	var r []T
	for i := it.Iterate(); i.HasNext(); {
		r = append(r, i.Next())
	}
	return r
}

func Take[T any](it Iterable[T], n int) []T {
	var r []T
	c := 0
	for i := it.Iterate(); i.HasNext(); {
		if c >= n {
			break
		}
		r = append(r, i.Next())
		c++
	}
	return r
}

//

func Keys[K comparable, V any](it Iterable[bt.Kv[K, V]]) Iterable[K] {
	return Map(it, func(kv bt.Kv[K, V]) K { return kv.K })
}

func Values[K comparable, V any](it Iterable[bt.Kv[K, V]]) Iterable[V] {
	return Map(it, func(kv bt.Kv[K, V]) V { return kv.V })
}

//

type NonUniqueError[T any] struct {
	Value T
}

func CheckUnique[T comparable](it Iterable[T]) Iterable[T] {
	return Factory(func() Iterator[T] {
		s := make(map[T]struct{})
		return Map(it, func(v T) T {
			if _, ok := s[v]; ok {
				panic(NonUniqueError[T]{v})
			}
			return v
		}).Iterate()
	})
}

func CheckUniqueKeys[K comparable, V any](it Iterable[bt.Kv[K, V]]) Iterable[bt.Kv[K, V]] {
	return Factory(func() Iterator[bt.Kv[K, V]] {
		s := make(map[K]struct{})
		return Map(it, func(kv bt.Kv[K, V]) bt.Kv[K, V] {
			if _, ok := s[kv.K]; ok {
				panic(NonUniqueError[bt.Kv[K, V]]{kv})
			}
			return kv
		}).Iterate()
	})
}
