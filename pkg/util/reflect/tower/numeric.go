package tower

//

type NumericType interface {
	ScalarType
	isNumericType()
}

type numericType struct {
	scalarType
}

func (t numericType) isNumericType() {}

//

type Numeric interface {
	Scalar
	isNumeric()
}

type numeric struct {
	scalar
}

func (t numeric) isNumeric() {}
