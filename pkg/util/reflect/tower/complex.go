package tower

//

type ComplexType struct {
	numericType
}

var (
	_complex64Type  = ComplexType{numericType{scalarType{type_{ty: typeOf[complex64]()}}}}
	_complex128Type = ComplexType{numericType{scalarType{type_{ty: typeOf[complex128]()}}}}
)

//

type Complex struct {
	numeric
}
