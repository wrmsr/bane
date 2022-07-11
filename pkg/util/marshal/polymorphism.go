package marshal

import "reflect"

///

type PolymorphismEntry struct {
	Tag  string
	Impl reflect.Type
}

type Polymorphism struct {
	ty reflect.Type
	es []PolymorphismEntry
}

func NewPolymorphism(ty reflect.Type, es ...PolymorphismEntry) *Polymorphism {
	return &Polymorphism{
		ty: ty,
		es: es,
	}
}

///

type PolymorphismMarshaler struct {
	p *Polymorphism
	m map[reflect.Type]Marshaler
}

func NewPolymorphismMarshaler(p *Polymorphism, m map[reflect.Type]Marshaler) PolymorphismMarshaler {
	return PolymorphismMarshaler{p: p, m: m}
}

var _ Marshaler = PolymorphismMarshaler{}

func (m PolymorphismMarshaler) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	panic("implement me")
}

///

type PolymorphismUnmarshaler struct {
	p *Polymorphism
	m map[string]Unmarshaler
}

func NewPolymorphismUnmarshaler(p *Polymorphism, m map[string]Unmarshaler) PolymorphismUnmarshaler {
	return PolymorphismUnmarshaler{p: p, m: m}
}

var _ Unmarshaler = PolymorphismUnmarshaler{}

func (u PolymorphismUnmarshaler) Unmarshal(ctx UnmarshalContext, mv Value) (reflect.Value, error) {
	panic("implement me")
}
