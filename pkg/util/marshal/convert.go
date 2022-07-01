package marshal

import (
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

//

type ConvertUnmarshaler struct {
	ty    reflect.Type
	child Unmarshaler
}

func NewConvertUnmarshaler(ty reflect.Type, child Unmarshaler) ConvertUnmarshaler {
	return ConvertUnmarshaler{ty: ty, child: child}
}

var _ Unmarshaler = ConvertUnmarshaler{}

func (u ConvertUnmarshaler) Unmarshal(ctx UnmarshalContext, mv Value) (reflect.Value, error) {
	rv, err := u.child.Unmarshal(ctx, mv)
	if err != nil {
		return rfl.Invalid(), err
	}

	return rv.Convert(u.ty), nil
}
