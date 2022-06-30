package tower

//

type SignedType struct {
	integralType
}

var (
	_intType   = SignedType{integralType{numericType{scalarType{type_{ty: typeOf[int]()}}}}}
	_int8Type  = SignedType{integralType{numericType{scalarType{type_{ty: typeOf[int8]()}}}}}
	_int16Type = SignedType{integralType{numericType{scalarType{type_{ty: typeOf[int16]()}}}}}
	_int32Type = SignedType{integralType{numericType{scalarType{type_{ty: typeOf[int32]()}}}}}
	_int64Type = SignedType{integralType{numericType{scalarType{type_{ty: typeOf[int64]()}}}}}
)

//

type Signed struct {
	integral
}
