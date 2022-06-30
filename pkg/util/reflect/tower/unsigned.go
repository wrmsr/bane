package tower

//

type UnsignedType struct {
	BaseIntegralType
}

var (
	_uintType    = UnsignedType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[uint]()}}}}}
	_uint8Type   = UnsignedType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[uint8]()}}}}}
	_uint16Type  = UnsignedType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[uint16]()}}}}}
	_uint32Type  = UnsignedType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[uint32]()}}}}}
	_uint64Type  = UnsignedType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[uint64]()}}}}}
	_uintptrType = UnsignedType{BaseIntegralType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[uintptr]()}}}}}
)

//

type Unsigned struct {
	BaseIntegral
}
