package json

import (
	"encoding/json"
	"errors"
)

//

var (
	_nullBytes        = []byte("null")
	_trueBytes        = []byte("true")
	_falseBytes       = []byte("false")
	_quoteBytes       = []byte("\"")
	_backslashBytes   = []byte("\\")
	_emptyObjectBytes = []byte("{}")
)

func NullBytes() []byte        { return _nullBytes }
func TrueBytes() []byte        { return _trueBytes }
func FalseBytes() []byte       { return _falseBytes }
func QuoteBytes() []byte       { return _quoteBytes }
func BackslashBytes() []byte   { return _backslashBytes }
func EmptyObjectBytes() []byte { return _emptyObjectBytes }

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

func MarshalPretty(v any) (string, error) {
	return MarshalIndentString(v, "", "  ")
}

//

func UnmarshalAs[T any](b []byte) (v T, err error) {
	err = json.Unmarshal(b, &v)
	return
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
