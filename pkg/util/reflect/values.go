package reflect

import (
	"fmt"
	"reflect"
)

//

type Value interface {
	isValue()
}

type value struct {
	v reflect.Value
}

func (v value) isValue() {}

//

type Scalar interface {
	isScalar()
}

type scalar struct {
	value
}

func (v scalar) isScalar() {}

//

type Bool struct {
	scalar
}

//

type Int struct {
	scalar
}

//

type Uint struct {
	scalar
}

//

type Uintptr struct {
	scalar
}

//

type Float struct {
	scalar
}

//

type Complex struct {
	scalar
}

//

type String struct {
	scalar
}

//

type Container interface {
	isContainer()
}

type container struct {
	value
}

func (v container) isContainer() {}

//

type Sequence interface {
	isSequence()
}

type sequence struct {
	container
}

func (v sequence) isSequence() {}

//

type Array struct {
	sequence
}

//

type Slice struct {
	sequence
}

//

type Map struct {
	container
}

//

type Chan struct {
	value
}

//

type Func struct {
	value
}

//

type Interface struct {
	value
}

//

type Pointer struct {
	value
}

//

type Struct struct {
	value
}

//

type UnsafePointer struct {
	value
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
		return Int{scalar{value{v: v}}}
	case
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		return Uint{scalar{value{v: v}}}
	case reflect.Uintptr:
		return Uintptr{scalar{value{v: v}}}
	case
		reflect.Float32,
		reflect.Float64:
		return Float{scalar{value{v: v}}}
	case
		reflect.Complex64,
		reflect.Complex128:
		return Complex{scalar{value{v: v}}}
	case reflect.Array:
		return Array{sequence{container{value{v: v}}}}
	case reflect.Chan:
		return Chan{value{v: v}}
	case reflect.Func:
		return Func{value{v: v}}
	case reflect.Interface:
		return Interface{value{v: v}}
	case reflect.Map:
		return Map{container{value{v: v}}}
	case reflect.Pointer:
		return Pointer{value{v: v}}
	case reflect.Slice:
		return Slice{sequence{container{value{v: v}}}}
	case reflect.String:
		return String{scalar{value{v: v}}}
	case reflect.Struct:
		return Struct{value{v: v}}
	case reflect.UnsafePointer:
		return UnsafePointer{value{v: v}}
	}
	panic(fmt.Errorf("unhandled value: %v", v))
}
