//go:build !nodev

package impl

import (
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/tools/go/packages"

	"github.com/wrmsr/bane/pkg/util/def"
	gg "github.com/wrmsr/bane/pkg/util/go/gen"
	iou "github.com/wrmsr/bane/pkg/util/io"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	"github.com/wrmsr/bane/pkg/util/slices"
	stru "github.com/wrmsr/bane/pkg/util/strings"
)

type FileGen struct {
	mod *modfile.File
	pkg *packages.Package
	ps  *def.PackageSpec

	ti *typeImporter

	decls     []gg.Decl
	initStmts []gg.Stmt
}

func NewFileGen(
	mod *modfile.File,
	pkg *packages.Package,
	ps *def.PackageSpec,
) *FileGen {
	return &FileGen{
		mod: mod,
		pkg: pkg,
		ps:  ps,

		ti: newTypeImporter(ps, mod.Module.Mod.Path, []string{
			"reflect",
			"sync",
		}),
	}
}

func newPtrFuncType(elem gg.Type) gg.FuncType {
	return gg.NewFuncType(
		gg.Func{
			Params: []gg.Param{
				{Type: gg.NewPtr(elem)},
			},
		},
	)
}

func (fg *FileGen) Gen() string {
	fg.genStructs()

	doInit := gg.Func{
		Body: gg.BlockOf(slices.DeepFlatten[gg.Stmt](
			gg.ShortVar{
				Name:  gg.IdentOf("spec"),
				Value: gg.NewCall(gg.NewSelect(gg.IdentOf("def"), gg.IdentOf("X_getPackageSpec")))},
			fg.initStmts,
		)...)}

	fg.decls = slices.DeepFlatten[gg.Decl](
		gg.StmtDecl{
			Stmt: gg.NewVar(
				gg.IdentOf("_def_init_once"),
				opt.Just[gg.Type](gg.NewNameType(gg.IdentOf("sync.Once"))), opt.None[gg.Expr]())},

		gg.Func{
			Name: gg.NewIdent("_def_init"),
			Body: gg.BlockOf(
				gg.NewExprStmt(gg.NewCall(
					gg.NewSelect(gg.IdentOf("_def_init_once"), gg.IdentOf("Do")),
					gg.NewFuncExpr(doInit))))},

		fg.decls)

	imps := fg.ti.imports()
	if len(imps) > 0 {
		fg.decls = slices.DeepFlatten[gg.Decl](gg.Imports{Imports: imps}, fg.decls)
	}

	_, pn, _ := stru.LastCut(fg.ps.Name(), "/")
	fg.decls = slices.DeepFlatten[gg.Decl](gg.Package{Name: gg.IdentOf(pn)}, fg.decls)

	var sb strings.Builder
	sw := iou.NewIndentWriter(iou.NewDiscardStringWriter(&sb), "\t")
	for i, n := range fg.decls {
		if i > 0 {
			sw.WriteString("\n")
		}
		gg.Render(sw, n)
	}
	return sb.String()
}
