package marshal

import (
	"math/big"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Value interface {
	isValue()
}

type value struct{}

func (v value) isValue() {}

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

//

type Float struct {
	numeric
	v float64
}

//

type Number struct {
	numeric
	v big.Rat
}

//

type Bool struct {
	simple
	v bool
}

//

type Null struct {
	simple
}

//

type Bytes struct {
	simple
	v []byte
}

//

type String struct {
	simple
	v string
}

//

type Any struct {
	simple
	v any
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

//

type Array struct {
	container
	a []Value
}
