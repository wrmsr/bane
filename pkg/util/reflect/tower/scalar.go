package tower

//

type ScalarType interface {
	isScalarType()
}

type BaseScalarType struct {
	BaseType
}

func (t BaseScalarType) isScalarType() {}

//

type Scalar interface {
	isScalar()
}

type BaseScalar struct {
	BaseValue
}

func (v BaseScalar) isScalar() {}
