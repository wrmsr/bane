package log

import "fmt"

//

type Arg interface {
	ArgName() string
	ArgValue() any

	isArg()
}

//

type StackOffset struct{ Num int }

func (a StackOffset) ArgName() string { return "" }
func (a StackOffset) ArgValue() any   { return a.Num }

func (a StackOffset) isArg() {}

//

type ErrorArg struct{ Err error }

func (a ErrorArg) ArgName() string { return "$error" }
func (a ErrorArg) ArgValue() any   { return a.Err }

func (a ErrorArg) isArg() {}

//

type PanicArg struct{ Val any }

func (a PanicArg) ArgName() string { return "$panic" }
func (a PanicArg) ArgValue() any   { return a.Val }

func (a PanicArg) isArg() {}

//

type NamedArg struct {
	Name  string
	Value any
}

func (a NamedArg) ArgName() string { return a.Name }
func (a NamedArg) ArgValue() any   { return a.Value }

func (a NamedArg) isArg() {}

//

func A(name string, value any) NamedArg {
	if name == "" || name[0] == '$' {
		panic(fmt.Errorf("illegal arg name: %s", name))
	}
	return NamedArg{Name: name, Value: value}
}

func E(err error) ErrorArg {
	return ErrorArg{Err: err}
}

func P(r any) PanicArg {
	return PanicArg{Val: r}
}
