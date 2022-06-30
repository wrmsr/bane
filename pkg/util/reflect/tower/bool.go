package tower

//

type BoolType struct {
	BaseScalarType
}

var _boolType = BoolType{BaseScalarType{BaseType{ty: typeOf[bool]()}}}

//

type Bool struct {
	BaseScalar
}

func (v Bool) Type() Type { return _boolType }
