package json

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type RawField = bt.Kv[string, json.RawMessage]
type RawObject []RawField

func RawFieldOf(k string, v json.RawMessage) bt.Kv[string, json.RawMessage] {
	return bt.KvOf(k, v)
}

//

func (o RawObject) MarshalJSON() ([]byte, error) {
	if o == nil {
		return _nullBytes, nil
	}

	buf := bytes.NewBuffer(nil)
	buf.WriteRune('{')
	enc := json.NewEncoder(buf)
	for i, kv := range o {
		if i > 0 {
			buf.WriteRune(',')
		}
		if err := enc.Encode(kv.K); err != nil {
			return nil, err
		}
		buf.WriteRune(':')
		buf.Write(kv.V)
	}
	buf.WriteRune('}')
	return buf.Bytes(), nil
}

func (o *RawObject) UnmarshalJSON(b []byte) (err error) {
	*o, err = DecodeRawObject(json.NewDecoder(bytes.NewReader(b)))
	return
}

//

func DecodeRawObject(dec *json.Decoder) (r RawObject, err error) {
	token, err := dec.Token()
	if token == nil || err != nil {
		return nil, err
	}

	var delim json.Delim
	var ok bool
	if delim, ok = token.(json.Delim); !ok {
		return nil, errors.New("expected delim")
	}

	switch delim {
	case '{':
		var tok json.Token
		for {
			if tok, err = dec.Token(); err != nil {
				return
			}

			if tok, ok := tok.(json.Delim); ok {
				if tok == '}' {
					return
				}
				return nil, fmt.Errorf("unexpected delim: %v", tok)
			}
			key := tok.(string)

			var value json.RawMessage
			if err := dec.Decode(&value); err != nil {
				return nil, err
			}

			r = append(r, bt.KvOf(key, value))
		}

	case '[':
		var tok json.Token
		for {
			if tok, err = dec.Token(); err != nil {
				return
			}

			var ok bool
			if delim, ok = tok.(json.Delim); !ok {
				return nil, fmt.Errorf("expected delim: %v", tok)
			}
			if delim == ']' {
				return
			}
			if delim != '[' {
				return nil, fmt.Errorf("unexpected delim: %v", tok)
			}

			if tok, err = dec.Token(); err != nil {
				return
			}
			key, ok := tok.(string)
			if !ok {
				return nil, fmt.Errorf("expected string: %v", tok)
			}

			var value json.RawMessage
			if err := dec.Decode(&value); err != nil {
				return nil, err
			}

			if tok, err = dec.Token(); err != nil {
				return
			}
			if delim, ok = tok.(json.Delim); !ok || delim != ']' {
				return nil, fmt.Errorf("expected ]: %v", tok)
			}

			r = append(r, bt.KvOf(key, value))
		}

	default:
		return nil, fmt.Errorf("unexpected delim: %v", delim)
	}
}

//

func MakeRawObject[K, V any](kvs []bt.Kv[K, V]) (RawObject, error) {
	if kvs == nil {
		return nil, nil
	}

	o := make(RawObject, len(kvs))
	for i, kv := range kvs {
		kb, err := json.Marshal(kv.K)
		if err != nil {
			return nil, err
		}

		var k string
		switch kb[0] {
		case '-', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			k = string(kb)
		case '"':
			if err := json.Unmarshal(kb, &k); err != nil {
				return nil, err
			}
		default:
			return nil, fmt.Errorf("unhandled key bytes: %v", kb)
		}

		v, err := json.Marshal(kv.V)
		if err != nil {
			return nil, err
		}

		o[i] = bt.KvOf[string, json.RawMessage](k, v)
	}

	return o, nil
}

func UnmarshalRawObjectFields[K, V any](o RawObject) ([]bt.Kv[K, V], error) {
	if o == nil {
		return nil, nil
	}

	isNum := rfl.IsNumericType(rfl.TypeOf[K]())

	r := make([]bt.Kv[K, V], len(o))
	for i, kv := range o {
		ks := kv.K
		if !isNum {
			ks = "\"" + kv.K + "\""
		}

		var k K
		if err := json.Unmarshal([]byte(ks), &k); err != nil {
			return nil, err
		}

		var v V
		if err := json.Unmarshal(kv.V, &v); err != nil {
			return nil, err
		}

		r[i] = bt.KvOf(k, v)
	}

	return r, nil
}

func UnmarshalObjectFields[K, V any](b []byte) ([]bt.Kv[K, V], error) {
	var ro RawObject
	if err := json.Unmarshal(b, &ro); err != nil {
		return nil, err
	}

	return UnmarshalRawObjectFields[K, V](ro)
}
