package marshal

import (
	"encoding"
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

///

type StdTextMarshaler struct{}

var _ Marshaler = StdTextMarshaler{}

func (m StdTextMarshaler) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	if rfl.SafeIsNil(rv) {
		return _nullValue, nil
	}
	tm := rv.Interface().(encoding.TextMarshaler)
	b, err := tm.MarshalText()
	if err != nil {
		return nil, err
	}
	return String{v: string(b)}, nil
}

//

var stdTextMarshalerTy = rfl.TypeOf[encoding.TextMarshaler]()

var stdTextMarshalerFactory = NewFuncFactory(func(ctx MarshalContext, ty reflect.Type) (Marshaler, error) {
	if !ty.AssignableTo(stdTextMarshalerTy) {
		return nil, nil
	}
	return StdTextMarshaler{}, nil
})

func NewStdTextMarshalerFactory() MarshalerFactory {
	return stdTextMarshalerFactory
}

///

type StdTextUnmarshaler struct{}

var _ Unmarshaler = StdTextUnmarshaler{}

func (u StdTextUnmarshaler) Unmarshal(ctx UnmarshalContext, mv Value) (reflect.Value, error) {
	switch mv := mv.(type) {

	case Null:

	}
	if rfl.SafeIsNil(rv) {
		return _nullValue, nil
	}
	tm := rv.Interface().(encoding.TextMarshaler)
	b, err := tm.MarshalText()
	if err != nil {
		return nil, err
	}
	return String{v: string(b)}, nil
}

//

var stdTextUnmarshalerFactory = NewFuncFactory(func(ctx UnmarshalContext, ty reflect.Type) (Unmarshaler, error) {
	if !ty.AssignableTo(stdTextMarshalerTy) {
		return nil, nil
	}
	return StdTextUnmarshaler{}, nil
})

func NewStdTextUnmarshalerFactory() UnmarshalerFactory {
	return stdTextUnmarshalerFactory
}
