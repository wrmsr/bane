package tower

//

type StringType struct {
	BaseScalarType
}

var _stringType = StringType{BaseScalarType{BaseType{ty: typeOf[string]()}}}

//

type String struct {
	BaseScalar
}
