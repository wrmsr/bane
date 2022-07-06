package marshal

import (
	"reflect"

	"github.com/wrmsr/bane/pkg/util/check"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type MapMarshaler struct {
	k Marshaler
	v Marshaler
}

func NewMapMarshaler(k, v Marshaler) MapMarshaler {
	return MapMarshaler{k: k, v: v}
}

var _ Marshaler = MapMarshaler{}

func (m MapMarshaler) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	if rv.IsNil() {
		return _nullValue, nil
	}

	vs := make([]bt.Kv[Value, Value], rv.Len())
	for i, mr := 0, rv.MapRange(); mr.Next(); i++ {
		k, err := m.k.Marshal(ctx, mr.Key())
		if err != nil {
			return nil, err
		}
		v, err := m.v.Marshal(ctx, mr.Value())
		if err != nil {
			return nil, err
		}
		vs[i] = bt.KvOf(k, v)
	}
	return Object{v: vs}, nil
}

//

type MapUnmarshaler struct {
	ty reflect.Type
	k  Unmarshaler
	v  Unmarshaler

	nv reflect.Value
}

func NewMapUnmarshaler(ty reflect.Type, k, v Unmarshaler) MapUnmarshaler {
	check.Equal(ty.Kind(), reflect.Map)
	return MapUnmarshaler{ty: ty, k: k, v: v, nv: rfl.ZeroFor(ty)}
}

var _ Unmarshaler = MapUnmarshaler{}

func (u MapUnmarshaler) Unmarshal(ctx UnmarshalContext, mv Value) (reflect.Value, error) {
	switch mv := mv.(type) {

	case Null:
		return u.nv, nil

	case Object:
		rv := reflect.New(u.ty).Elem()
		for _, kv := range mv.v {
			k, err := u.k.Unmarshal(ctx, kv.K)
			if err != nil {
				return rfl.Invalid(), err
			}
			v, err := u.v.Unmarshal(ctx, kv.V)
			if err != nil {
				return rfl.Invalid(), err
			}
			rv.SetMapIndex(k, v)
		}
		return rv, nil

	}
	return rfl.Invalid(), _unhandledType
}