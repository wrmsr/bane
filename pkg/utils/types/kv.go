package types

type Kv[K comparable, V any] struct {
	K K
	V V
}

func KvOf[K comparable, V any](k K, v V) Kv[K, V] {
	return Kv[K, V]{k, v}
}
