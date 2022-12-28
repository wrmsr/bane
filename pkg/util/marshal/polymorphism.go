package marshal

import (
	"fmt"
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

///

type SetImpl struct {
	Impl reflect.Type
	Tag  string
	Alt  []string
}

func (ri SetImpl) isRegistryItem() {}

func SetImplOf[T any](tags ...string) SetImpl {
	var tag string
	var alt []string
	if len(tags) > 0 {
		tag = tags[0]
		alt = tags[1:]
	}
	return SetImpl{
		Impl: rfl.TypeOf[T](),
		Tag:  tag,
		Alt:  alt,
	}
}

///

type InheritImpls struct {
	Iface reflect.Type
}

func (hi InheritImpls) isRegistryItem() {}

func InheritImplsOf[T any]() InheritImpls {
	return InheritImpls{
		Iface: rfl.TypeOf[T](),
	}
}

///

type Polymorphism struct {
	ty reflect.Type
	es []SetImpl
	hs []InheritImpls

	im map[reflect.Type]*SetImpl
	tm map[string]*SetImpl
	hm map[reflect.Type]*InheritImpls
	nm map[string]*InheritImpls
}

func NewPolymorphism(ty reflect.Type, es []SetImpl, hs []InheritImpls, naming Naming) *Polymorphism {
	if naming == nil {
		naming = NopNaming()
	}

	im := make(map[reflect.Type]*SetImpl, len(es))
	tm := make(map[string]*SetImpl, len(es))
	for i := range es {
		e := &es[i]
		if e.Impl.Kind() == reflect.Interface {
			panic(fmt.Errorf("interfaces cannot be impls: %s -> %s", e.Impl, ty))
		}
		if !e.Impl.AssignableTo(ty) {
			panic(fmt.Errorf("not assignable: %s -> %s", e.Impl, ty))
		}
		if _, ok := im[e.Impl]; ok {
			panic(fmt.Errorf("duplicate impl: %s", e.Impl))
		}
		im[e.Impl] = e
		if e.Tag == "" {
			e.Tag = naming.Name(e.Impl.Name())
		}
		if _, ok := tm[e.Tag]; ok {
			panic(fmt.Errorf("duplicate tag: %s", e.Tag))
		}
		tm[e.Tag] = e
		for _, a := range e.Alt {
			if _, ok := tm[a]; ok {
				panic(fmt.Errorf("duplicate tag: %s", a))
			}
			tm[a] = e
		}
	}

	hm := make(map[reflect.Type]*InheritImpls, len(es))
	nm := make(map[string]*InheritImpls, len(es))
	for i := range hs {
		h := &hs[i]
		if h.Iface.Kind() != reflect.Interface {
			panic(fmt.Errorf("inherits must be interfaces: %s <- %s", h.Iface, ty))
		}
	}

	return &Polymorphism{
		ty: ty,
		es: es,
		hs: hs,

		im: im,
		tm: tm,
		hm: hm,
		nm: nm,
	}
}

///

type polymorphismMarshalerEntry struct {
	m Marshaler
	k String
}

type PolymorphismMarshaler struct {
	m map[reflect.Type]polymorphismMarshalerEntry
}

func NewPolymorphismMarshaler(p *Polymorphism, m map[reflect.Type]Marshaler) PolymorphismMarshaler {
	pm := make(map[reflect.Type]polymorphismMarshalerEntry, len(m))
	for ty, em := range m {
		e, ok := p.im[ty]
		if !ok {
			panic(fmt.Errorf("entry not found: %s", ty))
		}
		pm[ty] = polymorphismMarshalerEntry{
			m: em,
			k: MakeString(e.Tag),
		}
	}
	return PolymorphismMarshaler{m: pm}
}

var _ Marshaler = PolymorphismMarshaler{}

func (m PolymorphismMarshaler) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	if rv.Kind() == reflect.Interface {
		rv = rv.Elem()
	}

	e, ok := m.m[rv.Type()]
	if !ok {
		return nil, unhandledType()
	}

	mv, err := e.m.Marshal(ctx, rv)
	if err != nil {
		return nil, err
	}

	return MakeObject(bt.KvOf[Value, Value](e.k, mv)), nil
}

