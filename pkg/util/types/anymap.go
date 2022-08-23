package types

//

type AnyMap interface {
	Traversable[Kv[any, any]]

	Len() int
	Contains(k any) bool
	Get(k any) any
	TryGet(k any) (any, bool)
}

type MutAnyMap interface {
	AnyMap

	Put(k, v any)
	Delete(k any)
	Default(k, v any) bool
}

//

type AnyMapImpl[K comparable, V any] struct {
	m map[K]V
}

func AnyMapOf[K comparable, V any](m map[K]V) AnyMapImpl[K, V] {
	return AnyMapImpl[K, V]{m: m}
}

var _ AnyMap = AnyMapImpl[int, string]{}

func (a AnyMapImpl[K, V]) Len() int {
	return len(a.m)
}

func (a AnyMapImpl[K, V]) Contains(k any) bool {
	_, ok := a.m[k.(K)]
	return ok
}

func (a AnyMapImpl[K, V]) Get(k any) any {
	return a.m[k.(K)]
}

func (a AnyMapImpl[K, V]) TryGet(k any) (any, bool) {
	v, ok := a.m[k.(K)]
	return v, ok
}

func (a AnyMapImpl[K, V]) ForEach(fn func(kv Kv[any, any]) bool) bool {
	for k, v := range a.m {
		if !fn(KvOf[any, any](k, v)) {
			return false
		}
	}
	return true
}

//

type MutAnyMapImpl[K comparable, V any] struct {
	AnyMapImpl[K, V]
}

func MutAnyMapOf[K comparable, V any](m map[K]V) MutAnyMapImpl[K, V] {
	return MutAnyMapImpl[K, V]{AnyMapImpl[K, V]{m: m}}
}

var _ MutAnyMap = MutAnyMapImpl[int, string]{}

func (m MutAnyMapImpl[K, V]) Put(k, v any) {
	m.m[k.(K)] = v.(V)
}

func (m MutAnyMapImpl[K, V]) Delete(k any) {
	delete(m.m, k.(K))
}

func (m MutAnyMapImpl[K, V]) Default(k, v any) bool {
	k_ := k.(K)
	if _, ok := m.m[k_]; ok {
		return false
	}
	m.m[k_] = v.(V)
	return true
}
