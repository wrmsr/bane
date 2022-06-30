package tower

//

type IntType struct {
	BaseIntegralType
}

var (
	_intType   = IntType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[int]()}}}}}
	_int8Type  = IntType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[int8]()}}}}}
	_int16Type = IntType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[int16]()}}}}}
	_int32Type = IntType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[int32]()}}}}}
	_int64Type = IntType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[int64]()}}}}}
)

//

type Int struct {
	BaseIntegral
}
