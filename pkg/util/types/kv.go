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

func KvsOf[K, V any](o ...any) []Kv[K, V] {
	if len(o)%2 != 0 {
		panic("must pass even number of args")
	}
	n := len(o) / 2
	r := make([]Kv[K, V], n)
	for i, j := 0, 0; i < n; i, j = i+1, j+2 {
		r[i] = KvOf(o[j].(K), o[j+1].(V))
	}
	return r
}
