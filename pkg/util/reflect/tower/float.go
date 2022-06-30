package tower

//

type FloatType struct {
	BaseScalarType
}

var (
	_float32Type = FloatType{BaseScalarType{BaseType{ty: typeOf[float32]()}}}
	_float64Type = FloatType{BaseScalarType{BaseType{ty: typeOf[float64]()}}}
)

//

type Float struct {
	BaseScalar
}
