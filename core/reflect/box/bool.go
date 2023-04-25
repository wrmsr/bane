package box

import "reflect"

//

type BoolType struct {
	scalarType
}

var _boolType = BoolType{scalarType{type_{t: typeOf[bool]()}}}

//

type Bool struct {
	scalar
}

func (v Bool) Type() Type { return _boolType }

var (
	_trueValue  = Bool{scalar{value{v: reflect.ValueOf(true)}}}
	_falseValue = Bool{scalar{value{v: reflect.ValueOf(false)}}}
)
