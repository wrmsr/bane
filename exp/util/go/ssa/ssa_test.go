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
	"testing"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"

	"github.com/wrmsr/bane/exp/util/go/analysis/runner"
	"github.com/wrmsr/bane/pkg/util/check"
	gg "github.com/wrmsr/bane/pkg/util/go/gen"
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

type FuncTransformer struct {
	fn *ssa.Function

	vns map[ssa.Value]int
}

func (ft *FuncTransformer) DoValue(value ssa.Value) gg.Expr {
	if _, ok := value.(ssa.Instruction); ok {
		return gg.IdentOf(fmt.Sprintf("_ssa_%d", ft.vns[value]))
	}
	switch value := value.(type) {
	case *ssa.Parameter:
		return gg.IdentOf(value.Name())
	case *ssa.Const:
		return gg.LitOf(value.Value.String())
	}
	panic(value)
}

func (ft *FuncTransformer) DoInstr(instr ssa.Instruction) gg.Stmt {
	switch instr := instr.(type) {
	case *ssa.BinOp:
		x := ft.DoValue(instr.X)
		y := ft.DoValue(instr.Y)
		return gg.AssignOf(
			ft.DoValue(instr),
			gg.InfixOf(gg.AddOp, x, y),
		)

	case *ssa.Return:
		v := ft.DoValue(check.Single(instr.Results))
		return gg.Return{Expr: v}

	// register
	case *ssa.UnOp:
	case *ssa.Call:
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
	case *ssa.Phi:
	case *ssa.Select:

	// void
	case *ssa.DebugRef:
	case *ssa.RunDefers:
	case *ssa.Panic:
	case *ssa.Send:
	case *ssa.Store:
	case *ssa.If:
	case *ssa.Go:
	case *ssa.Jump:
	case *ssa.Defer:
	case *ssa.MapUpdate:
	}
	panic(instr)
}

func (ft *FuncTransformer) DoBlock(block *ssa.BasicBlock) {
	for _, instr := range block.Instrs {
		s := ft.DoInstr(instr)
		fmt.Println(gg.RenderString(s))
	}
}

func (ft *FuncTransformer) Transform() {
	ft.vns = make(map[ssa.Value]int)
	for _, b := range ft.fn.Blocks {
		for _, i := range b.Instrs {
			if v, ok := i.(ssa.Value); ok {
				ft.vns[v] = len(ft.vns)
			}
		}
	}

	for _, b := range ft.fn.Blocks {
		ft.DoBlock(b)
	}
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

	fn := ssainfo.SrcFuncs[0]
	ft := FuncTransformer{fn: fn}
	ft.Transform()

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
	testSsa(t, []string{"github.com/wrmsr/bane/exp/util/go/ssa/test"})
}

func TestSsa2(t *testing.T) {
	testSsa(t, []string{"github.com/wrmsr/bane/pkg/util/ndarray"})
}
