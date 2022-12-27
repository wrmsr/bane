//go:generate go run github.com/wrmsr/bane/exp/util/wasm/instrs/dev/gen

package instrs

import wt "github.com/wrmsr/bane/exp/util/wasm/types"

//

type Instr int16

type Def struct {
	i Instr

	class Class
	name  string
	name2 string
	opPfx uint8
	op    uint8

	t wt.Prim
	a wt.Prim
	b wt.Prim
	c wt.Prim

	ma Access
	mz int
}

func (d Def) I() Instr { return d.i }

func (d Def) Class() Class  { return d.class }
func (d Def) Name() string  { return d.name }
func (d Def) Name2() string { return d.name2 }
func (d Def) OpPfx() uint8  { return d.opPfx }
func (d Def) Op() uint8     { return d.op }

func (d Def) T() wt.Prim { return d.t }
func (d Def) A() wt.Prim { return d.a }
func (d Def) B() wt.Prim { return d.b }
func (d Def) C() wt.Prim { return d.c }

func (d Def) Ma() Access { return d.ma }
func (d Def) Mz() int    { return d.mz }

//

//var defs []Def
//var opMap [256]*[256]Instr
//var nameMap = map[string]Instr
//var name2Map = map[string]Instr

//

func Get(i Instr) Def {
	return defs[i]
}

func ByOp(opPfx, op int8) Def {
	return defs[opMap[opPfx][op]]
}

func ByName(n string) Def {
	return defs[nameMap[n]]
}

func ByName2(n2 string) Def {
	return defs[name2Map[n2]]
}
