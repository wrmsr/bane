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

func (v BaseType) isType() {}

func (v BaseType) Reflect() reflect.Type {
	return v.ty
}

//

type Scalar interface {
	isScalar()
}

type BaseScalar struct {
	BaseType
}

func (v BaseScalar) isScalar() {}

//

type Bool struct {
	BaseScalar
}

//

type Int struct {
	BaseScalar
}

//

type Uint struct {
	BaseScalar
}

//

type Uintptr struct {
	BaseScalar
}

//

type Float struct {
	BaseScalar
}

//

type Complex struct {
	BaseScalar
}

//

type String struct {
	BaseScalar
}

//

type Container interface {
	isContainer()
}

type BaseContainer struct {
	BaseType
}

func (v BaseContainer) isContainer() {}

//

type Sequence interface {
	isSequence()
}

type BaseSequence struct {
	BaseContainer
}

func (v BaseSequence) isSequence() {}

//

type Array struct {
	BaseSequence
}

//

type Slice struct {
	BaseSequence
}

//

type Map struct {
	BaseContainer
}

//

type Chan struct {
	BaseType
}

//

type Func struct {
	BaseType
}

//

type Interface struct {
	BaseType
}

//

type Pointer struct {
	BaseType
}

//

type Field struct {
	s *Struct
	f reflect.StructField
}

type Struct struct {
	BaseType
}

//

type UnsafePointer struct {
	BaseType
}

//

var (
	_boolType = Bool{BaseScalar{BaseType{ty: TypeOf[bool]()}}}

	_intType   = Int{BaseScalar{BaseType{ty: TypeOf[int]()}}}
	_int8Type  = Int{BaseScalar{BaseType{ty: TypeOf[int8]()}}}
	_int16Type = Int{BaseScalar{BaseType{ty: TypeOf[int16]()}}}
	_int32Type = Int{BaseScalar{BaseType{ty: TypeOf[int32]()}}}
	_int64Type = Int{BaseScalar{BaseType{ty: TypeOf[int64]()}}}

	_uintType   = Uint{BaseScalar{BaseType{ty: TypeOf[uint]()}}}
	_uint8Type  = Uint{BaseScalar{BaseType{ty: TypeOf[uint8]()}}}
	_uint16Type = Uint{BaseScalar{BaseType{ty: TypeOf[uint16]()}}}
	_uint32Type = Uint{BaseScalar{BaseType{ty: TypeOf[uint32]()}}}
	_uint64Type = Uint{BaseScalar{BaseType{ty: TypeOf[uint64]()}}}

	_uintptrType = Uintptr{BaseScalar{BaseType{ty: TypeOf[uintptr]()}}}

	_float32Type = Float{BaseScalar{BaseType{ty: TypeOf[float32]()}}}
	_float64Type = Float{BaseScalar{BaseType{ty: TypeOf[float64]()}}}

	_complex64Type  = Complex{BaseScalar{BaseType{ty: TypeOf[complex64]()}}}
	_complex128Type = Complex{BaseScalar{BaseType{ty: TypeOf[complex128]()}}}

	_stringType = String{BaseScalar{BaseType{ty: TypeOf[string]()}}}
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
		return Array{BaseSequence{BaseContainer{BaseType{ty: ty}}}}
	case reflect.Slice:
		return Slice{BaseSequence{BaseContainer{BaseType{ty: ty}}}}

	case reflect.Map:
		return Map{BaseContainer{BaseType{ty: ty}}}

	case reflect.Chan:
		return Chan{BaseType{ty: ty}}
	case reflect.Func:
		return Func{BaseType{ty: ty}}
	case reflect.Interface:
		return Interface{BaseType{ty: ty}}
	case reflect.Pointer:
		return Pointer{BaseType{ty: ty}}
	case reflect.Struct:
		return Struct{BaseType{ty: ty}}
	case reflect.UnsafePointer:
		return UnsafePointer{BaseType{ty: ty}}
	}

	panic(fmt.Errorf("unhandled type: %v", ty))
}
