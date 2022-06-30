package marshal

import (
	"encoding/json"
	"errors"
)

func (v Array) MarshalJSON() ([]byte, error) {
	panic("nyi")
}

func (v Bool) MarshalJSON() ([]byte, error) {
	if v.v {
		return []byte("true"), nil
	} else {
		return []byte("false"), nil
	}
}

func (v Bytes) MarshalJSON() ([]byte, error) {
	return nil, errors.New("cannot json bytes")
}

func (v Float) MarshalJSON() ([]byte, error) {
	panic("nyi")
}

func (v Int) MarshalJSON() ([]byte, error) {
	panic("nyi")
}

func (v Null) MarshalJSON() ([]byte, error) {
	return []byte("null"), nil
}

func (v Number) MarshalJSON() ([]byte, error) {
	panic("nyi")
}

func (v Object) MarshalJSON() ([]byte, error) {
	panic("nyi")
}

func (v String) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.v)
}
