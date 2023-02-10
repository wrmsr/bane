package marshal

import (
	"fmt"
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	bt "github.com/wrmsr/bane/pkg/util/types"
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
	iv := rv.Interface()
	o := iv.(bt.AnyOptional)
	if !o.Present() {
		return _nullValue, nil
	}

	return m.elem.Marshal(ctx, reflect.ValueOf(o.Interface()))
}

//

var _optionalInterfaceTy = rfl.TypeOf[bt.AnyOptional]()

var optionalMarshalerFactory = NewFuncFactory(func(ctx MarshalContext, ty reflect.Type) (Marshaler, error) {
	if !ty.AssignableTo(_optionalInterfaceTy) {
		return nil, nil
	}

	irv := reflect.New(ty).Interface()
	aov := irv.(bt.AnyOptional)
	av := aov.Interface()
	fmt.Println(av)
	ety := aov.Type()
	if ety == nil {
		panic("?")
	}
	elem, err := ctx.Make(ctx, ety)
	if err != nil {
		return nil, err
	}

	return NewOptionalMarshaler(ty, elem), nil
})

func NewOptionalMarshalerFactory() MarshalerFactory {
	return optionalMarshalerFactory
}

///

type OptionalUnmarshaler struct {
	ty   reflect.Type
	elem Unmarshaler

	nv  reflect.Value
	nvi bt.AnyOptional
}

func NewOptionalUnmarshaler(ty reflect.Type, elem Unmarshaler) OptionalUnmarshaler {
	nv := rfl.ZeroFor(ty)
	return OptionalUnmarshaler{ty: ty, elem: elem, nv: nv, nvi: nv.Interface().(bt.AnyOptional)}
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

var optionalUnmarshalerFactory = NewFuncFactory(func(ctx UnmarshalContext, ty reflect.Type) (Unmarshaler, error) {
	if !ty.AssignableTo(_optionalInterfaceTy) {
		return nil, nil
	}

	ety := reflect.TypeOf(reflect.New(ty).Interface().(bt.AnyOptional).ZeroInterface())
	elem, err := ctx.Make(ctx, ety)
	if err != nil {
		return nil, err
	}

	return NewOptionalUnmarshaler(ty, elem), nil
})

func NewOptionalUnmarshalerFactory() UnmarshalerFactory {
	return optionalUnmarshalerFactory
}
