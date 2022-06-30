package tower

//

type NumericType interface {
	isNumericType()
}

type BaseNumericType struct {
	BaseScalarType
}

func (t BaseNumericType) isNumericType() {}

//

type Numeric interface {
	isNumeric()
}

type BaseNumeric struct {
	BaseScalar
}

func (t BaseNumeric) isNumeric() {}
