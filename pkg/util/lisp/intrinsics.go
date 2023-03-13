package lisp

import (
	"fmt"

	fnu "github.com/wrmsr/bane/pkg/util/funcs"
)

//

type Intrinsic struct {
	name string
	fn   func(args []Value) Value
}

var _ Value = &Intrinsic{}

func (it *Intrinsic) isValue() {}

func (it *Intrinsic) Call(args []Value) Value {
	return it.fn(args)
}

func (it *Intrinsic) String() string {
	return fmt.Sprintf("#[intrinsic-%s]", it.name)
}

func (it *Intrinsic) IsIdentity() bool {
	return true
}

//

func intrinsicNumBinOp(fn func(Value, Value) Value, args []Value) Value {
	switch len(args) {
	case 0:
		return Int(0)
	case 1:
		return AsNumber(args[0])
	case 2:
		return fn(args[0], args[1])
	default:
		panic("nyi")
	}
}

func intrinsicBoolReduce(fn func(Value, Value) bool, args []Value) Value {
	switch len(args) {
	case 0:
		fallthrough
	case 1:
		return Bool(true)
	case 2:
		return Bool(fn(args[0], args[1]))
	default:
		return Bool(reduceConvolution(args, fn))
	}
}

var intrinsics = []*Intrinsic{
	{"+", fnu.Bind1x1x1(intrinsicNumBinOp, NumberAdd)},
	{"-", fnu.Bind1x1x1(intrinsicNumBinOp, NumberSub)},
	{"*", fnu.Bind1x1x1(intrinsicNumBinOp, NumberMul)},
	{"/", fnu.Bind1x1x1(intrinsicNumBinOp, NumberDiv)},

	{"==", fnu.Bind1x1x1(intrinsicBoolReduce, NumberEq)},
	{"!=", fnu.Bind1x1x1(intrinsicBoolReduce, NumberNe)},
	{"<", fnu.Bind1x1x1(intrinsicBoolReduce, NumberLt)},
	{"<=", fnu.Bind1x1x1(intrinsicBoolReduce, NumberLe)},
	{">", fnu.Bind1x1x1(intrinsicBoolReduce, NumberGt)},
	{">=", fnu.Bind1x1x1(intrinsicBoolReduce, NumberGe)},

	{"display", func(args []Value) Value {
		if len(args) != 1 && len(args) != 2 {
			panic("display: proc requires 1 or 2 arguments")
		}

		fmt.Println(AsDisplay(args[0]))
		return nil
	}},

	{"gf", func(args []Value) Value {
		if len(args) != 2 {
			panic("no")
		}
		v := args[0].(Reflect).v
		n := args[1].(String)
		r := v.FieldByName(string(n))
		return Reflect{v: r}
	}},
}

func addIntrinsics(sc *Scope) *Scope {
	for _, i := range intrinsics {
		sc.Set(i.name, i)
	}
	return sc
}
