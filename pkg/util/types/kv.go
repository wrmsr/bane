package types

import "reflect"

//

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

//

func AsKey[K, V any](o any) K {
	if kv, ok := o.(Kv[K, V]); ok {
		return kv.K
	}
	return o.(K)
}

func KeyCmpImpl[K, V any](cmp CmpImpl[K]) CmpImpl[Kv[K, V]] {
	return func(r, l Kv[K, V]) CmpResult {
		return cmp(l.K, r.K)
	}
}

//

func (kv Kv[K, V]) ReflectTypeArgs() []reflect.Type {
	return []reflect.Type{reflect.TypeOf(kv.K), reflect.TypeOf(kv.V)}
}
