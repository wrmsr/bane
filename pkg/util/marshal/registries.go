package marshal

import (
	"reflect"
)

//

type RegistryItem interface {
	isRegistryItem()
}

//

type SetMarshaler struct {
	M Marshaler
}

func (ri SetMarshaler) isRegistryItem() {}

type SetUnmarshaler struct {
	U Unmarshaler
}

func (ri SetUnmarshaler) isRegistryItem() {}

//

type typeRegistry struct {
	ty reflect.Type
	m  map[reflect.Type][]RegistryItem
}

type Registry struct {
	m  map[reflect.Type]*typeRegistry
	ps []*Registry
}

func NewRegistry(parents []*Registry) *Registry {
	return &Registry{
		m: make(map[reflect.Type]*typeRegistry),

		ps: parents,
	}
}

func (r *Registry) Register(ty reflect.Type, items ...RegistryItem) *Registry {
	ti := r.m[ty]
	if ti == nil {
		ti = &typeRegistry{
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
