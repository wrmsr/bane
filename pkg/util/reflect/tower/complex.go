package tower

//

type ComplexType struct {
	BaseScalarType
}

var (
	_complex64Type  = ComplexType{BaseScalarType{BaseType{ty: typeOf[complex64]()}}}
	_complex128Type = ComplexType{BaseScalarType{BaseType{ty: typeOf[complex128]()}}}
)

//

type Complex struct {
	BaseScalar
}
