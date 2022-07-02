package marshal

import (
	"fmt"
	"reflect"

	opt "github.com/wrmsr/bane/pkg/util/optional"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	stu "github.com/wrmsr/bane/pkg/util/structs"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type ObjectFieldGetter = func(ctx MarshalContext, rv reflect.Value) (opt.Optional[reflect.Value], error)

func NewObjectFieldGetter(fi *stu.FieldInfo) ObjectFieldGetter {
	return func(ctx MarshalContext, rv reflect.Value) (opt.Optional[reflect.Value], error) {
		fv, ok := fi.GetValue(rv)
		if !ok {
			return opt.None[reflect.Value](), nil
		}
		return opt.Just(fv), nil
	}
}

type ObjectMarshalerField struct {
	Name string
	Get  ObjectFieldGetter
	Impl Marshaler
}

func NewObjectMarshalerField(
	name string,
	get ObjectFieldGetter,
	impl Marshaler,
) ObjectMarshalerField {
	return ObjectMarshalerField{Name: name, Get: get, Impl: impl}
}

type ObjectMarshaler struct {
	flds []ObjectMarshalerField
}

func NewObjectMarshaler(flds ...ObjectMarshalerField) ObjectMarshaler {
	return ObjectMarshaler{flds: flds}
}

var _ Marshaler = ObjectMarshaler{}

func (m ObjectMarshaler) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	kvs := make([]bt.Kv[Value, Value], 0, len(m.flds))
	for _, fld := range m.flds {
		frv, err := fld.Get(ctx, rv)
		if err != nil {
			return nil, err
		}
		if !frv.Present() {
			continue
		}

		fmv, err := fld.Impl.Marshal(ctx, frv.Value())
		if err != nil {
			return nil, err
		}
		kvs = append(kvs, bt.KvOf[Value, Value](String{v: fld.Name}, fmv))
	}
	return Object{v: kvs}, nil
}

//

type ObjectFieldSetter = func(ctx UnmarshalContext, ov, fv reflect.Value) error

func NewObjectFieldSetter(fi *stu.FieldInfo) ObjectFieldSetter {
	return func(ctx UnmarshalContext, ov, fv reflect.Value) error {
		panic("nyi")
	}
}

type ObjectUnmarshalerField struct {
	Name string
	Set  ObjectFieldSetter
	Impl Unmarshaler
}

func NewObjectUnmarshalerField(
	name string,
	set ObjectFieldSetter,
	impl Unmarshaler,
) ObjectUnmarshalerField {
	return ObjectUnmarshalerField{Name: name, Set: set, Impl: impl}
}

type ObjectUnmarshaler struct {
	fac  func() reflect.Value
	flds []ObjectUnmarshalerField

	m map[string]*ObjectUnmarshalerField
}

func NewObjectUnmarshaler(fac func() reflect.Value, flds ...ObjectUnmarshalerField) ObjectUnmarshaler {
	m := make(map[string]*ObjectUnmarshalerField, len(flds))
	for i, f := range flds {
		if _, ok := m[f.Name]; ok {
			if ok {
				panic(fmt.Errorf("duplicate field: %s", f.Name))
			}
		}
		m[f.Name] = &flds[i]
	}
	return ObjectUnmarshaler{fac: fac, flds: flds, m: m}
}

var _ Unmarshaler = ObjectUnmarshaler{}

func (u ObjectUnmarshaler) Unmarshal(ctx UnmarshalContext, mv Value) (reflect.Value, error) {
	switch mv := mv.(type) {

	case Object:
		ov := u.fac()
		for _, kv := range mv.v {
			k, ok := kv.K.(String)
			if !ok {
				return rfl.Invalid(), _unhandledType
			}
			f, ok := u.m[k.v]
			if !ok {
				return rfl.Invalid(), fmt.Errorf("unknown field: %s", k.v)
			}
			fv, err := f.Impl.Unmarshal(ctx, kv.V)
			if err != nil {
				return rfl.Invalid(), err
			}
			if err := f.Set(ctx, ov, fv); err != nil {
				return rfl.Invalid(), err
			}
		}
		return ov, nil

	}
	return rfl.Invalid(), _unhandledType
}
