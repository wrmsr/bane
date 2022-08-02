package container

import (
	"reflect"

	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type StdTypeMap[T any] struct {
	m map[reflect.Type]T
}

func NewStdTypeMap[T any](it bt.Iterable[T]) StdTypeMap[T] {
	r := StdTypeMap[T]{}
	if it != nil {
		m := make(map[reflect.Type]T)
		for it := it.Iterate(); it.HasNext(); {
			v := it.Next()
			m[reflect.TypeOf(v)] = v
		}
		r.m = m
	}
	return r
}

var _ TypeMap[any] = StdTypeMap[any]{}

func (m StdTypeMap[T]) Len() int {
	if m.m == nil {
		return 0
	}
	return len(m.m)
}

func (m StdTypeMap[T]) Contains(ty reflect.Type) bool {
	if m.m == nil {
		return false
	}
	_, ok := m.m[ty]
	return ok
}

func (m StdTypeMap[T]) Get(ty reflect.Type) T {
	if m.m == nil {
		var z T
		return z
	}
	v, _ := m.m[ty]
	return v
}

func (m StdTypeMap[T]) TryGet(ty reflect.Type) (T, bool) {
	if m.m == nil {
		var z T
		return z, false
	}
	v, ok := m.m[ty]
	return v, ok
}

func (m StdTypeMap[T]) Iterate() bt.Iterator[T] {
	if m.m == nil {
		return its.Empty[T]().Iterate()
	}
	vs := make([]T, len(m.m))
	i := 0
	for _, v := range m.m {
		vs[i] = v
		i++
	}
	return its.OfSlice(vs).Iterate()
}

func (m StdTypeMap[T]) ForEach(fn func(v T) bool) bool {
	if m.m == nil {
		return true
	}
	for _, v := range m.m {
		if !fn(v) {
			return false
		}
	}
	return true
}

//

type MutStdTypeMap[T any] struct {
	m StdTypeMap[T]
}

func NewMutTypeMap[T any](it bt.Iterable[T]) *MutStdTypeMap[T] {
	return &MutStdTypeMap[T]{NewStdTypeMap(it)}
}

var _ MutTypeMap[any] = &MutStdTypeMap[any]{}

func (m *MutStdTypeMap[T]) isMut() {}

func (m *MutStdTypeMap[T]) Len() int                         { return m.m.Len() }
func (m *MutStdTypeMap[T]) Contains(ty reflect.Type) bool    { return m.m.Contains(ty) }
func (m *MutStdTypeMap[T]) Get(ty reflect.Type) T            { return m.m.Get(ty) }
func (m *MutStdTypeMap[T]) TryGet(ty reflect.Type) (T, bool) { return m.m.TryGet(ty) }
func (m *MutStdTypeMap[T]) Iterate() bt.Iterator[T]          { return m.m.Iterate() }
func (m *MutStdTypeMap[T]) ForEach(fn func(v T) bool) bool   { return m.m.ForEach(fn) }

func (m *MutStdTypeMap[T]) lazyInit() {
	if m.m.m == nil {
		m.m.m = make(map[reflect.Type]T)
	}
}

func (m *MutStdTypeMap[T]) Put(v T) {
	m.lazyInit()
	m.m.m[reflect.TypeOf(v)] = v
}

func (m *MutStdTypeMap[T]) Delete(ty reflect.Type) {
	m.lazyInit()
	delete(m.m.m, ty)
}

func (m *MutStdTypeMap[T]) Default(v T) bool {
	m.lazyInit()
	ty := reflect.TypeOf(v)
	if _, ok := m.m.m[ty]; ok {
		return false
	}
	m.m.m[ty] = v
	return true
}
