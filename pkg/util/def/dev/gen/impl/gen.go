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

	ti  *typeImporter
	rsv *Resolver

	decls     gg.Decls
	initStmts gg.Stmts
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

		rsv: newResolver(pkg),
	}
}

func newPtrFuncType(elem any) gg.FuncType {
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
	fg.inlineFuncs()

	doInit := gg.Func{
		Body: &gg.Block{Body: slices.DeepFlatten[gg.Stmt](
			gg.ShortVarOf(
				"spec",
				gg.CallOf(gg.SelectOf("def", "X_getPackageSpec"))),
			fg.initStmts,
		)}}

	fg.decls = slices.DeepFlatten[gg.Decl](
		gg.StmtDeclOf(gg.Var{
			Name: gg.IdentOf("_def_init_once"),
			Type: gg.TypeOf("sync.Once")}),

		gg.Func{
			Name: gg.NewIdent("_def_init"),
			Body: gg.NewBlock(
				gg.CallOf(
					gg.SelectOf("_def_init_once", "Do"),
					doInit))},

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
