package marshal

import "reflect"

//

type PrimitiveMarshaller struct{}

var _ Marshaller = PrimitiveMarshaller{}

func (p PrimitiveMarshaller) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
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
	return nil, UnhandledTypeError{}
}

//

type PrimitiveUnmarshaller struct{}

var _ Unmarshaller = PrimitiveUnmarshaller{}

func (p PrimitiveUnmarshaller) Unmarshal(ctx MarshalContext, mv Value) (reflect.Value, error) {
	panic("implement me")
}
