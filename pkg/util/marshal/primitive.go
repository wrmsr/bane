package marshal

import (
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

//

type Primitive interface {
	isPrimitive()
}

func (v Bool) isPrimitive()   {}
func (v Int) isPrimitive()    {}
func (v Float) isPrimitive()  {}
func (v String) isPrimitive() {}

//

type PrimitiveMarshaler struct{}

var _ Marshaler = PrimitiveMarshaler{}

func (p PrimitiveMarshaler) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	switch rv.Kind() {

	case reflect.Bool:
		return Bool{v: rv.Bool()}, nil

	case
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		return Int{v: rv.Int()}, nil

	case
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr:
		return Int{v: int64(rv.Uint()), u: true}, nil

	case
		reflect.Float32,
		reflect.Float64:
		return Float{v: rv.Float()}, nil

	case reflect.String:
		return String{v: rv.String()}, nil

	}
	return nil, _unhandledType
}

//

type PrimitiveUnmarshaler struct{}

var _ Unmarshaler = PrimitiveUnmarshaler{}

func (p PrimitiveUnmarshaler) Unmarshal(ctx UnmarshalContext, mv Value) (reflect.Value, error) {
	switch mv := mv.(type) {

	case Bool:
		if mv.v {
			return rfl.True(), nil
		} else {
			return rfl.False(), nil
		}

	case Int:
		if mv.u {
			return reflect.ValueOf(uint64(mv.v)), nil
		} else {
			return reflect.ValueOf(mv.v), nil
		}

	case Float:
		return reflect.ValueOf(mv.v), nil

	case String:
		return reflect.ValueOf(mv.v), nil

	}
	return rfl.Invalid(), _unhandledType
}
