package marshal

import (
	"reflect"

	opt "github.com/wrmsr/bane/pkg/util/optional"
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
