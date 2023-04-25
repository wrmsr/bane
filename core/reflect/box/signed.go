package box

///

type SignedType interface {
	IntegralType
	isSignedType()
}

type signedType struct {
	integralType
}

func (t signedType) isSignedType() {}

//

type IntType struct{ signedType }
type Int8Type struct{ signedType }
type Int16Type struct{ signedType }
type Int32Type struct{ signedType }
type Int64Type struct{ signedType }

var (
	_intType   = IntType{signedType{integralType{numericType{scalarType{type_{t: typeOf[int]()}}}}}}
	_int8Type  = Int8Type{signedType{integralType{numericType{scalarType{type_{t: typeOf[int8]()}}}}}}
	_int16Type = Int16Type{signedType{integralType{numericType{scalarType{type_{t: typeOf[int16]()}}}}}}
	_int32Type = Int32Type{signedType{integralType{numericType{scalarType{type_{t: typeOf[int32]()}}}}}}
	_int64Type = Int64Type{signedType{integralType{numericType{scalarType{type_{t: typeOf[int64]()}}}}}}
)

///

type Signed interface {
	Integral
	isSigned()
}

type signed struct {
	integral
}

func (v signed) isSigned() {}

//

type Int struct{ signed }
type Int8 struct{ signed }
type Int16 struct{ signed }
type Int32 struct{ signed }
type Int64 struct{ signed }

func (v Int) Type() Type   { return _intType }
func (v Int8) Type() Type  { return _int8Type }
func (v Int16) Type() Type { return _int16Type }
func (v Int32) Type() Type { return _int32Type }
func (v Int64) Type() Type { return _int64Type }
