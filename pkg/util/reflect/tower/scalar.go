package tower

//

type ScalarType interface {
	isScalarType()
}

type scalarType struct {
	type_
}

func (t scalarType) isScalarType() {}

//

type Scalar interface {
	isScalar()
}

type scalar struct {
	value
}

func (v scalar) isScalar() {}
