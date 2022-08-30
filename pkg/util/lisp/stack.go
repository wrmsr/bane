package lisp

type stack []Value

func (st *stack) remove(nb int) (v []Value) {
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

func (st *stack) pop() (v Value) {
	i := len(*st) - 1
	if i < 0 {
		panic("stack underflow")
	}

	v, *st = (*st)[i], (*st)[:i]
	return
}

func (st stack) top() Value {
	if nb := len(st); nb == 0 {
		panic("stack underflow")
	} else {
		return st[nb-1]
	}
}

func (st stack) sub(vv Value) {
	st[len(st)-1] = vv
}
