package lisp

import (
	"fmt"
	"sync/atomic"
)

//

var _nextId int64

func nextId() int64 {
	return atomic.AddInt64(&_nextId, 1)
}

//

type Compiler struct{}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (co *Compiler) compileArgs(p *Program, v *Cons, n int) int {
	na := 0
	for s := v; s != nil; {
		vv, ok := AsCons(s.Cdr)
		if !ok {
			panic(fmt.Errorf("list is not applicable: %s", v))
		}

		co.compileValue(p, s.Car)
		s = vv
		na++
	}

	if !(n < 0 || n == na) {
		panic(fmt.Errorf("expect %d arguments, got %d", n, na))
	}
	return na
}

func (co *Compiler) compileBlock(p *Program, v *Cons) {
	for v != nil {
		co.compileValue(p, v.Car)

		if v.Cdr != nil {
			p.add(mkIns(OpDrop, nil))
		}

		var ok bool
		if v, ok = AsCons(v.Cdr); !ok {
			panic(fmt.Errorf("block must be a proper list: %s", v))
		}
	}
}

func (co *Compiler) compileCondition(p *Program, v *Cons) {
	var ok bool
	var al *Cons
	var pp *Cons

	if pp, ok = v.Cdr.(*Cons); !ok {
		panic(fmt.Errorf("malformed if construct: %s", v))
	}
	if al, ok = AsCons(pp.Cdr); !ok {
		panic(fmt.Errorf("malformed if construct: %s", v))
	}
	if al != nil && al.Cdr != nil {
		panic(fmt.Errorf("malformed if construct: %s", v))
	}

	co.compileValue(p, v.Car)
	i := p.pc()
	p.add(mkIns(OpIfFalse, nil))

	co.compileValue(p, pp.Car)
	j := p.pc()
	p.add(mkIns(OpGoto, nil))
	p.pin(i)

	if al == nil {
		p.add(mkIns(OpLdConst, nil))
		p.pin(j)
		return
	}

	co.compileValue(p, al.Car)
	p.pin(j)
}

func (co *Compiler) compileDefine(p *Program, v *Cons) {
	var name Atom
	var decl *Cons

	pp := v
	ok := false

	if pp, ok = v.Cdr.(*Cons); !ok {
		panic("malformed define construct: " + v.String())
	}
	if name, ok = v.Car.(Atom); ok && pp.Cdr != nil {
		panic("malformed define construct: " + v.String())
	}

	if ok {
		co.compileValue(p, pp.Car)
		p.add(mkIns(OpDefine, name))
		return
	}

	if decl, ok = v.Car.(*Cons); !ok {
		panic("malformed define construct: " + v.String())
	}
	if name, ok = decl.Car.(Atom); !ok {
		panic("malformed define construct: " + v.String())
	}
	if decl, ok = AsCons(decl.Cdr); !ok {
		panic("malformed define construct: " + v.String())
	}

	co.compileLambda(p, AsPair(decl, pp), name.String())
	p.add(mkIns(OpDefine, name))
}

func (co *Compiler) compileLambda(p *Program, v *Cons, name string) {
	var atom Atom
	var decl *Cons
	var proc *Cons
	var args []string

	pp := v
	ok := true

	if decl, ok = pp.Car.(*Cons); !ok {
		panic("malformed proc construct: " + v.String())
	}
	if proc, ok = AsCons(pp.Cdr); !ok {
		panic("malformed proc construct: " + v.String())
	}

	for q := decl; ok && q != nil; q, ok = AsCons(q.Cdr) {
		if atom, ok = q.Car.(Atom); ok {
			args = append(args, string(atom))
		} else {
			panic("malformed proc construct: " + v.String())
		}
	}

	if !ok {
		panic("malformed proc construct: " + v.String())
	}

	p.add(mkIns(OpLdProc, &Proc{
		Args: args,
		Name: name,
		Code: co.Compile(AsPair(Atom("begin"), proc)),
	}))
}

