package box

//

type ScalarType interface {
	Type
	isScalarType()
}

type scalarType struct {
	type_
}

func (t scalarType) isScalarType() {}

//

type Scalar interface {
	Value
	isScalar()
}

type scalar struct {
	value
}

func (v scalar) isScalar() {}
