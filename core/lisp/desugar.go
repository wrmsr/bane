package lisp

import "fmt"

//

func desugarDo(v *Cons) *Cons {
	var decl *Cons
	var cond *Cons
	var body *Cons
	var defs []Value
	var init []Value
	var step []Value

	p := v
	ok := false

	if p == nil {
		panic("malformed do construct: " + v.String())
	}
	if decl, ok = p.Car.(*Cons); !ok {
		panic("malformed do construct: " + v.String())
	}
	if p, ok = p.Cdr.(*Cons); !ok {
		panic("malformed do construct: " + v.String())
	}
	if cond, ok = p.Car.(*Cons); !ok {
		panic("malformed do construct: " + v.String())
	}
	if body, ok = p.Cdr.(*Cons); !ok {
		panic("malformed do construct: " + v.String())
	}

	for p = decl; p != nil; {
		var s Atom
		var q *Cons
		var i Value
		var r Value

		if q, ok = p.Car.(*Cons); !ok {
			panic("malformed do construct: " + decl.String())
		}
		if p, ok = AsCons(p.Cdr); !ok {
			panic("malformed do construct: " + decl.String())
		}
		if s, ok = q.Car.(Atom); !ok {
			panic("malformed do construct: " + decl.String())
		}
		if q, ok = q.Cdr.(*Cons); !ok {
			panic("malformed do construct: " + decl.String())
		}

		if i, r = q.Car, s; q.Cdr != nil {
			if q, ok = q.Cdr.(*Cons); !ok {
				panic("malformed do construct: " + decl.String())
			}
			if r, ok = q.Car.(*Cons); !ok {
				panic("malformed do construct: " + decl.String())
			}
			if q.Cdr != nil {
				panic("malformed do construct: " + decl.String())
			}
		}

		defs = append(defs, s)
		init = append(init, i)
		step = append(step, r)
	}

	if p, ok = AsCons(cond.Cdr); !ok {
		panic("malformed do construct: " + cond.String())
	}
	if p != nil && p.Cdr != nil {
		panic("malformed do construct: " + cond.String())
	}

	if p == nil {
		return rebuildDo(defs, init, step, cond.Car, nil, body)
	} else {
		return rebuildDo(defs, init, step, cond.Car, p.Car, body)
	}
}

func rebuildDo(defs []Value, init []Value, step []Value, cond Value, retv Value, body *Cons) *Cons {
	var pb, qb *Cons
	var pd, qd *Cons
	var pi, qi *Cons
	var ps, qs *Cons

	ok := true
	name := fmt.Sprintf("#[desugar-do-%d]", nextId())

	for _, v := range defs {
		AppendValue(&pd, &qd, v)
	}
	for _, v := range init {
		AppendValue(&pi, &qi, v)
	}
	for _, v := range step {
		AppendValue(&ps, &qs, v)
	}

	for p := body; ok && p != nil; p, ok = AsCons(p.Cdr) {
		AppendValue(&pb, &qb, p.Car)
	}

	if !ok || body == nil {
		panic("loop body must be a proper list: " + body.String())
	}

	if retv == nil {
		retv = MakeList(Atom("quote"), nil)
	}

	AppendValue(&pb, &qb, AsPair(
		Atom(name),
		ps,
	))

	loop := MakeList(
		Atom("if"),
		cond,
		retv,
		AsPair(Atom("begin"), pb),
	)

	return AsPair(Atom("letrec"), AsPair(
		MakeList(MakeList(Atom(name), MakeList(Atom("lambda"), pd, loop))),
		MakeList(AsPair(Atom(name), pi)),
	))
}

//

type letKind int8

const (
	let letKind = iota
	letRec
	letStar
)

func desugarLet(v *Cons, kind letKind) *Cons {
	var decl *Cons
	var body *Cons
	var defs []Value
	var init []Value

	p := v
	n := 0
	ok := false

	if p == nil {
		panic("malformed let construct: " + v.String())
	}
	if decl, ok = AsCons(p.Car); !ok {
		panic("malformed let construct: " + v.String())
	}
	if body, ok = p.Cdr.(*Cons); !ok {
		panic("malformed let construct: " + v.String())
	}

	for p = decl; p != nil; n++ {
		var s Atom
		var q *Cons

		if q, ok = p.Car.(*Cons); !ok {
			panic("malformed let construct: " + decl.String())
		}
		if p, ok = AsCons(p.Cdr); !ok {
			panic("malformed let construct: " + decl.String())
		}
		if s, ok = q.Car.(Atom); !ok {
			panic("malformed let construct: " + decl.String())
		}
		if q, ok = q.Cdr.(*Cons); !ok {
			panic("malformed let construct: " + decl.String())
		}
		if q.Cdr != nil {
			panic("malformed let construct: " + decl.String())
		}

		defs = append(defs, s)
		init = append(init, q.Car)
	}

	switch kind {
	case let:
		break
	case letRec:
		return rebuildLetRec(defs, init, body)
	case letStar:
		n = 1
	default:
		panic("fatal: invalid let kind")
	}

	for i := len(defs) - n; i >= 0; i-- {
		arg := MakeList(defs[i : i+n]...)
		ref := AsPair(Atom("lambda"), AsPair(arg, body))
		body = MakeList(MakeList(append([]Value{ref}, init[i:i+n]...)...))
	}

	return body.Car.(*Cons)
}

func rebuildLetRec(defs []Value, init []Value, body *Cons) *Cons {
	var pd, qd *Cons
	var pi, qi *Cons

	for _, v := range defs {
		AppendValue(&pd, &qd, MakeList(v, MakeList(Atom("quote"), nil)))
	}

	for i, v := range init {
		AppendValue(&pi, &qi, MakeList(Atom("set!"), defs[i], v))
	}

	qi.Cdr = body
	return AsPair(Atom("let"), AsPair(pd, pi))
}
