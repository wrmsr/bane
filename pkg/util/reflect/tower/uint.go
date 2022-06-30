package tower

//

type UintType struct {
	BaseScalarType
}

var (
	_uintType    = UintType{BaseScalarType{BaseType{ty: typeOf[uint]()}}}
	_uint8Type   = UintType{BaseScalarType{BaseType{ty: typeOf[uint8]()}}}
	_uint16Type  = UintType{BaseScalarType{BaseType{ty: typeOf[uint16]()}}}
	_uint32Type  = UintType{BaseScalarType{BaseType{ty: typeOf[uint32]()}}}
	_uint64Type  = UintType{BaseScalarType{BaseType{ty: typeOf[uint64]()}}}
	_uintptrType = UintType{BaseScalarType{BaseType{ty: typeOf[uintptr]()}}}
)

//

type Uint struct {
	BaseScalar
}
