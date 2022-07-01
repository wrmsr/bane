package marshal

import (
	"reflect"

	opt "github.com/wrmsr/bane/pkg/util/optional"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type ObjectMarshalerField struct {
	Name string
	Get  func(ctx MarshalContext, rv reflect.Value) (opt.Optional[reflect.Value], error)
	Impl Marshaler
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
