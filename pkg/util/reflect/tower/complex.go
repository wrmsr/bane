package tower

//

type ComplexType struct {
	BaseNumericType
}

var (
	_complex64Type  = ComplexType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[complex64]()}}}}
	_complex128Type = ComplexType{BaseNumericType{BaseScalarType{BaseType{ty: typeOf[complex128]()}}}}
)

//

type Complex struct {
	BaseNumeric
}
