package marshal

import (
	"reflect"

	"github.com/wrmsr/bane/pkg/util/check"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type kvIterableMarshalerMethods struct {
	hnm, nm reflect.Value
}

type KvIterableMarshaler struct {
	k Marshaler
	v Marshaler

	itm reflect.Value
	mc  map[reflect.Type]kvIterableMarshalerMethods
}

func NewKvIterableMarshaler(ty reflect.Type, k, v Marshaler) KvIterableMarshaler {
	itm := check.Ok1(ty.MethodByName("Iterate")).Func
	return KvIterableMarshaler{k: k, v: v, itm: itm, mc: make(map[reflect.Type]kvIterableMarshalerMethods)}
}

func (m KvIterableMarshaler) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	if rv.IsNil() {
		return _nullValue, nil
	}

	itv := check.Single(m.itm.Call([]reflect.Value{rv}))
	if itv.Kind() == reflect.Interface {
		itv = itv.Elem()
	}
	ity := itv.Type()

	ms, ok := m.mc[ity]
	if !ok {
		ms = kvIterableMarshalerMethods{
			nm:  check.Ok1(ity.MethodByName("Next")).Func,
			hnm: check.Ok1(ity.MethodByName("HasNext")).Func,
		}
		m.mc[ity] = ms
	}

	var vs []bt.Kv[Value, Value]
	for check.Single(ms.hnm.Call([]reflect.Value{itv})).Bool() {
		kv := check.Single(ms.nm.Call([]reflect.Value{itv})).Interface().(bt.AnyKv)
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

//

var kvIterableMarshalerFactory = NewFuncFactory(func(ctx MarshalContext, ty reflect.Type) (Marshaler, error) {
	if ty.Kind() != reflect.Pointer {
		return nil, nil
	}

	elem, err := ctx.Make(ctx, ty.Elem())
	if err != nil {
		return nil, err
	}
	return NewPointerMarshaler(elem), nil
})

func NewKvIterableMarshalerFactory() MarshalerFactory {
	return kvIterableMarshalerFactory
}