func (co *Compiler) compileList(p *Program, v *Cons) {
	if v == nil {
		p.add(mkIns(OpLdConst, nil))
		return
	}

	at, ok := v.Car.(Atom)
	if !ok {
		na := co.compileArgs(p, v, -1)
		p.add(mkIns(OpApply, AsValue(na)))
		return
	}

	vv, ok := AsCons(v.Cdr)
	if !ok {
		panic(fmt.Errorf("not applicable: %s", v))
	}

	switch at {

	case "begin":
		co.compileBlock(p, vv)
	case "define":
		co.compileDefine(p, vv)
	case "if":
		co.compileCondition(p, vv)
	case "quote":
		co.compileQuote(p, vv)
	case "set!":
		co.compileSet(p, vv)

	case "and":
		co.compileShortCircuit(p, vv, conj)
	case "or":
		co.compileShortCircuit(p, vv, disj)

	case "lambda":
		co.compileLambda(p, vv, fmt.Sprintf("#[lambda-%d]", nextId()))

	case "car":
		co.compileArgs(p, vv, 1)
		p.add(mkIns(OpCar, nil))
	case "cdr":
		co.compileArgs(p, vv, 1)
		p.add(mkIns(OpCdr, nil))
	case "cons":
		co.compileArgs(p, vv, 2)
		p.add(mkIns(OpCons, nil))

	case "do":
		co.compileList(p, desugarDo(vv))

	case "let":
		co.compileList(p, desugarLet(vv, let))
	case "let*":
		co.compileList(p, desugarLet(vv, letStar))
	case "letrec":
		co.compileList(p, desugarLet(vv, letRec))

	default:
		na := co.compileArgs(p, v, -1)
		p.add(mkIns(OpApply, AsValue(na)))
	}
}

func (co *Compiler) compileQuote(p *Program, v *Cons) {
	if v == nil || v.Cdr != nil {
		panic("`quote` takes exact 1 argument: " + v.String())
	}
	p.add(mkIns(OpLdConst, v.Car))
}

func (co *Compiler) compileSet(p *Program, v *Cons) {
	var ok bool
	var sn Atom
	var vv *Cons

	if sn, ok = v.Car.(Atom); !ok {
		panic("malformed set! construct: " + v.String())
	}
	if vv, ok = v.Cdr.(*Cons); !ok {
		panic("malformed set! construct: " + v.String())
	}
	if vv.Cdr != nil {
		panic("malformed set! construct: " + v.String())
	}

	co.compileValue(p, vv.Car)
	p.add(mkIns(OpSet, sn))
}

type relKind int8

const (
	conj relKind = iota
	disj
)

func (co *Compiler) compileShortCircuit(p *Program, v *Cons, kind relKind) {
	var ok bool
	var pp *Cons
	var br []int

	if v == nil {
		panic("empty condition")
	}

	co.compileValue(p, v.Car)
	pp, ok = AsCons(v.Cdr)

	for ok && pp != nil {
		car := pp.Car
		cdr := pp.Cdr

		switch br = append(br, p.pc()); kind {
		case conj:
			p.add(mkIns(OpAsTrue, nil))
		case disj:
			p.add(mkIns(OpAsFalse, nil))
		default:
			panic("invalid relationship kind")
		}

		pp, ok = AsCons(cdr)
		co.compileValue(p, car)
	}

	if !ok {
		panic("malformed short-circuit construct: " + v.String())
	}

	for _, pc := range br {
		p.pin(pc)
	}
}

func (co *Compiler) compileValue(p *Program, v Value) {
	if v.IsIdentity() {
		p.add(mkIns(OpLdConst, v))
		return
	}

	if at, ok := v.(Atom); ok {
		p.add(mkIns(OpLdVar, at))
		return
	}

	if sl, ok := AsCons(v); ok {
		co.compileList(p, sl)
		return
	}

	panic("invalid value type")
}

func (co *Compiler) Compile(src *Cons) (p Program) {
	co.compileList(&p, src)
	p.add(mkIns(OpReturn, nil))
	return p
}
