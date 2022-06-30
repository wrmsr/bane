package box

import (
	"fmt"
	"reflect"
)

func BoxType(t reflect.Type) Type {
	switch t.Kind() {

	case reflect.Bool:
		return _boolType

	case reflect.Int:
		return _intType
	case reflect.Int8:
		return _int8Type
	case reflect.Int16:
		return _int16Type
	case reflect.Int32:
		return _int32Type
	case reflect.Int64:
		return _int64Type

	case reflect.Uint:
		return _uintType
	case reflect.Uint8:
		return _uint8Type
	case reflect.Uint16:
		return _uint16Type
	case reflect.Uint32:
		return _uint32Type
	case reflect.Uint64:
		return _uint64Type
	case reflect.Uintptr:
		return _uintptrType

	case reflect.Float32:
		return _float32Type
	case reflect.Float64:
		return _float64Type

	case reflect.Complex64:
		return _complex64Type
	case reflect.Complex128:
		return _complex128Type

	case reflect.String:
		return _stringType

	case reflect.Array:
		return ArrayType{sequenceType{containerType{type_{t: t}}}}
	case reflect.Slice:
		return SliceType{sequenceType{containerType{type_{t: t}}}}

	case reflect.Map:
		return MapType{containerType{type_{t: t}}}

	case reflect.Chan:
		return ChanType{type_{t: t}}
	case reflect.Func:
		return FuncType{type_{t: t}}
	case reflect.Interface:
		return InterfaceType{type_{t: t}}
	case reflect.Pointer:
		return PointerType{type_{t: t}}
	case reflect.Struct:
		return StructType{type_{t: t}}
	case reflect.UnsafePointer:
		return UnsafePointerType{type_{t: t}}

	}
	panic(fmt.Errorf("unhandled type: %v", t))
}

//

func Box(v reflect.Value) Value {
	switch v.Kind() {

	case reflect.Bool:
		return Bool{scalar{value{v: v}}}

	case reflect.Int:
		return Int{signed{integral{numeric{scalar{value{v: v}}}}}}
	case reflect.Int8:
		return Int8{signed{integral{numeric{scalar{value{v: v}}}}}}
	case reflect.Int16:
		return Int16{signed{integral{numeric{scalar{value{v: v}}}}}}
	case reflect.Int32:
		return Int32{signed{integral{numeric{scalar{value{v: v}}}}}}
	case reflect.Int64:
		return Int64{signed{integral{numeric{scalar{value{v: v}}}}}}

	case reflect.Uint:
		return Uint{unsigned{integral{numeric{scalar{value{v: v}}}}}}
	case reflect.Uint8:
		return Uint8{unsigned{integral{numeric{scalar{value{v: v}}}}}}
	case reflect.Uint16:
		return Uint16{unsigned{integral{numeric{scalar{value{v: v}}}}}}
	case reflect.Uint32:
		return Uint32{unsigned{integral{numeric{scalar{value{v: v}}}}}}
	case reflect.Uint64:
		return Uint64{unsigned{integral{numeric{scalar{value{v: v}}}}}}
	case reflect.Uintptr:
		return Uintptr{unsigned{integral{numeric{scalar{value{v: v}}}}}}

	case reflect.Float32:
		return Float32{float{numeric{scalar{value{v: v}}}}}
	case reflect.Float64:
		return Float64{float{numeric{scalar{value{v: v}}}}}

	case reflect.Complex64:
		return Complex64{complex{numeric{scalar{value{v: v}}}}}
	case reflect.Complex128:
		return Complex128{complex{numeric{scalar{value{v: v}}}}}

	case reflect.String:
		return String{scalar{value{v: v}}}

	case reflect.Array:
		return Array{sequence{container{value{v: v}}}}
	case reflect.Slice:
		return Slice{sequence{container{value{v: v}}}}

	case reflect.Map:
		return Map{container{value{v: v}}}

	case reflect.Chan:
		return Chan{value{v: v}}
	case reflect.Func:
		return Func{value{v: v}}
	case reflect.Interface:
		return Interface{value{v: v}}
	case reflect.Pointer:
		return Pointer{value{v: v}}
	case reflect.Struct:
		return Struct{value{v: v}}
	case reflect.UnsafePointer:
		return UnsafePointer{value{v: v}}

	}
	panic(fmt.Errorf("unhandled value: %v", v))
}
