package box

//

type StringType struct {
	scalarType
}

var _stringType = StringType{scalarType{type_{t: typeOf[string]()}}}

//

type String struct {
	scalar
}

func (v String) Type() Type { return _stringType }
