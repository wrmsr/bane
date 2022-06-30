package tower

//

type IntegralType interface {
	NumericType
	isIntegralType()
}

type integralType struct {
	numericType
}

func (t integralType) isIntegralType() {}

//

type Integral interface {
	Numeric
	isIntegral()
}

type integral struct {
	numeric
}

func (t integral) isIntegral() {}
