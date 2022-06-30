package marshal

import (
	"math/big"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Value interface {
	Interface() any

	isValue()
}

type value struct{}

func (v value) isValue() {}

var (
	_ Value = Int{}
	_ Value = Float{}
	_ Value = Number{}
	_ Value = Bool{}
	_ Value = Null{}
	_ Value = Bytes{}
	_ Value = String{}
	_ Value = Any{}
	_ Value = Object{}
	_ Value = Array{}
)

//

type Simple interface {
	isSimple()
}

type simple struct {
	value
}

func (v simple) isSimple() {}

//

type Numeric interface {
	isNumeric()
}

type numeric struct {
	simple
}

func (v numeric) isNumeric() {}

//

type Int struct {
	numeric
	v int64
	u bool
}

func (v Int) Interface() any {
	if v.u {
		return uint64(v.v)
	} else {
		return v.v
	}
}

//

type Float struct {
	numeric
	v float64
}

func (v Float) Interface() any {
	return v.v
}

//

type Number struct {
	numeric
	v big.Rat
}

func (v Number) Interface() any {
	return v.v
}

//

type Bool struct {
	simple
	v bool
}

var (
	_trueValue  = Bool{v: true}
	_falseValue = Bool{v: false}
)

func (v Bool) Interface() any {
	return v.v
}

//

type Null struct {
	simple
}

var _nullValue = Null{}

func (v Null) Interface() any {
	return nil
}

//

type Bytes struct {
	simple
	v []byte
}

func (v Bytes) Interface() any {
	return v.v
}

//

type String struct {
	simple
	v string
}

func (v String) Interface() any {
	return v.v
}

//

type Any struct {
	simple
	v any
}

func (v Any) Interface() any {
	return v.v
}

//

type Container interface {
	isContainer()
}

type container struct {
	value
}

func (v container) isContainer() {}

//

type Object struct {
	container
	kvs []bt.Kv[Value, Value]
}

func (v Object) Interface() any {
	return v.kvs // FIXME: ?
}

//

type Array struct {
	container
	v []Value
}

func (v Array) Interface() any {
	return v
}
