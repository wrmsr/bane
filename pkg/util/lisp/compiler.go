package lisp

import (
	"fmt"
	"strings"
)

//

type Program struct {
	insns []Instr
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

func (co *Compiler) compileValue(p *Program, v Value) {
	if v.IsIdentity() {
		p.add(mkIns(OpLdConst, v))
	} else if _, ok := v.(Atom); ok {
		panic("nyi")
	} else if sl, ok := AsCons(v); ok {
		co.compileList(p, sl)
	} else {
		panic("invalid value type")
	}
}

func (co *Compiler) compileBlock(p *Program, v *Cons) {
	for v != nil {
		co.compileValue(p, v.Car)

		if v.Cdr != nil {
			panic("nyi")
		}

		var ok bool
		if v, ok = AsCons(v.Cdr); !ok {
			panic(fmt.Sprintf("block must be a proper list: %s", v))
		}
	}
}

func (co *Compiler) compileList(p *Program, v *Cons) {
	if v == nil {
		p.add(mkIns(OpLdConst, nil))
		return
	}

	at, ok := v.Car.(Atom)
	if !ok {
		panic("nyi")
	}

	vv, ok := AsCons(v.Cdr)
	if !ok {
		panic(fmt.Sprintf("not applicable: %s", v))
	}

	switch at {
	case "begin":
		co.compileBlock(p, vv)
	default:
		panic("nyi")
	}
}

func (co *Compiler) Compile(src *Cons) (p Program) {
	co.compileList(&p, src)
	p.add(mkIns(OpReturn, nil))
	return p
}
