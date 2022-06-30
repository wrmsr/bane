package tower

//

type NumericType interface {
	isNumericType()
}

type numericType struct {
	scalarType
}

func (t numericType) isNumericType() {}

//

type Numeric interface {
	isNumeric()
}

type numeric struct {
	scalar
}

func (t numeric) isNumeric() {}
