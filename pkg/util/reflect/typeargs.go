package reflect

import (
	"fmt"
	"reflect"
)

type TypeArgsReflector interface {
	ReflectTypeArgs() []reflect.Type
}

var _typeArgsReflectorTy = TypeOf[TypeArgsReflector]()

func TypeArgs(ty reflect.Type) []reflect.Type {
	if ty.AssignableTo(_typeArgsReflectorTy) {
		inst := reflect.New(ty).Elem().Interface()
		if r, ok := inst.(TypeArgsReflector); ok {
			return r.ReflectTypeArgs()
		}
		panic(fmt.Errorf("can't reflect: %s", ty))
	}

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
