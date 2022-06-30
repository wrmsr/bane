package tower

//

type IntegralType interface {
	isIntegralType()
}

type BaseIntegralType struct {
	BaseNumericType
}

func (t BaseIntegralType) isIntegralType() {}

//

type Integral interface {
	isIntegral()
}

type BaseIntegral struct {
	BaseNumeric
}

func (t BaseIntegral) isIntegral() {}
