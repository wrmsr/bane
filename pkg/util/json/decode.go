package json

import (
	"encoding/json"
	"errors"
	"fmt"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

type RawObject []bt.Kv[string, json.RawMessage]

func readDelim(dec *json.Decoder) (json.Delim, error) {
	token, err := dec.Token()
	if err != nil {
		return 0, err
	}
	if delim, ok := token.(json.Delim); ok {
		return delim, nil
	}
	return 0, errors.New("expected delim")
}

func endDelim(d json.Delim) json.Delim {
	switch d {
	case '{':
		return '}'
	case '[':
		return ']'
	}
	panic(errors.New("unreachable"))
}

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
