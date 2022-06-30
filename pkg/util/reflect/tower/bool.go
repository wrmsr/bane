package tower

//

type BoolType struct {
	scalarType
}

var _boolType = BoolType{scalarType{type_{ty: typeOf[bool]()}}}

//

type Bool struct {
	scalar
}

func (v Bool) Type() Type { return _boolType }
