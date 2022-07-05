package types

type Kv[K, V any] struct {
	K K
	V V
}

func (kv Kv[K, V]) GetK() K { return kv.K }
func (kv Kv[K, V]) GetV() V { return kv.V }

func (kv Kv[K, V]) Unpack() (K, V) { return kv.K, kv.V }

func KvOf[K, V any](k K, v V) Kv[K, V] {
	return Kv[K, V]{k, v}
}

func KvMaker[K, V any](fn func(v V) K) func(v V) Kv[K, V] {
	return func(v V) Kv[K, V] {
		return Kv[K, V]{K: fn(v), V: v}
	}
}
