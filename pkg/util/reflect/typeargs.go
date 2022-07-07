package reflect

import (
	"fmt"
	"reflect"
)

func TypeArgs(ty reflect.Type) []reflect.Type {
	switch ty.Kind() {

	case
		reflect.Bool,
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
		reflect.Complex128,
		reflect.String,
		reflect.UnsafePointer:
		return nil

	case
		reflect.Array,
		reflect.Slice,
		reflect.Chan,
		reflect.Pointer:
		return []reflect.Type{ty.Elem()}

	case reflect.Map:
		return []reflect.Type{ty.Key(), ty.Elem()}

	}

	panic(fmt.Errorf("unhandled type: %s", ty))
}
