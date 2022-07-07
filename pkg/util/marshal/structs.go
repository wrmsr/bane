package marshal

import (
	"errors"
	"reflect"

	opt "github.com/wrmsr/bane/pkg/util/optional"
	stu "github.com/wrmsr/bane/pkg/util/structs"
)

///

func NewStructFieldGetter(fi *stu.FieldInfo) ObjectFieldGetter {
	return func(ctx MarshalContext, rv reflect.Value) (opt.Optional[reflect.Value], error) {
		fv, ok := fi.GetValue(rv)
		if !ok {
			return opt.None[reflect.Value](), nil
		}
		return opt.Just(fv), nil
	}
}

//

type StructMarshalerFactory struct {
	sic *stu.StructInfoCache
}

func NewStructMarshalerFactory(sic *stu.StructInfoCache) StructMarshalerFactory {
	return StructMarshalerFactory{sic: sic}
}

var _ MarshalerFactory = StructMarshalerFactory{}

func (mf StructMarshalerFactory) Make(ctx MarshalContext, ty reflect.Type) (Marshaler, error) {
	if ty.Kind() != reflect.Struct {
		return nil, nil
	}

	si := mf.sic.Info(ty)
	si.Fields().ByFlat()

	panic("nyi")
}

///

func NewStructFactory(si *stu.StructInfo) ObjectFactory {
	return func(ctx UnmarshalContext) reflect.Value {
		return reflect.New(si.Type()).Elem()
	}
}

func NewStructFieldSetter(fi *stu.FieldInfo) ObjectFieldSetter {
	return func(ctx UnmarshalContext, ov, fv reflect.Value) error {
		if !fi.SetValue(ov.Addr().Interface(), fv.Interface()) {
			return errors.New("can't set struct value")
		}
		return nil
	}
}

//

type StructUnmarshalerFactory struct {
	sic *stu.StructInfoCache
}

func NewStructUnmarshalerFactory(sic *stu.StructInfoCache) StructUnmarshalerFactory {
	return StructUnmarshalerFactory{sic: sic}
}

var _ UnmarshalerFactory = StructUnmarshalerFactory{}

func (mf StructUnmarshalerFactory) Make(ctx UnmarshalContext, ty reflect.Type) (Unmarshaler, error) {
	if ty.Kind() != reflect.Struct {
		return nil, nil
	}

	panic("nyi")
}
