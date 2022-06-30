package tower

import "reflect"

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

type Value interface {
	isValue()

	Type() Type
}

type BaseValue struct {
	v reflect.Value
}

func (v BaseValue) isValue() {}

func (v BaseValue) Reflect() reflect.Value {
	return v.v
}

//

func typeOf[T any]() reflect.Type {
	var z T
	return reflect.TypeOf(z)
}
