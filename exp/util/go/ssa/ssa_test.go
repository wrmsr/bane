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
	"go/ast"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"

	"github.com/wrmsr/bane/exp/util/go/analysis/runner"
	"github.com/wrmsr/bane/pkg/util/check"
)

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

func TestSsa(t *testing.T) {
	pkgs := check.Must1(packages.Load(
		&packages.Config{
			Mode: runner.PackageMode,
		},
		"github.com/wrmsr/bane/exp/util/go/ssa/test",
	))

	pkg := check.Single(pkgs)
	ssainfo := check.Must1(run(
		pkg.Fset,
		pkg.Syntax,
		pkg.Types,
		pkg.TypesInfo,
	))

	_ = ssainfo

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