package marshal

import (
	"bytes"
	"io"
	"strconv"

	iou "github.com/wrmsr/bane/pkg/util/io"
	ju "github.com/wrmsr/bane/pkg/util/json"
)

func JsonEncode(w io.Writer, v Value) error {
	switch v := v.(type) {

	case Null:
		return iou.WriteErr(w, ju.NullBytes())

	case Bool:
		return iou.WriteErr(w, ju.BoolBytes(v.v))

	case Int:
		if v.u {
			return iou.WriteErr(w, []byte(strconv.FormatUint(uint64(v.v), 10)))
		} else {
			return iou.WriteErr(w, []byte(strconv.FormatInt(v.v, 10)))
		}

	case Float:
		return ju.EncodeFloat64(w, v.v)

	case Number:
		panic("nyi")

	case String:
		return ju.EncodeString(w, v.v, false)

	case Bytes:
		return _unhandledType

	case Array:
		if _, err := w.Write([]byte{'['}); err != nil {
			return err
		}
		for i, e := range v.v {
			if i > 0 {
				if _, err := w.Write([]byte{','}); err != nil {
					return err
				}
			}
			if err := JsonEncode(w, e); err != nil {
				return err
			}
		}
		if _, err := w.Write([]byte{']'}); err != nil {
			return err
		}
		return nil

	case Object:
		return _unhandledType

	case Any:
		return _unhandledType

	}
	panic("unreachable")
}

func JsonMarshal(v Value) ([]byte, error) {
	b := bytes.NewBuffer(nil)
	if err := JsonEncode(b, v); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (v Null) MarshalJSON() ([]byte, error)   { return JsonMarshal(v) }
func (v Bool) MarshalJSON() ([]byte, error)   { return JsonMarshal(v) }
func (v Int) MarshalJSON() ([]byte, error)    { return JsonMarshal(v) }
func (v Float) MarshalJSON() ([]byte, error)  { return JsonMarshal(v) }
func (v Number) MarshalJSON() ([]byte, error) { return JsonMarshal(v) }
func (v String) MarshalJSON() ([]byte, error) { return JsonMarshal(v) }
func (v Bytes) MarshalJSON() ([]byte, error)  { return JsonMarshal(v) }
func (v Array) MarshalJSON() ([]byte, error)  { return JsonMarshal(v) }
func (v Object) MarshalJSON() ([]byte, error) { return JsonMarshal(v) }
func (v Any) MarshalJSON() ([]byte, error)    { return JsonMarshal(v) }
