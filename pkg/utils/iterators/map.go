package iterators

import bt "github.com/wrmsr/bane/pkg/utils/types"

func IterateMap[K comparable, V any](m map[K]V) Iterator[bt.Kv[K, V]] {
	kvs := make([]bt.Kv[K, V], len(m))
	i := 0
	for k, v := range m {
		kvs[i] = bt.Kv[K, V]{k, v}
		i++
	}
	return IterateSlice[bt.Kv[K, V]](kvs)
}
