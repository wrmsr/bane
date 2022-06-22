package lisp

import "fmt"

//

type Scope struct {
	parent *Scope
	defs   map[string]Value
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		parent: parent,
		defs:   make(map[string]Value),
	}
}

//

func Evaluate(s *Scope, p Program) Value {
	pc := 0
	st := make([]Value, 0, 16)

	for pc < len(p.insns) {
		ins := p.insns[pc]
		switch pc++; ins.op {

		case OpLdConst:
			st = append(st, ins.arg.(Value))

		case OpReturn:
			if len(st) != 1 {
				panic("unbalanced stack")
			} else {
				return st[0]
			}

		default:
			panic(fmt.Sprintf("invalid instruction: %s", ins))
		}
	}

	panic(fmt.Sprintf("program not returned properly:\n%s", p))
}
