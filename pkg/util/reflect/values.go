package reflect

import (
	"fmt"
	"reflect"
)

//

type Value interface {
	isValue()
}

type BaseValue struct {
	v reflect.Value
}

func (v BaseValue) isValue() {}

func (v BaseValue) Reflect() reflect.Value {
	return v.v
}

//

type Scalar interface {
	isScalar()
}

type BaseScalar struct {
	BaseValue
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
	BaseValue
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
	BaseValue
}

//

type Func struct {
	BaseValue
}

//

type Interface struct {
	BaseValue
}

//

type Pointer struct {
	BaseValue
}

//

type Struct struct {
	BaseValue
}

//

type UnsafePointer struct {
	BaseValue
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
		return Int{BaseScalar{BaseValue{v: v}}}
	case
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		return Uint{BaseScalar{BaseValue{v: v}}}
	case reflect.Uintptr:
		return Uintptr{BaseScalar{BaseValue{v: v}}}
	case
		reflect.Float32,
		reflect.Float64:
		return Float{BaseScalar{BaseValue{v: v}}}
	case
		reflect.Complex64,
		reflect.Complex128:
		return Complex{BaseScalar{BaseValue{v: v}}}
	case reflect.Array:
		return Array{BaseSequence{BaseContainer{BaseValue{v: v}}}}
	case reflect.Chan:
		return Chan{BaseValue{v: v}}
	case reflect.Func:
		return Func{BaseValue{v: v}}
	case reflect.Interface:
		return Interface{BaseValue{v: v}}
	case reflect.Map:
		return Map{BaseContainer{BaseValue{v: v}}}
	case reflect.Pointer:
		return Pointer{BaseValue{v: v}}
	case reflect.Slice:
		return Slice{BaseSequence{BaseContainer{BaseValue{v: v}}}}
	case reflect.String:
		return String{BaseScalar{BaseValue{v: v}}}
	case reflect.Struct:
		return Struct{BaseValue{v: v}}
	case reflect.UnsafePointer:
		return UnsafePointer{BaseValue{v: v}}
	}
	panic(fmt.Errorf("unhandled value: %v", v))
}
