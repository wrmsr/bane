package box

///

type FloatType interface {
	NumericType
	isFloatType()
}

type floatType struct {
	numericType
}

func (t floatType) isFloatType() {}

//

type Float32Type struct{ floatType }
type Float64Type struct{ floatType }

var (
	_float32Type = Float32Type{floatType{numericType{scalarType{type_{t: typeOf[float32]()}}}}}
	_float64Type = Float64Type{floatType{numericType{scalarType{type_{t: typeOf[float64]()}}}}}
)

///

type Float interface {
	Numeric
	isFloat()
}

type float struct {
	numeric
}

func (v float) isFloat() {}

//

type Float32 struct{ float }
type Float64 struct{ float }

func (v Float32) Type() Type { return _float32Type }
func (v Float64) Type() Type { return _float64Type }
