package tower

//

type IntType struct {
	BaseScalarType
}

var (
	_intType   = IntType{BaseScalarType{BaseType{ty: typeOf[int]()}}}
	_int8Type  = IntType{BaseScalarType{BaseType{ty: typeOf[int8]()}}}
	_int16Type = IntType{BaseScalarType{BaseType{ty: typeOf[int16]()}}}
	_int32Type = IntType{BaseScalarType{BaseType{ty: typeOf[int32]()}}}
	_int64Type = IntType{BaseScalarType{BaseType{ty: typeOf[int64]()}}}
)

//

type Int struct {
	BaseScalar
}
