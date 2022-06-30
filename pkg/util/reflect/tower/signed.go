package tower

//

type SignedType struct {
	BaseIntegralType
}

var (
	_intType   = SignedType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[int]()}}}}}
	_int8Type  = SignedType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[int8]()}}}}}
	_int16Type = SignedType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[int16]()}}}}}
	_int32Type = SignedType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[int32]()}}}}}
	_int64Type = SignedType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[int64]()}}}}}
)

//

type Signed struct {
	BaseIntegral
}
