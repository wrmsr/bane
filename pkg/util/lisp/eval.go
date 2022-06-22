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

func (sc *Scope) Get(key string) (v Value, ok bool) {
	for p := sc; !ok && p != nil; p = p.parent {
		v, ok = p.defs[key]
	}
	return
}

func (sc *Scope) Set(key string, val Value) {
	sc.defs[key] = val
}

//

func stackRemove(st *[]Value, nb int) (v []Value) {
	if nb == 0 {
		panic("taking nothing")
	}

	i := len(*st) - nb
	if i < 0 {
		panic("stack underflow")
	}

	v, *st = (*st)[i:], (*st)[:i]
	return
}

func Evaluate(s *Scope, p Program) Value {
	pc := 0
	st := make([]Value, 0, 16)

	for pc < len(p.insns) {
		ins := p.insns[pc]
		switch pc++; ins.op {

		case OpLdConst:
			st = append(st, ins.arg.(Value))

		case OpLdVar:
			vv, ok := s.Get(string(ins.arg.(Atom)))
			if !ok {
				panic(fmt.Sprintf("undefined reference: %s", ins.arg))
			}
			st = append(st, vv)

		case OpApply:
			nb := int(ins.arg.(Int))
			vv := stackRemove(&st, nb)

			fn, ok := vv[0].(Callable)
			if !ok {
				panic("eval: object is not appliable: " + AsString(vv[0]))
			}
			st = append(st, fn.Call(vv[1:]))

		case OpReturn:
			if len(st) != 1 {
				panic("unbalanced stack")
			}
			return st[0]

		default:
			panic(fmt.Sprintf("invalid instruction: %s", ins))
		}
	}

	panic(fmt.Sprintf("program not returned properly:\n%s", p))
}
