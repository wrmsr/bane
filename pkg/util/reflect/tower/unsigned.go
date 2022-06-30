package tower

//

type UintType struct {
	BaseIntegralType
}

var (
	_uintType    = UintType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[uint]()}}}}}
	_uint8Type   = UintType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[uint8]()}}}}}
	_uint16Type  = UintType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[uint16]()}}}}}
	_uint32Type  = UintType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[uint32]()}}}}}
	_uint64Type  = UintType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[uint64]()}}}}}
	_uintptrType = UintType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[uintptr]()}}}}}
)

//

type Uint struct {
	BaseIntegral
}
