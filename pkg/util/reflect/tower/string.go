package tower

//

type StringType struct {
	scalarType
}

var _stringType = StringType{scalarType{type_{ty: typeOf[string]()}}}

//

type String struct {
	scalar
}
