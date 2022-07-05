package json

import (
	"encoding"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

//

var (
	_nullBytes      = []byte("null")
	_trueBytes      = []byte("true")
	_falseBytes     = []byte("false")
	_quoteBytes     = []byte("\"")
	_backslashBytes = []byte("\\")
)

func NullBytes() []byte      { return _nullBytes }
func TrueBytes() []byte      { return _trueBytes }
func FalseBytes() []byte     { return _falseBytes }
func QuoteBytes() []byte     { return _quoteBytes }
func BackslashBytes() []byte { return _backslashBytes }

func IsNullBytes(b []byte) bool {
	return len(b) == 4 && b[0] == 'n' && b[1] == 'u' && b[2] == 'l' && b[3] == 'l'
}

func IsTrueBytes(b []byte) bool {
	return len(b) == 4 && b[0] == 't' && b[1] == 'r' && b[2] == 'u' && b[3] == 'e'
}

func IsFalseBytes(b []byte) bool {
	return len(b) == 5 && b[0] == 'f' && b[1] == 'a' && b[2] == 'l' && b[3] == 's' && b[4] == 'e'
}

func IsQuoteBytes(b []byte) bool {
	return len(b) == 1 && b[0] == '"'
}

func IsBackslashBytes(b []byte) bool {
	return len(b) == 1 && b[0] == '\\'
}

func BoolBytes(v bool) []byte {
	if v {
		return _trueBytes
	} else {
		return _falseBytes
	}
}

//

func MarshalString(v any) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func MarshalIndentString(v any, prefix, indent string) (string, error) {
	b, err := json.MarshalIndent(v, prefix, indent)
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

//

func MarshalAsKey(k reflect.Value) (string, error) {
	if k.Kind() == reflect.String {
		return k.String(), nil
	}
	if tm, ok := k.Interface().(encoding.TextMarshaler); ok {
		if k.Kind() == reflect.Pointer && k.IsNil() {
			return "", nil
		}
		buf, err := tm.MarshalText()
		return string(buf), err
	}
	switch k.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(k.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(k.Uint(), 10), nil
	}
	return "", fmt.Errorf("unhandled key type: %s", k.Type().Name())
}
