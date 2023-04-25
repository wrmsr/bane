package box

//

type ComplexType interface {
	NumericType
	isComplexType()
}

type complexType struct {
	numericType
}

func (t complexType) isComplexType() {}

//

type Complex64Type struct{ complexType }
type Complex128Type struct{ complexType }

var (
	_complex64Type  = Complex64Type{complexType{numericType{scalarType{type_{t: typeOf[complex64]()}}}}}
	_complex128Type = Complex128Type{complexType{numericType{scalarType{type_{t: typeOf[complex128]()}}}}}
)

///

type Complex interface {
	Numeric
	isComplex()
}

type complex struct {
	numeric
}

func (v complex) isComplex() {}

//

type Complex64 struct{ complex }
type Complex128 struct{ complex }

func (v Complex64) Type() Type  { return _complex64Type }
func (v Complex128) Type() Type { return _complex128Type }
