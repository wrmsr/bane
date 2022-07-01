package marshal

import (
	"encoding/base64"
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

//

type Base64Marshaller struct{}

var _ Marshaller = Base64Marshaller{}

func (m Base64Marshaller) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	if rv.Type() != rfl.Bytes() {
		return nil, _unhandledType
	}

	if rv.IsNil() {
		return _nullValue, nil
	}

	return String{v: base64.StdEncoding.EncodeToString(rv.Bytes())}, nil
}

//

type Base64Unmarshaller struct{}

var _ Unmarshaller = Base64Unmarshaller{}

func (u Base64Unmarshaller) Unmarshal(ctx MarshalContext, mv Value) (reflect.Value, error) {
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
