/*
Copyright (c) 2009 The Go Authors. All rights reserved.

Redistribution and use in source and binary forms, with or without modification, are permitted provided that the
following conditions are met:

  - Redistributions of source code must retain the above copyright notice, this list of conditions and the following
    disclaimer.
  - Redistributions in binary form must reproduce the above copyright notice, this list of conditions and the following
    disclaimer in the documentation and/or other materials provided with the distribution.
  - Neither the name of Google Inc. nor the names of its contributors may be used to endorse or promote products derived
    from this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES,
INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF
THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/
package ssa

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"math"
	"strings"
	"testing"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"

	"github.com/wrmsr/bane/x/go/analysis/runner"

	"github.com/wrmsr/bane/core/check"
	gg "github.com/wrmsr/bane/core/go/gen"
)

//

type SSA struct {
	Pkg      *ssa.Package
	SrcFuncs []*ssa.Function
}

func run(
	fset *token.FileSet,
	files []*ast.File,
	pkg *types.Package,
	ti *types.Info,
) (*SSA, error) {
	var mode ssa.BuilderMode

	mode |= ssa.PrintPackages | ssa.PrintFunctions
	mode |= ssa.InstantiateGenerics

	prog := ssa.NewProgram(fset, mode)

	created := make(map[*types.Package]bool)
	var createAll func(pkgs []*types.Package)
	createAll = func(pkgs []*types.Package) {
		for _, p := range pkgs {
			if !created[p] {
				created[p] = true
				prog.CreatePackage(p, nil, nil, true)
				createAll(p.Imports())
			}
		}
	}
	createAll(pkg.Imports())

	ssapkg := prog.CreatePackage(pkg, files, ti, false)
	ssapkg.Build()

	var funcs []*ssa.Function
	for _, f := range files {
		for _, decl := range f.Decls {
			if fdecl, ok := decl.(*ast.FuncDecl); ok {
				if fdecl.Name.Name == "_" {
					continue
				}

				fn := ti.Defs[fdecl.Name].(*types.Func)
				if fn == nil {
					panic(fn)
				}

				f := ssapkg.Prog.FuncValue(fn)
				if f == nil {
					panic(fn)
				}

				var addAnons func(f *ssa.Function)
				addAnons = func(f *ssa.Function) {
					funcs = append(funcs, f)
					for _, anon := range f.AnonFuncs {
						addAnons(anon)
					}
				}
				addAnons(f)
			}
		}
	}

	return &SSA{Pkg: ssapkg, SrcFuncs: funcs}, nil
}

//

type ftBlockDstPhi struct {
	phi *ssa.Phi
	v   ssa.Value
}

type ftBlockDst struct {
	b    *ssa.BasicBlock
	phis []ftBlockDstPhi
}

type ftBlock struct {
	b    *ssa.BasicBlock
	dsts map[*ssa.BasicBlock]*ftBlockDst
}

type FuncTransformer struct {
	fn *ssa.Function

	vns map[ssa.Value]int

	bs map[*ssa.BasicBlock]*ftBlock
}

func (ft *FuncTransformer) instrIdent(instr ssa.Instruction) gg.Ident {
	n, ok := ft.vns[instr.(ssa.Value)]
	check.Ok(ok)
	return gg.IdentOf(fmt.Sprintf("_ssa_%d", n))
}

func (ft *FuncTransformer) DoValue(value ssa.Value) gg.Expr {
	if instr, ok := value.(ssa.Instruction); ok {
		return ft.instrIdent(instr)
	}

	switch value := value.(type) {

	case *ssa.Parameter:
		return gg.IdentOf(value.Name())

	case *ssa.Const:
		return gg.LitOf(value.Value.String())

	}
	panic(value)
}

func (ft *FuncTransformer) blockIdent(block *ssa.BasicBlock) gg.Ident {
	return gg.Ident{Name: fmt.Sprintf("_ssa_block_%d", block.Index)}
}

func (ft *FuncTransformer) doPhiSrc(src, dst *ssa.BasicBlock, rest ...gg.Stmt) []gg.Stmt {
	ftsrc := ft.bs[src]
	ftdst := ftsrc.dsts[dst]
	if ftdst == nil || len(ftdst.phis) < 1 {
		return rest
	}

	ret := make([]gg.Stmt, len(ftdst.phis)+len(rest))

	for i, p := range ftdst.phis {
		ret[i] = gg.AssignOf(
			ft.DoValue(p.phi),
			ft.DoValue(p.v),
		)
	}

	copy(ret[len(ftdst.phis):], rest)
	return ret
}

func (ft *FuncTransformer) DoInstr(instr ssa.Instruction) []gg.Stmt {
	switch instr := instr.(type) {

	// register

	case *ssa.BinOp:
		x := ft.DoValue(instr.X)
		y := ft.DoValue(instr.Y)
		return []gg.Stmt{gg.AssignOf(
			ft.DoValue(instr),
			gg.BinaryOf(gg.BinaryOpFromToken(instr.Op), x, y),
		)}

	case *ssa.Call:
		cc := instr.Common()
		var gfn gg.Expr
		if !cc.IsInvoke() {
			gfn = gg.Raw{Raw: ft.rawRelValue(cc.Value)}
		} else {
			panic("nyi")
		}
		args := make([]gg.Expr, len(cc.Args))
		for i, a := range cc.Args {
			args[i] = ft.DoValue(a)
		}
		return []gg.Stmt{gg.AssignOf(
			ft.DoValue(instr),
			gg.Call{
				Func: gfn,
				Args: args,
			},
		)}

	case *ssa.Phi:
		return nil

	case *ssa.UnOp:
		x := ft.DoValue(instr.X)
		return []gg.Stmt{gg.AssignOf(
			ft.DoValue(instr),
			gg.UnaryOf(gg.UnaryOpFromToken(instr.Op), x),
		)}

	case *ssa.ChangeInterface:
	case *ssa.ChangeType:
	case *ssa.Convert:
	case *ssa.SliceToArrayPointer:
	case *ssa.MakeInterface:
	case *ssa.Extract:
	case *ssa.Slice:
	case *ssa.MakeChan:
	case *ssa.Alloc:
	case *ssa.MakeSlice:
	case *ssa.MakeMap:
	case *ssa.Range:
	case *ssa.Next:
	case *ssa.FieldAddr:
	case *ssa.Field:
	case *ssa.IndexAddr:
	case *ssa.Index:
	case *ssa.Lookup:
	case *ssa.TypeAssert:
	case *ssa.MakeClosure:
	case *ssa.Select:

	// void

	case *ssa.If:
		bl := instr.Block()
		x := ft.DoValue(instr.Cond)
		return []gg.Stmt{gg.IfElseOf(
			x,
			gg.Block{Body: ft.doPhiSrc(
				bl,
				bl.Succs[0],
				gg.Goto{Name: ft.blockIdent(bl.Succs[0])},
			)},
			gg.Block{Body: ft.doPhiSrc(
				bl,
				bl.Succs[1],
				gg.Goto{Name: ft.blockIdent(bl.Succs[1])},
			)},
		)}

	case *ssa.Jump:
		dst := instr.Block().Succs[0]
		return ft.doPhiSrc(
			instr.Block(),
			dst,
			gg.Goto{
				Name: ft.blockIdent(dst),
			},
		)

	case *ssa.Return:
		v := ft.DoValue(check.Single(instr.Results))
		return []gg.Stmt{gg.Return{Expr: v}}

	case *ssa.DebugRef:
	case *ssa.RunDefers:
	case *ssa.Panic:
	case *ssa.Send:
	case *ssa.Store:
	case *ssa.Go:
	case *ssa.Defer:
	case *ssa.MapUpdate:

	}
	panic(instr)
}

func (ft *FuncTransformer) DoBlock(block *ssa.BasicBlock) []gg.Stmt {
	var body []gg.Stmt

	if len(block.Preds) > 0 {
		body = append(body, gg.Label{Name: ft.blockIdent(block)})
	}

	for _, instr := range block.Instrs {
		s := ft.DoInstr(instr)
		body = append(body, s...)
	}

	return body
}

func (ft *FuncTransformer) rawType(typ types.Type) gg.Type {
	s := types.TypeString(typ, types.RelativeTo(ft.fn.Package().Pkg))
	s = strings.ReplaceAll(s, "interface{}", "any")
	return gg.Raw{Raw: s}
}

func (ft *FuncTransformer) rawMember(m ssa.Member) string {
	if pkg := m.Package().Pkg; pkg != nil && pkg != ft.fn.Pkg.Pkg {
		return fmt.Sprintf("%s.%s", pkg.Path(), m.Name())
	}
	return m.Name()
}

func (ft *FuncTransformer) rawRelValue(v ssa.Value) string {
	from := ft.fn.Pkg.Pkg
	switch v := v.(type) {

	case ssa.Member: // *Function or *Global
		return v.RelString(from)

	case *ssa.Const:
		return v.RelString(from)

	}
	return v.Name()
}

func (ft *FuncTransformer) Transform() gg.Func {
	ft.bs = make(map[*ssa.BasicBlock]*ftBlock)
	for _, b := range ft.fn.Blocks {
		ft.bs[b] = &ftBlock{
			b:    b,
			dsts: make(map[*ssa.BasicBlock]*ftBlockDst),
		}
	}

	for _, b := range ft.fn.Blocks {
		for _, i := range b.Instrs {
			if phi, ok := i.(*ssa.Phi); ok {
				for predi, predv := range phi.Edges {
					ftpred := ft.bs[b.Preds[predi]]

					ftd := ftpred.dsts[b]
					if ftd == nil {
						ftd = &ftBlockDst{b: b}
						ftpred.dsts[b] = ftd
					}

					ftd.phis = append(ftd.phis, ftBlockDstPhi{
						phi: phi,
						v:   predv,
					})
				}
			}
		}
	}

	gfn := gg.Func{
		Name: gg.NewIdent(ft.fn.Name()),
	}

	for _, p := range ft.fn.Params {
		gfn.Params = append(gfn.Params, gg.Param{
			Name: gg.NewIdent(p.Name()),
			Type: ft.rawType(p.Type()),
		})
	}

	rtt := ft.fn.Signature.Results()
	if rtt.Len() > 0 {
		check.Equal(rtt.Len(), 1)
		gfn.Type = ft.rawType(rtt.At(0).Type())
	}

	var body []gg.Stmt

	ft.vns = make(map[ssa.Value]int)
	for _, b := range ft.fn.Blocks {
		for _, i := range b.Instrs {
			if v, ok := i.(ssa.Value); ok {
				ft.vns[v] = len(ft.vns)
				id := ft.instrIdent(i)
				body = append(
					body,
					gg.Var{
						Name: id,
						Type: ft.rawType(v.Type()),
					},
					gg.Assign{
						Var: gg.Ident{
							Name: "_",
						},
						Value: id,
					},
				)
			}
		}
	}

	for _, b := range ft.fn.Blocks {
		s := ft.DoBlock(b)
		body = append(body, s...)
	}

	gfn.Body = &gg.Block{
		Body: body,
	}

	return gfn
}

func testSsa(t *testing.T, patterns []string) {
	pkgs := check.Must1(packages.Load(
		&packages.Config{
			Mode: runner.PackageMode,
		},
		patterns...,
	))

	pkg := check.Single(pkgs)
	ssainfo := check.Must1(run(
		pkg.Fset,
		pkg.Syntax,
		pkg.Types,
		pkg.TypesInfo,
	))

	_ = ssainfo
	fmt.Println(ssainfo)

	for i := 0; i < 2; i++ {
		fn := ssainfo.SrcFuncs[i]
		ft := FuncTransformer{fn: fn}

		gfn := ft.Transform()
		fmt.Println(gg.RenderString(gfn))
		fmt.Println()
	}

	//results := runner.Run(buildssa.Analyzer, pkgs, runner.Options{})
	//ssainfo := results[0].Result.(*buildssa.SSA)

	//fmt.Println(ssainfo)
	//for _, f := range ssainfo.SrcFuncs {
	//	fmt.Println(f)
	//	for _, b := range f.Blocks {
	//		for _, i := range b.Instrs {
	//			fmt.Println(i)
	//		}
	//		fmt.Println()
	//	}
	//	fmt.Println()
	//}
}

func TestSsa(t *testing.T) {
	testSsa(t, []string{"github.com/wrmsr/bane/x/go/ssa/test"})
}

func TestSsa2(t *testing.T) {
	testSsa(t, []string{"github.com/wrmsr/bane/core/ndarray"})
}

//

func init() {
	math.Jn(0, 0)
}

//

func A(x int, y int) int {
	var _ssa_0 bool
	_ = _ssa_0
	var _ssa_1 int
	_ = _ssa_1
	var _ssa_2 int
	_ = _ssa_2
	var _ssa_3 int
	_ = _ssa_3
	var _ssa_4 int
	_ = _ssa_4
	var _ssa_5 int
	_ = _ssa_5
	_ssa_0 = x > 0
	if _ssa_0 {
		goto _ssa_block_1
	} else {
		_ssa_2 = x
		goto _ssa_block_2
	}
_ssa_block_1:
	_ssa_1 = x + 1
	_ssa_2 = _ssa_1
	goto _ssa_block_2
_ssa_block_2:
	_ssa_3 = -_ssa_2
	_ssa_4 = _ssa_3 + y
	_ssa_5 = _ssa_4 + 1
	return _ssa_5
}

func B(x int, y int) int {
	var _ssa_0 float64
	_ = _ssa_0
	var _ssa_1 int
	_ = _ssa_1
	var _ssa_2 int
	_ = _ssa_2
	var _ssa_3 int
	_ = _ssa_3
	var _ssa_4 int
	_ = _ssa_4
	var _ssa_5 int
	_ = _ssa_5
	var _ssa_6 int
	_ = _ssa_6
	var _ssa_7 bool
	_ = _ssa_7
	_ssa_0 = math.Jn(x, 0)
	_ssa_5 = x
	_ssa_6 = 0
	goto _ssa_block_3
_ssa_block_1:
	_ssa_1 = A(_ssa_5, y)
	_ssa_2 = _ssa_1 * 2
	_ssa_3 = _ssa_5 + _ssa_2
	_ssa_4 = _ssa_6 + 1
	_ssa_5 = _ssa_3
	_ssa_6 = _ssa_4
	goto _ssa_block_3
_ssa_block_2:
	return _ssa_5
_ssa_block_3:
	_ssa_7 = _ssa_6 < y
	if _ssa_7 {
		goto _ssa_block_1
	} else {
		goto _ssa_block_2
	}
}
