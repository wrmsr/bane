package lisp

import (
	"fmt"
	"strings"
)

//

type Program struct {
	insns []Instr
}

func (p *Program) pc() int {
	return len(p.insns)
}

func (p *Program) pin(i int) {
	p.insns[i].arg = Int(p.pc())
}

func (p *Program) add(ins Instr) {
	p.insns = append(p.insns, ins)
}

func (p Program) String() string {
	var sb strings.Builder
	for _, i := range p.insns {
		sb.WriteString(i.String())
		sb.WriteRune('\n')
	}
	return sb.String()
}

//

type Compiler struct{}

func NewCompiler() *Compiler {
	return &Compiler{}
}

func (co Compiler) compileArgs(p *Program, v *Cons, n int) int {
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
		panic(fmt.Errorf("expect %d arguments, got %d.", n, na))
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

func (co Compiler) compileCondition(p *Program, v *Cons) {
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
	case "if":
		co.compileCondition(p, vv)
	case "define":
		co.compileDefine(p, vv)
	default:
		na := co.compileArgs(p, v, -1)
		p.add(mkIns(OpApply, AsValue(na)))
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
