package lisp

import (
	"fmt"

	"github.com/wrmsr/bane/pkg/util/check"
)

//

type Op int8

const (
	OpInvalid Op = iota

	OpApply
	OpAsFalse
	OpAsTrue
	OpDefine
	OpDrop
	OpGoto
	OpIfFalse
	OpLdConst
	OpLdProc
	OpLdVar
	OpReturn
)

var opNames = map[Op]string{
	OpApply:   "apply",
	OpAsFalse: "asfalse",
	OpAsTrue:  "astrue",
	OpDefine:  "define",
	OpDrop:    "drop",
	OpGoto:    "goto",
	OpIfFalse: "iffalse",
	OpLdConst: "ldconst",
	OpLdProc:  "ldproc",
	OpLdVar:   "ldvar",
	OpReturn:  "return",
}

func (o Op) String() string {
	s, ok := opNames[o]
	check.Ok(ok)
	return s
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
