package tower

import (
	"fmt"
	"reflect"
)

func MakeType(ty reflect.Type) Type {
	switch ty.Kind() {

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
		return ArrayType{sequenceType{containerType{type_{ty: ty}}}}
	case reflect.Slice:
		return SliceType{sequenceType{containerType{type_{ty: ty}}}}

	case reflect.Map:
		return MapType{containerType{type_{ty: ty}}}

	case reflect.Chan:
		return ChanType{type_{ty: ty}}
	case reflect.Func:
		return FuncType{type_{ty: ty}}
	case reflect.Interface:
		return InterfaceType{type_{ty: ty}}
	case reflect.Pointer:
		return PointerType{type_{ty: ty}}
	case reflect.Struct:
		return StructType{type_{ty: ty}}
	case reflect.UnsafePointer:
		return UnsafePointerType{type_{ty: ty}}
	}

	panic(fmt.Errorf("unhandled type: %v", ty))
}

//

func MakeValue(v reflect.Value) Value {
	switch v.Kind() {

	case reflect.Bool:
		return Bool{scalar{value{v: v}}}

	case
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		return Signed{scalar{value{v: v}}}

	case
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr:
		return Unsigned{scalar{value{v: v}}}

	case
		reflect.Float32,
		reflect.Float64:
		return Float{scalar{value{v: v}}}

	case
		reflect.Complex64,
		reflect.Complex128:
		return Complex{scalar{value{v: v}}}

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
