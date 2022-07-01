package marshal

import (
	"encoding/json"
	"io"

	ju "github.com/wrmsr/bane/pkg/util/json"
)

func JsonEncode(w io.Writer, v Value) error {
	switch v := v.(type) {

	case Array:
		panic("nyi")

	case Bool:
		_, err := w.Write(ju.BoolBytes(v.v))
		return err

	case Bytes:
		return _unhandledType

	case Float:

	}

}

func (v Array) MarshalJSON() ([]byte, error) {
	panic("nyi")
}

func (v Bool) MarshalJSON() ([]byte, error) {
	return ju.BoolBytes(v.v), nil
}

func (v Bytes) MarshalJSON() ([]byte, error) {
	return nil, _unhandledType
}

func (v Float) MarshalJSON() ([]byte, error) {
	panic("nyi")
}

func (v Int) MarshalJSON() ([]byte, error) {
	panic("nyi")
}

func (v Null) MarshalJSON() ([]byte, error) {
	return ju.NullBytes(), nil
}

func (v Number) MarshalJSON() ([]byte, error) {
	return []byte(v.v.String()), nil
}

func (v Object) MarshalJSON() ([]byte, error) {
	panic("nyi")
}

func (v String) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.v)
}
