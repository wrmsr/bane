package iterators

import bt "github.com/wrmsr/bane/pkg/utils/types"

func Map[K comparable, V any](m map[K]V) Iterable[bt.Kv[K, V]] {
	return factory(func() Iterator[bt.Kv[K, V]] {
		kvs := make([]bt.Kv[K, V], len(m))
		i := 0
		for k, v := range m {
			kvs[i] = bt.KvOf(k, v)
			i++
		}
		return Slice[bt.Kv[K, V]](kvs).Iterate()
	})
}
