package marshal

import (
	"reflect"

	"github.com/wrmsr/bane/pkg/util/check"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

///

type KvIterableMarshaler struct {
	k Marshaler
	v Marshaler
}

func (m KvIterableMarshaler) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	if rv.IsNil() {
		return _nullValue, nil
	}

	itm := check.Ok1(rv.Type().MethodByName("Iterate"))
	itv := check.Single(itm.Func.Call([]reflect.Value{rv}))
	if itv.Kind() == reflect.Interface {
		itv = itv.Elem()
	}
	ity := itv.Type()

	nm := check.Ok1(ity.MethodByName("Next")).Func
	hnm := check.Ok1(ity.MethodByName("HasNext")).Func

	var vs []bt.Kv[Value, Value]
	for check.Single(hnm.Call([]reflect.Value{itv})).Bool() {
		kv := check.Single(nm.Call([]reflect.Value{itv})).Interface().(bt.AnyKv)
		k, err := m.k.Marshal(ctx, reflect.ValueOf(kv.AnyK()))
		if err != nil {
			return nil, err
		}
		v, err := m.v.Marshal(ctx, reflect.ValueOf(kv.AnyV()))
		if err != nil {
			return nil, err
		}
		vs = append(vs, bt.KvOf(k, v))
	}
	return Object{v: vs}, nil
}
