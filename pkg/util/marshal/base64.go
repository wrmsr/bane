package marshal

import (
	"encoding/base64"
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

//

type Base64Marshaler struct{}

var _ Marshaler = Base64Marshaler{}

func (m Base64Marshaler) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	if rv.Type() != rfl.Bytes() {
		return nil, _unhandledType
	}

	if rv.IsNil() {
		return _nullValue, nil
	}

	return String{v: base64.StdEncoding.EncodeToString(rv.Bytes())}, nil
}

//

type Base64Unmarshaler struct{}

var _ Unmarshaler = Base64Unmarshaler{}

func (u Base64Unmarshaler) Unmarshal(ctx UnmarshalContext, mv Value) (reflect.Value, error) {
	switch mv := mv.(type) {

	case Null:
		return rfl.NilBytes(), nil

	case String:
		b, err := base64.StdEncoding.DecodeString(mv.v)
		if err != nil {
			return rfl.Invalid(), err
		}
		return reflect.ValueOf(b), nil

	}
	return rfl.Invalid(), _unhandledType
}
