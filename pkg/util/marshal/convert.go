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

//

var convertPrimitiveUnmarshalerFactory = func() UnmarshalerFactory {
	m := make(map[reflect.Type]Unmarshaler, len(primitiveTypes))
	for ty, cty := range primitiveTypes {
		if ty != cty {
			m[ty] = NewConvertUnmarshaler(ty, PrimitiveUnmarshaler{})
		} else {
			m[ty] = PrimitiveUnmarshaler{}
		}
	}
	return NewTypeMapUnmarshalerFactory(m)
}()

func NewConvertPrimitiveUnmarshalerFactory() UnmarshalerFactory {
	return convertPrimitiveUnmarshalerFactory
}
