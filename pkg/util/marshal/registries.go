package marshal

import (
	"reflect"
)

//

type RegistryItem interface {
	isRegistryItem()
}

//

type SetType struct {
	Marshaler   Marshaler
	Unmarshaler Unmarshaler
}

func (ri SetType) isRegistryItem() {}

//

type TypeRegistry struct {
	ty reflect.Type
	m  map[reflect.Type][]RegistryItem
}

func (tr *TypeRegistry) Type() reflect.Type { return tr.ty }

func (tr *TypeRegistry) Get(ity reflect.Type) []RegistryItem {
	return tr.m[ity]
}

//

type Registry struct {
	m  map[reflect.Type]*TypeRegistry
	ps []*Registry
}

func NewRegistry(parents []*Registry) *Registry {
	return &Registry{
		m: make(map[reflect.Type]*TypeRegistry),

		ps: parents,
	}
}

func (r *Registry) Register(ty reflect.Type, items ...RegistryItem) *Registry {
	ti := r.m[ty]
	if ti == nil {
		ti = &TypeRegistry{
			ty: ty,
			m:  make(map[reflect.Type][]RegistryItem),
		}
		r.m[ty] = ti
	}

	for _, i := range items {
		ity := reflect.TypeOf(i)
		ti.m[ity] = append(ti.m[ity], i)
	}

	return r
}

func (r *Registry) Get(ty reflect.Type) *TypeRegistry {
	return r.m[ty]
}
