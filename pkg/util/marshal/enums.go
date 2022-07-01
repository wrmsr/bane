package marshal

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/wrmsr/bane/pkg/util/maps"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

//

type EnumMarshaler struct {
	m map[any]Value
}

func NewEnumMarshaler(m map[any]string) EnumMarshaler {
	return EnumMarshaler{m: maps.MapValues(MakeString, m)}
}

var _ Marshaler = EnumMarshaler{}

func (m EnumMarshaler) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	iv := rv.Interface()
	mv, ok := m.m[iv]
	if !ok {
		return nil, errors.New("unknown enum value")
	}
	return mv, nil
}

//

type EnumUnmarshaler struct {
	m map[string]reflect.Value
}

func NewEnumUnmarshaler(m map[string]any) EnumUnmarshaler {
	return EnumUnmarshaler{m: maps.MapValues(reflect.ValueOf, m)}
}

var _ Unmarshaler = EnumUnmarshaler{}

func (u EnumUnmarshaler) Unmarshal(ctx UnmarshalContext, mv Value) (reflect.Value, error) {
	switch mv := mv.(type) {

	case String:
		s := mv.v
		v, ok := u.m[s]
		if !ok {
			return rfl.Invalid(), fmt.Errorf("unknown enum value: %s", s)
		}
		return v, nil

	}
	return rfl.Invalid(), _unhandledType
}
