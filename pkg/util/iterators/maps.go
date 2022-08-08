package iterators

import (
	"reflect"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

func OfMap[K comparable, V any](m map[K]V) bt.Iterable[bt.Kv[K, V]] {
	return Factory(func() bt.Iterator[bt.Kv[K, V]] {
		kvs := make([]bt.Kv[K, V], len(m))
		i := 0
		for k, v := range m {
			kvs[i] = bt.KvOf(k, v)
			i++
		}
		return OfSlice[bt.Kv[K, V]](kvs).Iterate()
	}, m)
}

func OfTypeMap[V any](m map[reflect.Type]V) bt.Iterable[bt.Kv[reflect.Type, V]] {
	return Factory(func() bt.Iterator[bt.Kv[reflect.Type, V]] {
		kvs := make([]bt.Kv[reflect.Type, V], len(m))
		i := 0
		for k, v := range m {
			kvs[i] = bt.KvOf(k, v)
			i++
		}
		return OfSlice[bt.Kv[reflect.Type, V]](kvs).Iterate()
	}, m)
}
