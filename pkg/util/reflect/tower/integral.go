package tower

//

type IntegralType interface {
	isIntegralType()
}

type integralType struct {
	numericType
}

func (t integralType) isIntegralType() {}

//

type Integral interface {
	isIntegral()
}

type integral struct {
	numeric
}

func (t integral) isIntegral() {}
