package marshal

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/wrmsr/bane/pkg/util/maps"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

//

type EnumMarshaler[T comparable] struct {
	m map[T]Value
}

func NewEnumMarshaler[T comparable](m map[T]string) EnumMarshaler[T] {
	return EnumMarshaler[T]{m: maps.MapValues(func(s string) Value { return MakeString(s) }, m)}
}

var _ Marshaler = EnumMarshaler[int]{}

func (m EnumMarshaler[T]) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	iv := rv.Interface()
	it, ok := iv.(T)
	if !ok {
		return nil, _unhandledType
	}
	mv, ok := m.m[it]
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
