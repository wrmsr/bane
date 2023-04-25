package box

///

type UnsignedType interface {
	IntegralType
	isUnsignedType()
}

type unsignedType struct {
	integralType
}

func (t unsignedType) isUnsignedType() {}

//

type UintType struct{ unsignedType }
type Uint8Type struct{ unsignedType }
type Uint16Type struct{ unsignedType }
type Uint32Type struct{ unsignedType }
type Uint64Type struct{ unsignedType }
type UintptrType struct{ unsignedType }

var (
	_uintType    = UintType{unsignedType{integralType{numericType{scalarType{type_{t: typeOf[uint]()}}}}}}
	_uint8Type   = Uint8Type{unsignedType{integralType{numericType{scalarType{type_{t: typeOf[uint8]()}}}}}}
	_uint16Type  = Uint16Type{unsignedType{integralType{numericType{scalarType{type_{t: typeOf[uint16]()}}}}}}
	_uint32Type  = Uint32Type{unsignedType{integralType{numericType{scalarType{type_{t: typeOf[uint32]()}}}}}}
	_uint64Type  = Uint64Type{unsignedType{integralType{numericType{scalarType{type_{t: typeOf[uint64]()}}}}}}
	_uintptrType = UintptrType{unsignedType{integralType{numericType{scalarType{type_{t: typeOf[uintptr]()}}}}}}
)

///

type Unsigned interface {
	Integral
	isUnsigned()
}

type unsigned struct {
	integral
}

func (v unsigned) isUnsigned() {}

//

type Uint struct{ unsigned }
type Uint8 struct{ unsigned }
type Uint16 struct{ unsigned }
type Uint32 struct{ unsigned }
type Uint64 struct{ unsigned }
type Uintptr struct{ unsigned }

func (v Uint) Type() Type    { return _uintType }
func (v Uint8) Type() Type   { return _uint8Type }
func (v Uint16) Type() Type  { return _uint16Type }
func (v Uint32) Type() Type  { return _uint32Type }
func (v Uint64) Type() Type  { return _uint64Type }
func (v Uintptr) Type() Type { return _uintptrType }