//

func NewPolymorphismMarshalerFactory(p *Polymorphism) MarshalerFactory {
	return NewFuncFactory(func(ctx MarshalContext, a reflect.Type) (Marshaler, error) {
		if a != p.ty {
			return nil, nil
		}
		m := make(map[reflect.Type]Marshaler, len(p.es))
		for _, e := range p.es {
			em, err := ctx.Make(ctx, e.Impl)
			if err != nil {
				return nil, err
			}
			m[e.Impl] = em
		}
		return NewPolymorphismMarshaler(p, m), nil
	})
}

var (
	_setImplTy      = rfl.TypeOf[SetImpl]()
	_inheritImplsTy = rfl.TypeOf[InheritImpls]()
)

func buildRegistryPolymorphism(r *Registry, ty reflect.Type) *Polymorphism {
	tr := r.Get(ty)
	if tr == nil {
		return nil
	}
	ris := tr.Get(_setImplTy)
	his := tr.Get(_inheritImplsTy)
	if len(ris) < 1 && len(his) < 1 {
		return nil
	}
	return NewPolymorphism(
		ty,
		slices.Map(func(ri RegistryItem) SetImpl { return ri.(SetImpl) }, ris),
		slices.Map(func(ri RegistryItem) InheritImpls { return ri.(InheritImpls) }, his),
		nil,
	)
}

func NewRegistryPolymorphismMarshalerFactory() MarshalerFactory {
	return NewFuncFactory(func(ctx MarshalContext, a reflect.Type) (Marshaler, error) {
		if ctx.Reg == nil {
			return nil, nil
		}
		p := buildRegistryPolymorphism(ctx.Reg, a)
		if p == nil {
			return nil, nil
		}
		return NewPolymorphismMarshalerFactory(p).Make(ctx, a)
	})
}

///

type PolymorphismUnmarshaler struct {
	m map[string]Unmarshaler
}

func NewPolymorphismUnmarshaler(m map[string]Unmarshaler) PolymorphismUnmarshaler {
	return PolymorphismUnmarshaler{m: m}
}

var _ Unmarshaler = PolymorphismUnmarshaler{}

func (u PolymorphismUnmarshaler) Unmarshal(ctx UnmarshalContext, mv Value) (reflect.Value, error) {
	switch mv := mv.(type) {

	case Object:
		if len(mv.v) != 1 {
			return rfl.Invalid(), fmt.Errorf("expected object of one item")
		}
		kv := mv.v[0]
		ks, ok := kv.K.(String)
		if !ok {
			return rfl.Invalid(), fmt.Errorf("expected object of one string key")
		}
		e, ok := u.m[ks.v]
		if !ok {
			return rfl.Invalid(), fmt.Errorf("impl not found: %s", ks.v)
		}
		return e.Unmarshal(ctx, kv.V)

	}
	return rfl.Invalid(), unhandledType()
}

//

func NewPolymorphismUnmarshalerFactory(p *Polymorphism) UnmarshalerFactory {
	return NewFuncFactory(func(ctx UnmarshalContext, a reflect.Type) (Unmarshaler, error) {
		if a != p.ty {
			return nil, nil
		}
		m := make(map[string]Unmarshaler, len(p.tm))
		for _, e := range p.tm {
			u, err := ctx.Make(ctx, e.Impl)
			if err != nil {
				return nil, err
			}
			m[e.Tag] = u
			for _, a := range e.Alt {
				m[a] = u
			}
		}
		return NewPolymorphismUnmarshaler(m), nil
	})
}

func NewRegistryPolymorphismUnmarshalerFactory() UnmarshalerFactory {
	return NewFuncFactory(func(ctx UnmarshalContext, a reflect.Type) (Unmarshaler, error) {
		if ctx.Reg == nil {
			return nil, nil
		}
		p := buildRegistryPolymorphism(ctx.Reg, a)
		if p == nil {
			return nil, nil
		}
		return NewPolymorphismUnmarshalerFactory(p).Make(ctx, a)
	})
}
