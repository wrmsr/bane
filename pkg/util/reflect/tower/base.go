/*
Scalar
 Bool
 Numeric
  Integral
   Unsigned
    Uint
    Uint8
    Uint16
    Uint32
    Uint64
    Uintptr
   Signed
    Int
    Int8
    Int16
    Int32
    Int64
  Float
   Float32
   Float64
  Complex
   Complex64
   Complex128
 String
Container
 Sequence
  Array
  Slice
 Map
Chan
Func
Interface
Pointer
Struct
UnsafePointer
*/
package tower

import "reflect"

//

type Type interface {
	isType()
}

type type_ struct {
	ty reflect.Type
}

func (t type_) isType() {}

func (t type_) Reflect() reflect.Type {
	return t.ty
}

//

type Value interface {
	isValue()

	Type() Type
}

type value struct {
	v reflect.Value
}

func (v value) isValue() {}

func (v value) Reflect() reflect.Value {
	return v.v
}

//

func typeOf[T any]() reflect.Type {
	var z T
	return reflect.TypeOf(z)
}
