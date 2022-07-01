package marshal

import (
	"fmt"
	"math/big"
	"strconv"

	iou "github.com/wrmsr/bane/pkg/util/io"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Value interface {
	Interface() any

	WriteString(w iou.DiscardStringWriter)
	String() string

	isValue()
}

type value struct{}

func (v value) isValue() {}

var (
	_ Value = Null{}
	_ Value = Bool{}
	_ Value = Int{}
	_ Value = Float{}
	_ Value = Number{}
	_ Value = String{}
	_ Value = Bytes{}
	_ Value = Array{}
	_ Value = Object{}
	_ Value = Any{}
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

type Null struct {
	simple
}

var _nullValue = Null{}

func (v Null) Interface() any {
	return nil
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

type String struct {
	simple
	v string
}

func (v String) Interface() any {
	return v.v
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

type Container interface {
	isContainer()
}

type container struct {
	value
}

func (v container) isContainer() {}

//

type Array struct {
	container
	v []Value
}

func (v Array) Interface() any {
	return v
}

//

type Object struct {
	container
	v []bt.Kv[Value, Value]
}

func (v Object) Interface() any {
	return v.v // FIXME: ?
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

func (v Null) WriteString(w iou.DiscardStringWriter) {
	w.WriteString("null")
}

func (v Bool) WriteString(w iou.DiscardStringWriter) {
	if v.v {
		w.WriteString("true")
	} else {
		w.WriteString("false")
	}
}

func (v Int) WriteString(w iou.DiscardStringWriter) {
	if v.u {
		w.WriteString(strconv.FormatUint(uint64(v.v), 10))
	} else {
		w.WriteString(strconv.FormatInt(v.v, 10))
	}
}

func (v Float) WriteString(w iou.DiscardStringWriter) {
	w.WriteString(strconv.FormatFloat(v.v, 'g', -1, 64))
}

func (v Number) WriteString(w iou.DiscardStringWriter) {
	w.WriteString(v.v.String())
}

func (v String) WriteString(w iou.DiscardStringWriter) {
	w.WriteString("\"")
	w.WriteString(v.v)
	w.WriteString("\"")
}

func (v Bytes) WriteString(w iou.DiscardStringWriter) {
	w.WriteString(fmt.Sprintf("%v", v.v))
}

func (v Array) WriteString(w iou.DiscardStringWriter) {
	w.WriteString("[")
	for i, e := range v.v {
		if i > 0 {
			w.WriteString(", ")
		}
		e.WriteString(w)
	}
	w.WriteString("]")
}

func (v Object) WriteString(w iou.DiscardStringWriter) {
	w.WriteString("{")
	for i, kv := range v.v {
		if i > 0 {
			w.WriteString(", ")
		}
		kv.K.WriteString(w)
		w.WriteString(": ")
		kv.V.WriteString(w)
	}
	w.WriteString("}")
}

func (v Any) WriteString(w iou.DiscardStringWriter) {
	w.WriteString(fmt.Sprintf("%v", v.v))
}

func (v Null) String() string   { return iou.InvokeToString(v.WriteString) }
func (v Bool) String() string   { return iou.InvokeToString(v.WriteString) }
func (v Int) String() string    { return iou.InvokeToString(v.WriteString) }
func (v Float) String() string  { return iou.InvokeToString(v.WriteString) }
func (v Number) String() string { return iou.InvokeToString(v.WriteString) }
func (v String) String() string { return iou.InvokeToString(v.WriteString) }
func (v Bytes) String() string  { return iou.InvokeToString(v.WriteString) }
func (v Array) String() string  { return iou.InvokeToString(v.WriteString) }
func (v Object) String() string { return iou.InvokeToString(v.WriteString) }
func (v Any) String() string    { return iou.InvokeToString(v.WriteString) }
