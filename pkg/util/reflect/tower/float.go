package tower

//

type FloatType struct {
	scalarType
}

type Float32Type struct{ FloatType }
type Float64Type struct{ FloatType }

var (
	_float32Type = Float32Type{FloatType{scalarType{type_{ty: typeOf[float32]()}}}}
	_float64Type = Float64Type{FloatType{scalarType{type_{ty: typeOf[float64]()}}}}
)

//

type Float struct {
	scalar
}
