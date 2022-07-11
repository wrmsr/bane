package marshal

import (
	"fmt"
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

///

type PolymorphismEntry struct {
	Impl reflect.Type
	Tag  string
	Alt  []string
}

type Polymorphism struct {
	ty reflect.Type
	es []PolymorphismEntry

	im map[reflect.Type]*PolymorphismEntry
	tm map[string]*PolymorphismEntry
}

func NewPolymorphism(ty reflect.Type, es ...PolymorphismEntry) *Polymorphism {
	im := make(map[reflect.Type]*PolymorphismEntry, len(es))
	tm := make(map[string]*PolymorphismEntry, len(es))
	for i := range es {
		e := &es[i]
		if _, ok := im[e.Impl]; ok {
			panic(fmt.Errorf("duplicate impl: %s", e.Impl))
		}
		im[e.Impl] = e
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
	return &Polymorphism{ty: ty, es: es, im: im, tm: tm}
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
	e, ok := m.m[rv.Type()]
	if !ok {
		return nil, _unhandledType
	}

	mv, err := e.m.Marshal(ctx, rv)
	if err != nil {
		return nil, err
	}

	return MakeObject(bt.KvOf[Value, Value](e.k, mv)), nil

}

//

func NewPolymorphicMarshalerFactory(p *Polymorphism) MarshalerFactory {
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
	return rfl.Invalid(), _unhandledType
}

//

func NewPolymorphicUnmarshalerFactory(p *Polymorphism) UnmarshalerFactory {
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
