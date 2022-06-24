package iterators

import bt "github.com/wrmsr/bane/pkg/util/types"

func OfMap[K comparable, V any](m map[K]V) Iterable[bt.Kv[K, V]] {
	return Factory(func() Iterator[bt.Kv[K, V]] {
		kvs := make([]bt.Kv[K, V], len(m))
		i := 0
		for k, v := range m {
			kvs[i] = bt.KvOf(k, v)
			i++
		}
		return OfSlice[bt.Kv[K, V]](kvs).Iterate()
	}, m)
}
