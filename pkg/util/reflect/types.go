package reflect

import (
	"fmt"
	"reflect"
)

//

type Type interface {
	isType()
}

type BaseType struct {
	ty reflect.Type
}

func (t BaseType) isType() {}

func (t BaseType) Reflect() reflect.Type {
	return t.ty
}

//

type ScalarType interface {
	isScalarType()
}

type BaseScalarType struct {
	BaseType
}

func (t BaseScalarType) isScalarType() {}

//

type BoolType struct {
	BaseScalarType
}

//

type IntType struct {
	BaseScalarType
}

//

type UintType struct {
	BaseScalarType
}

//

type UintptrType struct {
	BaseScalarType
}

//

type FloatType struct {
	BaseScalarType
}

//

type ComplexType struct {
	BaseScalarType
}

//

type StringType struct {
	BaseScalarType
}

//

type ContainerType interface {
	isContainerType()
}

type BaseContainerType struct {
	BaseType
}

func (t BaseContainerType) isContainerType() {}

//

type SequenceType interface {
	isSequenceType()
}

type BaseSequenceType struct {
	BaseContainerType
}

func (t BaseSequenceType) isSequenceType() {}

//

type ArrayType struct {
	BaseSequenceType
}

//

type SliceType struct {
	BaseSequenceType
}

//

type MapType struct {
	BaseContainerType
}

//

type ChanType struct {
	BaseType
}

//

type FuncType struct {
	BaseType
}

//

type InterfaceType struct {
	BaseType
}

//

type PointerType struct {
	BaseType
}

//

type FieldType struct {
	s *StructType
	f reflect.StructField
}

type StructType struct {
	BaseType
}

//

type UnsafePointerType struct {
	BaseType
}

//

var (
	_boolType = BoolType{BaseScalarType{BaseType{ty: TypeOf[bool]()}}}

	_intType   = IntType{BaseScalarType{BaseType{ty: TypeOf[int]()}}}
	_int8Type  = IntType{BaseScalarType{BaseType{ty: TypeOf[int8]()}}}
	_int16Type = IntType{BaseScalarType{BaseType{ty: TypeOf[int16]()}}}
	_int32Type = IntType{BaseScalarType{BaseType{ty: TypeOf[int32]()}}}
	_int64Type = IntType{BaseScalarType{BaseType{ty: TypeOf[int64]()}}}

	_uintType   = UintType{BaseScalarType{BaseType{ty: TypeOf[uint]()}}}
	_uint8Type  = UintType{BaseScalarType{BaseType{ty: TypeOf[uint8]()}}}
	_uint16Type = UintType{BaseScalarType{BaseType{ty: TypeOf[uint16]()}}}
	_uint32Type = UintType{BaseScalarType{BaseType{ty: TypeOf[uint32]()}}}
	_uint64Type = UintType{BaseScalarType{BaseType{ty: TypeOf[uint64]()}}}

	_uintptrType = UintptrType{BaseScalarType{BaseType{ty: TypeOf[uintptr]()}}}

	_float32Type = FloatType{BaseScalarType{BaseType{ty: TypeOf[float32]()}}}
	_float64Type = FloatType{BaseScalarType{BaseType{ty: TypeOf[float64]()}}}

	_complex64Type  = ComplexType{BaseScalarType{BaseType{ty: TypeOf[complex64]()}}}
	_complex128Type = ComplexType{BaseScalarType{BaseType{ty: TypeOf[complex128]()}}}

	_stringType = StringType{BaseScalarType{BaseType{ty: TypeOf[string]()}}}
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
