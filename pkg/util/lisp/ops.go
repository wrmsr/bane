package lisp

import "fmt"

//

type Op int8

const (
	OpInvalid Op = iota
	OpLdConst
	OpReturn
)

func (o Op) String() string {
	switch o {
	case OpLdConst:
		return "ldconst"
	case OpReturn:
		return "return"
	}
	panic("unreachable")
}

//

type Instr struct {
	op  Op
	arg Value
}

func mkIns(op Op, arg Value) Instr {
	return Instr{op: op, arg: arg}
}

func (i Instr) String() string {
	return fmt.Sprintf("%-8s%v", i.op, i.arg)
}
