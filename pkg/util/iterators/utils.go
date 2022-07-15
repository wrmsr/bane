package iterators

import bt "github.com/wrmsr/bane/pkg/util/types"

//

func Seq[T any](it bt.Iterable[T]) []T {
	var r []T
	for i := it.Iterate(); i.HasNext(); {
		r = append(r, i.Next())
	}
	return r
}

func Take[T any](it bt.Iterable[T], n int) []T {
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

func Apply[T any](fn func(v T), it bt.Iterable[T]) {
	for it := it.Iterate(); it.HasNext(); {
		fn(it.Next())
	}
}

func Len[T any](it bt.Iterable[T]) int {
	i := 0
	for it := it.Iterate(); it.HasNext(); it.Next() {
		i++
	}
	return i
}

func Nth[T any](it bt.Iterable[T], n int) (T, bool) {
	if n < 0 {
		panic(n)
	}
	i := it.Iterate()
	for ; ; n-- {
		if !i.HasNext() {
			var z T
			return z, false
		}
		if n == 0 {
			return i.Next(), true
		}
		i.Next()
	}
}

//

type NonUniqueError[T any] struct {
	Value T
}

func CheckUnique[T comparable](it bt.Iterable[T]) bt.Iterable[T] {
	return Factory(func() bt.Iterator[T] {
		s := make(map[T]struct{})
		return Map(it, func(v T) T {
			if _, ok := s[v]; ok {
				panic(NonUniqueError[T]{v})
			}
			return v
		}).Iterate()
	}, it)
}

func CheckUniqueKeys[K comparable, V any](it bt.Iterable[bt.Kv[K, V]]) bt.Iterable[bt.Kv[K, V]] {
	return Factory(func() bt.Iterator[bt.Kv[K, V]] {
		s := make(map[K]struct{})
		return Map(it, func(kv bt.Kv[K, V]) bt.Kv[K, V] {
			if _, ok := s[kv.K]; ok {
				panic(NonUniqueError[bt.Kv[K, V]]{kv})
			}
			return kv
		}).Iterate()
	}, it)
}

//

func Any[T any](fn func(v T) bool, it bt.Iterable[T]) bool {
	for it := it.Iterate(); it.HasNext(); {
		if fn(it.Next()) {
			return true
		}
	}
	return false
}
