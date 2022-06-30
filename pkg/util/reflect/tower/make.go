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
		return ArrayType{BaseSequenceType{BaseContainerType{BaseType{ty: ty}}}}
	case reflect.Slice:
		return SliceType{BaseSequenceType{BaseContainerType{BaseType{ty: ty}}}}

	case reflect.Map:
		return MapType{BaseContainerType{BaseType{ty: ty}}}

	case reflect.Chan:
		return ChanType{BaseType{ty: ty}}
	case reflect.Func:
		return FuncType{BaseType{ty: ty}}
	case reflect.Interface:
		return InterfaceType{BaseType{ty: ty}}
	case reflect.Pointer:
		return PointerType{BaseType{ty: ty}}
	case reflect.Struct:
		return StructType{BaseType{ty: ty}}
	case reflect.UnsafePointer:
		return UnsafePointerType{BaseType{ty: ty}}
	}

	panic(fmt.Errorf("unhandled type: %v", ty))
}

//

func MakeValue(v reflect.Value) Value {
	switch v.Kind() {

	case reflect.Bool:
		return Bool{BaseScalar{BaseValue{v: v}}}

	case
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		return Signed{BaseScalar{BaseValue{v: v}}}

	case
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr:
		return Unsigned{BaseScalar{BaseValue{v: v}}}

	case
		reflect.Float32,
		reflect.Float64:
		return Float{BaseScalar{BaseValue{v: v}}}

	case
		reflect.Complex64,
		reflect.Complex128:
		return Complex{BaseScalar{BaseValue{v: v}}}

	case reflect.String:
		return String{BaseScalar{BaseValue{v: v}}}

	case reflect.Array:
		return Array{BaseSequence{BaseContainer{BaseValue{v: v}}}}
	case reflect.Slice:
		return Slice{BaseSequence{BaseContainer{BaseValue{v: v}}}}

	case reflect.Map:
		return Map{BaseContainer{BaseValue{v: v}}}

	case reflect.Chan:
		return Chan{BaseValue{v: v}}
	case reflect.Func:
		return Func{BaseValue{v: v}}
	case reflect.Interface:
		return Interface{BaseValue{v: v}}
	case reflect.Pointer:
		return Pointer{BaseValue{v: v}}
	case reflect.Struct:
		return Struct{BaseValue{v: v}}
	case reflect.UnsafePointer:
		return UnsafePointer{BaseValue{v: v}}
	}

	panic(fmt.Errorf("unhandled value: %v", v))
}
