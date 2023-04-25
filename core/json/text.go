package json

import (
	"encoding"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	rfl "github.com/wrmsr/bane/core/reflect"
)

//

func IsPrimitiveTextType(ty reflect.Type) bool {
	switch ty.Kind() {
	case
		reflect.String,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return true
	}
	return false
}

//

var _textMarshalerTy = rfl.TypeOf[encoding.TextMarshaler]()

func CanMarshalAsText(ty reflect.Type) bool {
	return IsPrimitiveTextType(ty) || ty.AssignableTo(_textMarshalerTy)
}

func MarshalAsText(v reflect.Value) (string, error) {
	if v.Kind() == reflect.String {
		return v.String(), nil
	}

	if tm, ok := v.Interface().(encoding.TextMarshaler); ok {
		if v.Kind() == reflect.Pointer && v.IsNil() {
			return "", nil
		}
		buf, err := tm.MarshalText()
		return string(buf), err
	}

	switch v.Kind() {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10), nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10), nil

	}

	return "", fmt.Errorf("unhandled text type: %s", v.Type().Name())
}

//

var _textUnmarshalerTy = rfl.TypeOf[encoding.TextUnmarshaler]()

func CanUnmarshalAsText(ty reflect.Type) bool {
	return IsPrimitiveTextType(ty) || ty.AssignableTo(_textUnmarshalerTy)
}

func UnmarshalAsText(s string, v reflect.Value) error {
	if v.Kind() == reflect.String {
		v.Set(reflect.ValueOf(s))
		return nil
	}

	if tm, ok := v.Interface().(encoding.TextUnmarshaler); ok {
		return tm.UnmarshalText([]byte(s))
	}

	switch v.Kind() {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		var i int64
		if err := json.Unmarshal([]byte(s), &i); err != nil {
			return err
		}
		v.SetInt(i)
		return nil

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		var u uint64
		if err := json.Unmarshal([]byte(s), &u); err != nil {
			return err
		}
		v.SetUint(u)
		return nil

	}

	return fmt.Errorf("unhandled text type: %s", v.Type().Name())
}
