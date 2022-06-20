package types

type Kv[K comparable, V any] struct {
	K K
	V V
}

func (kv Kv[K, V]) GetK() K { return kv.K }
func (kv Kv[K, V]) GetV() V { return kv.V }

func KvOf[K comparable, V any](k K, v V) Kv[K, V] {
	return Kv[K, V]{k, v}
}
