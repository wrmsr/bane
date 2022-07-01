package marshal

import (
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

//

type IndexMarshaler struct {
	elem Marshaler
}

func NewIndexMarshaler(elem Marshaler) IndexMarshaler {
	return IndexMarshaler{elem: elem}
}

var _ Marshaler = IndexMarshaler{}

func (m IndexMarshaler) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	if rv.Kind() == reflect.Slice && rv.IsNil() {
		return _nullValue, nil
	}

	l := rv.Len()
	vs := make([]Value, l)
	for i := 0; i < l; i++ {
		v, err := m.elem.Marshal(ctx, rv.Index(i))
		if err != nil {
			return nil, err
		}
		vs[i] = v
	}
	return Array{v: vs}, nil
}

//

type SliceUnmarshaler struct {
	ty   reflect.Type
	elem Unmarshaler

	nv reflect.Value
}

func NewSliceUnmarshaler(ty reflect.Type, elem Unmarshaler) SliceUnmarshaler {
	return SliceUnmarshaler{ty: ty, elem: elem, nv: reflect.New(ty).Elem()}
}

var _ Unmarshaler = SliceUnmarshaler{}

func (u SliceUnmarshaler) Unmarshal(ctx UnmarshalContext, mv Value) (reflect.Value, error) {
	switch mv := mv.(type) {

	case Null:
		return u.nv, nil

	case Array:
		l := len(mv.v)
		r := reflect.MakeSlice(u.ty, l, l)
		for i, em := range mv.v {
			er, err := u.elem.Unmarshal(ctx, em)
			if err != nil {
				return rfl.Invalid(), err
			}
			r.Index(i).Set(er)
		}
		return r, nil

	}
	return rfl.Invalid(), _unhandledType
}
