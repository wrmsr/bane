package marshal

import (
	"reflect"

	opt "github.com/wrmsr/bane/pkg/util/optional"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

///

type OptionalMarshaler struct {
	ty   reflect.Type
	elem Marshaler
}

func NewOptionalMarshaler(ty reflect.Type, elem Marshaler) OptionalMarshaler {
	return OptionalMarshaler{ty: ty, elem: elem}
}

var _ Marshaler = OptionalMarshaler{}

func (m OptionalMarshaler) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	o := rv.Interface().(opt.Interface)
	if !o.Present() {
		return _nullValue, nil
	}

	return m.elem.Marshal(ctx, reflect.ValueOf(o.Interface()))
}

//

var optionalMarshalerFactory = NewFuncMarshalerFactory(func(ctx MarshalerFactoryContext, ty reflect.Type) (Marshaler, error) {

})

func NewOptionalMarshalerFactory() MarshalerFactory {
	return optionalMarshalerFactory
}

///

type OptionalUnmarshaler struct {
	ty   reflect.Type
	elem Unmarshaler

	nv  reflect.Value
	nvi opt.Interface
}

func NewOptionalUnmarshaler(ty reflect.Type, elem Unmarshaler) OptionalUnmarshaler {
	nv := rfl.ZeroFor(ty)
	return OptionalUnmarshaler{ty: ty, elem: elem, nv: nv, nvi: nv.Interface().(opt.Interface)}
}

var _ Unmarshaler = OptionalUnmarshaler{}

func (u OptionalUnmarshaler) Unmarshal(ctx UnmarshalContext, mv Value) (reflect.Value, error) {
	switch mv := mv.(type) {

	case Null:
		return u.nv, nil

	default:
		ev, err := u.elem.Unmarshal(ctx, mv)
		if err != nil {
			return rfl.Invalid(), err
		}
		return reflect.ValueOf(u.nvi.Replace(ev.Interface())), nil
	}
}

//

var optionalUnmarshalerFactory = NewFuncUnmarshalerFactory(func(ctx UnmarshalerFactoryContext, ty reflect.Type) (Unmarshaler, error) {

})

func NewOptionalUnmarshalerFactory() UnmarshalerFactory {
	return optionalUnmarshalerFactory
}
