package lisp

import "strings"

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
