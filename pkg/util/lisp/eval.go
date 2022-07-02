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

func stackPop(st *[]Value) (v Value) {
	i := len(*st) - 1
	if i < 0 {
		panic("stack underflow")
	}

	v, *st = (*st)[i], (*st)[:i]
	return
}

func isTrue(v Value) bool {
	switch vv := v.(type) {
	case nil:
		return false
	case Int:
		return vv != 0
	case Bool:
		return bool(vv)
	case Char:
		return vv != 0
	case Float:
		return vv != 0.0
	case String:
		return vv != ""
	case Complex:
		return vv != 0
	default:
		return true
	}
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

		/* branch if the stack top is #f */
		case OpIfFalse:
			if !isTrue(stackPop(&st)) {
				if pc = int(ins.arg.(Int)); pc < 0 || pc >= len(p.insns) {
					panic(fmt.Sprintf("branch out of scope: %s", ins.arg))
				}
			}

		/* unconditional jump */
		case OpGoto:
			if pc = int(ins.arg.(Int)); pc < 0 || pc >= len(p.insns) {
				panic(fmt.Sprintf("branch out of scope: %s", ins.arg))
			}

		case OpApply:
			nb := int(ins.arg.(Int))
			vv := stackRemove(&st, nb)

			fn, ok := vv[0].(Callable)
			if !ok {
				panic(fmt.Sprintf("object is not applicable: %s", AsString(vv[0])))
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
