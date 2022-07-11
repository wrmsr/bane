package marshal

import (
	"reflect"

	ctr "github.com/wrmsr/bane/pkg/util/container"
)

type RegistryItem interface {
	isRegistryItem()
}

type typeRegistry struct {
	ty reflect.Type
	m  ctr.MutTypeMap[RegistryItem]
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

func (r *Registry) Register(ty reflect.Type, item RegistryItem) *Registry {
	ti := r.m[ty]
	if ti == nil {
		ti = &typeRegistry{
			ty: ty,
			m:  ctr.NewMutTypeMap[RegistryItem](nil),
		}
		r.m[ty] = ti
	}

	return r
}
