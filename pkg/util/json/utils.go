package json

import (
	"encoding/json"
	"errors"
)

//

var _nullBytes = []byte("null")

func NullBytes() []byte {
	return _nullBytes
}

//

func MarshalString(v any) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

//

func ReadDelim(dec *json.Decoder) (json.Delim, error) {
	token, err := dec.Token()
	if err != nil {
		return 0, err
	}
	if delim, ok := token.(json.Delim); ok {
		return delim, nil
	}
	return 0, errors.New("expected delim")
}

func EndDelim(d json.Delim) json.Delim {
	switch d {
	case '{':
		return '}'
	case '[':
		return ']'
	}
	panic(errors.New("unreachable"))
}
