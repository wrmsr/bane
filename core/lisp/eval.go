package lisp

import "fmt"

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
	st := make(stack, 0, 16)

	for pc < len(p.insns) {
		ins := p.insns[pc]
		switch pc++; ins.op {

		case OpApply:
			nb := int(ins.arg.(Int))
			vv := st.remove(nb)

			fn, ok := vv[0].(Callable)
			if !ok {
				panic(fmt.Errorf("object is not applicable: %s", AsString(vv[0])))
			}
			st = append(st, fn.Call(vv[1:]))

		case OpAsFalse:
			if !isTrue(st.top()) {
				st.pop()
			} else if pc = int(AsNumber(ins.arg).AsInt()); pc < 0 || pc >= len(p.insns) {
				panic("branch out of scope: " + ins.String())
			}

		case OpAsTrue:
			if isTrue(st.top()) {
				st.pop()
			} else if pc = int(AsNumber(ins.arg).AsInt()); pc < 0 || pc >= len(p.insns) {
				panic("branch out of scope: " + ins.String())
			}

		case OpCar:
			if r, ok := st.top().(*Cons); ok {
				st.sub(r.Car)
			} else {
				panic("invalid argument type for car: " + AsString(st.top()))
			}

		case OpCdr:
			if r, ok := st.top().(*Cons); ok {
				st.sub(r.Cdr)
			} else {
				panic("invalid argument type for cdr: " + AsString(st.top()))
			}

		case OpCons:
			cdr := st.pop()
			car := st.top()
			st.sub(AsPair(car, cdr))

		case OpDefine:
			s.Set(ins.arg.String(), st.top())

		case OpDrop:
			st.pop()

		case OpGoto:
			if pc = int(ins.arg.(Int)); pc < 0 || pc >= len(p.insns) {
				panic(fmt.Errorf("branch out of scope: %s", ins.arg))
			}

		case OpIfFalse:
			if !isTrue(st.pop()) {
				if pc = int(ins.arg.(Int)); pc < 0 || pc >= len(p.insns) {
					panic(fmt.Errorf("branch out of scope: %s", ins.arg))
				}
			}

		case OpLdConst:
			st = append(st, ins.arg)

		case OpLdProc:
			st = append(st, ins.arg.(*Proc).LoadWithScope(s))

		case OpLdVar:
			vv, ok := s.Get(string(ins.arg.(Atom)))
			if !ok {
				panic(fmt.Errorf("undefined reference: %s", ins.arg))
			}
			st = append(st, vv)

		case OpReturn:
			if len(st) != 1 {
				panic("unbalanced stack")
			}
			return st[0]

		case OpSet:
			s.resolve(string(ins.arg.(Atom))).update(st.top())

		default:
			panic(fmt.Errorf("invalid instruction: %s", ins))
		}
	}

	panic(fmt.Errorf("program not returned properly:\n%s", p))
}
