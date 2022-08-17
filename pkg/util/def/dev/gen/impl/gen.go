//go:build !nodev

package impl

import (
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/tools/go/packages"

	"github.com/wrmsr/bane/pkg/util/def"
	gg "github.com/wrmsr/bane/pkg/util/go/gen"
	iou "github.com/wrmsr/bane/pkg/util/io"
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
	return gg.FuncTypeOf(
		gg.Func{
			Params: []gg.Param{
				{Type: gg.PtrOf(elem)},
			},
		},
	)
}

func (fg *FileGen) Gen() string {
	fg.genStructs()

	doInit := gg.Func{
		Body: gg.NewBlock(slices.DeepFlatten[gg.Stmt](
			gg.ShortVarOf(
				gg.IdentOf("spec"),
				gg.CallOf(gg.SelectOf("def", "X_getPackageSpec"))),
			fg.initStmts,
		)...)}

	fg.decls = slices.DeepFlatten[gg.Decl](
		gg.StmtDeclOf(gg.Var{
			Name: gg.IdentOf("_def_init_once"),
			Type: gg.NameTypeOf("sync.Once")}),

		gg.Func{
			Name: gg.NewIdent("_def_init"),
			Body: gg.NewBlock(
				gg.CallOf(
					gg.SelectOf("_def_init_once", "Do"),
					gg.FuncExpr{Func: doInit}))},

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
