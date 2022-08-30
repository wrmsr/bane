package lisp

import (
	"fmt"
	"strings"
)

type Proc struct {
	Name string
	Code Program
	Args []string
}

var _ Value = &Proc{}

func (p *Proc) isValue() {}

func (p *Proc) String() string {
	if len(p.Args) == 0 {
		return fmt.Sprintf("#[proc (%s)]", p.Name)
	} else {
		return fmt.Sprintf("#[proc (%s %s)]", p.Name, strings.Join(p.Args, " "))
	}
}

func (p *Proc) IsIdentity() bool {
	return true
}

func (p *Proc) LoadWithScope(scope *Scope) LoadedProc {
	return LoadedProc{
		Proc:  p,
		Scope: scope,
	}
}

type LoadedProc struct {
	*Proc
	*Scope
}

func (p LoadedProc) Call(args []Value) Value {
	return Evaluate(p.Scope.Derive(p.Proc, args), p.Code)
}
