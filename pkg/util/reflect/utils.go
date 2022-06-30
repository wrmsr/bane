package reflect

import (
	"fmt"
	"reflect"
)

//

func AsValue(v any) reflect.Value {
	if v, ok := v.(reflect.Value); ok {
		return v
	}

	return reflect.ValueOf(v)
}

func AsType(v any) reflect.Type {
	if v, ok := v.(reflect.Type); ok {
		return v
	}

	rv := AsValue(v)
	if rv.Kind() == reflect.Pointer && rv.IsNil() {
		return rv.Type().Elem()
	}

	panic(fmt.Errorf("cannot make type from %T %v", v, v))
}

func TypeOf[T any]() reflect.Type {
	var z T
	return reflect.TypeOf(z)
}

//

func UnwrapPointerValue(v reflect.Value) (reflect.Value, bool) {
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return v, false
		}
		v = v.Elem()
	}
	return v, true
}

func UnwrapPointerType(ty reflect.Type) reflect.Type {
	for ty.Kind() == reflect.Pointer {
		ty = ty.Elem()
	}
	return ty
}

//

func IsEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case
		reflect.Array,
		reflect.Map,
		reflect.Slice,
		reflect.String:
		return v.Len() == 0
	case
		reflect.Bool:
		return !v.Bool()
	case
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		return v.Int() == 0
	case
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr:
		return v.Uint() == 0
	case
		reflect.Float32,
		reflect.Float64:
		return v.Float() == 0
	case
		reflect.Interface,
		reflect.Pointer:
		return v.IsNil()
	}
	return false
}

//

func IsNumericType(ty reflect.Type) bool {
	switch ty.Kind() {
	case
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64,
		reflect.Complex64,
		reflect.Complex128:
		return true
	}
	return false
}
