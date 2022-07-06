package marshal

import (
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

//

var primitiveTypes = []reflect.Type{
	rfl.TypeOf[int](),
	rfl.TypeOf[int8](),
	rfl.TypeOf[int16](),
	rfl.TypeOf[int32](),
	rfl.TypeOf[int64](),
	rfl.TypeOf[uint](),
	rfl.TypeOf[uint8](),
	rfl.TypeOf[uint16](),
	rfl.TypeOf[uint32](),
	rfl.TypeOf[uint64](),
	rfl.TypeOf[uintptr](),
	rfl.TypeOf[float32](),
	rfl.TypeOf[float64](),
	rfl.TypeOf[string](),
}

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

var primitiveMarshalerFactory = NewSimpleMarshalerFactory(primitiveTypes, PrimitiveMarshaler{})

func NewPrimitiveMarshalerFactory() MarshalerFactory {
	return primitiveMarshalerFactory
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

//

var primitiveUnmarshalerFactory = NewSimpleUnmarshalerFactory(primitiveTypes, PrimitiveUnmarshaler{})

func NewPrimitiveUnmarshalerFactory() UnmarshalerFactory {
	return primitiveUnmarshalerFactory
}
