package marshal

import (
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

///

type RegistryItem interface {
	isRegistryItem()
}

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

func (r *Registry) GetOf(ty, ity reflect.Type) []RegistryItem {
	e := r.m[ty]
	if e == nil {
		return nil
	}
	return e.m[ity]
}

//

var globalRegistry = NewRegistry(nil)

func GlobalRegistry() *Registry {
	return globalRegistry
}

func Register(ty reflect.Type, items ...RegistryItem) *Registry {
	return globalRegistry.Register(ty, items...)
}

///

type SetType struct {
	Marshaler   Marshaler
	Unmarshaler Unmarshaler

	MarshalerFactory   MarshalerFactory
	UnmarshalerFactory UnmarshalerFactory
}

func (ri SetType) isRegistryItem() {}

//

type RegistryMarshalerFactory struct{}

var _ MarshalerFactory = RegistryMarshalerFactory{}

var _setTypeTy = rfl.TypeOf[SetType]()

func (f RegistryMarshalerFactory) Make(ctx MarshalContext, a reflect.Type) (Marshaler, error) {
	if ctx.Reg == nil {
		return nil, nil
	}
	sts := ctx.Reg.GetOf(a, _setTypeTy)
	for i := len(sts) - 1; i >= 0; i-- {
		st := sts[i].(SetType)
		if st.MarshalerFactory != nil {
			return st.MarshalerFactory.Make(ctx, a)
		}
		if st.Marshaler != nil {
			return st.Marshaler, nil
		}
	}
	return nil, nil
}

//

type RegistryUnmarshalerFactory struct{}

var _ UnmarshalerFactory = RegistryUnmarshalerFactory{}

func (f RegistryUnmarshalerFactory) Make(ctx UnmarshalContext, a reflect.Type) (Unmarshaler, error) {
	if ctx.Reg == nil {
		return nil, nil
	}
	sts := ctx.Reg.GetOf(a, _setTypeTy)
	for i := len(sts) - 1; i >= 0; i-- {
		st := sts[i].(SetType)
		if st.UnmarshalerFactory != nil {
			return st.UnmarshalerFactory.Make(ctx, a)
		}
		if st.Unmarshaler != nil {
			return st.Unmarshaler, nil
		}
	}
	return nil, nil
}
