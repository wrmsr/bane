package tower

//

type UnsignedType struct {
	integralType
}

var (
	_uintType    = UnsignedType{integralType{numericType{scalarType{type_{ty: typeOf[uint]()}}}}}
	_uint8Type   = UnsignedType{integralType{numericType{scalarType{type_{ty: typeOf[uint8]()}}}}}
	_uint16Type  = UnsignedType{integralType{numericType{scalarType{type_{ty: typeOf[uint16]()}}}}}
	_uint32Type  = UnsignedType{integralType{numericType{scalarType{type_{ty: typeOf[uint32]()}}}}}
	_uint64Type  = UnsignedType{integralType{numericType{scalarType{type_{ty: typeOf[uint64]()}}}}}
	_uintptrType = UnsignedType{integralType{numericType{scalarType{type_{ty: typeOf[uintptr]()}}}}}
)

//

type Unsigned struct {
	integral
}
