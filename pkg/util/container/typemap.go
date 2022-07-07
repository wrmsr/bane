package container

import (
	"reflect"

	its "github.com/wrmsr/bane/pkg/util/iterators"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

//

type TypeMap[T any] interface {
	Len() int
	Contains(ty reflect.Type) bool
	Get(ty reflect.Type) T
	TryGet(ty reflect.Type) (T, bool)

	its.Iterable[T]
	its.Traversable[T]

	rfl.TypeArgsReflector
}

type MutTypeMap[T any] interface {
	TypeMap[T]
	Mutable

	Put(v T)
	Delete(ty reflect.Type)
	Default(v T) bool
}

//

type typeMapImpl[T any] struct {
	m map[reflect.Type]T
}

func newTypeMapImpl[T any](it its.Iterable[T]) typeMapImpl[T] {
	m := make(map[reflect.Type]T)
	if it != nil {
		for it := it.Iterate(); it.HasNext(); {
			v := it.Next()
			m[reflect.TypeOf(v)] = v
		}
	}
	return typeMapImpl[T]{m: m}
}

func NewTypeMap[T any](it its.Iterable[T]) TypeMap[T] {
	return newTypeMapImpl(it)
}

var _ TypeMap[any] = typeMapImpl[any]{}

func (m typeMapImpl[T]) ReflectTypeArgs() []reflect.Type {
	return []reflect.Type{rfl.TypeOf[T]()}
}

func (m typeMapImpl[T]) Len() int {
	return len(m.m)
}

func (m typeMapImpl[T]) Contains(ty reflect.Type) bool {
	_, ok := m.m[ty]
	return ok
}

func (m typeMapImpl[T]) Get(ty reflect.Type) T {
	v, _ := m.m[ty]
	return v
}

func (m typeMapImpl[T]) TryGet(ty reflect.Type) (T, bool) {
	v, ok := m.m[ty]
	return v, ok
}

func (m typeMapImpl[T]) Iterate() its.Iterator[T] {
	vs := make([]T, len(m.m))
	i := 0
	for _, v := range m.m {
		vs[i] = v
		i++
	}
	return its.OfSlice(vs).Iterate()
}

func (m typeMapImpl[T]) ForEach(fn func(v T) bool) bool {
	for _, v := range m.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

//

type mutTypeMapImpl[T any] struct {
	typeMapImpl[T]
}

func NewMutTypeMap[T any](it its.Iterable[T]) MutTypeMap[T] {
	return mutTypeMapImpl[T]{newTypeMapImpl(it)}
}

var _ MutTypeMap[any] = mutTypeMapImpl[any]{}

func (m mutTypeMapImpl[T]) isMutable() {}

func (m mutTypeMapImpl[T]) Put(v T) {
	m.m[reflect.TypeOf(v)] = v
}

func (m mutTypeMapImpl[T]) Delete(ty reflect.Type) {
	delete(m.m, ty)
}

func (m mutTypeMapImpl[T]) Default(v T) bool {
	ty := reflect.TypeOf(v)
	if _, ok := m.m[ty]; ok {
		return false
	}
	m.m[ty] = v
	return true
}
