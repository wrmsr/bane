package marshal

import (
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

//

type PointerMarshaler struct {
	elem Marshaler
}

func NewPointerMarshaler(elem Marshaler) PointerMarshaler {
	return PointerMarshaler{elem: elem}
}

var _ Marshaler = PointerMarshaler{}

func (m PointerMarshaler) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	if rv.IsNil() {
		return _nullValue, nil
	}

	return m.Marshal(ctx, rv.Elem())
}

//

type PointerUnmarshaler struct {
	ty   reflect.Type
	elem Unmarshaler

	nv reflect.Value
}

func NewPointerUnmarshaler(ty reflect.Type, elem Unmarshaler) PointerUnmarshaler {
	return PointerUnmarshaler{ty: ty, elem: elem, nv: rfl.ZeroFor(ty)}
}

var _ Unmarshaler = PointerUnmarshaler{}

func (u PointerUnmarshaler) Unmarshal(ctx UnmarshalContext, mv Value) (reflect.Value, error) {
	switch mv := mv.(type) {

	case Null:
		return u.nv, nil

	default:
		return u.elem.Unmarshal(ctx, mv)
	}
}
