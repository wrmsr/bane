package lisp

import "fmt"

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

var intrinsics = []*Intrinsic{
	{"+", func(args []Value) Value {
		switch len(args) {
		case 0:
			return Int(0)
		case 1:
			return AsNumber(args[0])
		case 2:
			return NumberAdd(args[0], args[1])
		default:
			panic("nyi")
		}
	}},

	{"<", func(args []Value) Value {
		switch len(args) {
		case 0:
			fallthrough
		case 1:
			return Bool(true)
		case 2:
			return Bool(NumberLt(args[0], args[1]))
		default:
			return Bool(reduceConvolution(args, NumberLt))
		}
	}},

	{">", func(args []Value) Value {
		switch len(args) {
		case 0:
			fallthrough
		case 1:
			return Bool(true)
		case 2:
			return Bool(NumberGt(args[0], args[1]))
		default:
			return Bool(reduceConvolution(args, NumberGt))
		}
	}},
}

func addIntrinsics(sc *Scope) *Scope {
	for _, i := range intrinsics {
		sc.Set(i.name, i)
	}
	return sc
}
