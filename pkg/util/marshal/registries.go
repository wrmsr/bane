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

type typeRegistry struct {
	ty reflect.Type
	s  []RegistryItem
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

	// FIXME: copy-on-write
	for _, i := range items {
		ity := reflect.TypeOf(i)
		ti.s = append(ti.s, i)
		ti.m[ity] = append(ti.m[ity], i)
	}

	return r
}

func (r *Registry) Get(ty reflect.Type) []RegistryItem {
	e := r.m[ty]
	if e == nil {
		return nil
	}
	return e.s
}

func (r *Registry) GetOf(ty, ity reflect.Type) []RegistryItem {
	e := r.m[ty]
	if e == nil {
		return nil
	}
	return e.m[ity]
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

func NewRegistryMarshalerFactory() RegistryMarshalerFactory {
	return RegistryMarshalerFactory{}
}

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

func NewRegistryUnmarshalerFactory() RegistryUnmarshalerFactory {
	return RegistryUnmarshalerFactory{}
}

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
