//go:build !nodev

package impl

import (
	"strings"

	"golang.org/x/mod/modfile"
	"golang.org/x/tools/go/packages"

	"github.com/wrmsr/bane/pkg/util/def"
	gg "github.com/wrmsr/bane/pkg/util/gogen"
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

		ti: newTypeImporter(ps, mod.Module.Mod.Path, []string{"sync"}),
	}
}

func newPtrFuncType(elem gg.Type) gg.FuncType {
	return gg.NewFuncType(
		gg.NewFunc(
			opt.None[gg.Param](),
			opt.None[gg.Ident](),
			[]gg.Param{
				gg.NewParam(opt.None[gg.Ident](), gg.NewPtr(elem)),
			},
			opt.None[gg.Type](),
			opt.None[gg.Block](),
		),
	)
}

func (fg *FileGen) Gen() string {
	fg.genStructs()

	doInit := gg.NewFunc(
		opt.None[gg.Param](),
		opt.None[gg.Ident](),
		nil,
		opt.None[gg.Type](),
		opt.Just(gg.NewBlock(slices.DeepFlatten[gg.Stmt](
			gg.NewShortVar(
				gg.NewIdent("spec"),
				gg.NewCall(gg.NewSelect(gg.NewIdent("def"), gg.NewIdent("X_getPackageSpec")))),
			fg.initStmts,
		)...)))

	fg.decls = slices.DeepFlatten[gg.Decl](
		gg.NewStmtDecl(
			gg.NewVar(
				gg.NewIdent("_def_init_once"),
				opt.Just[gg.Type](gg.NewNameType(gg.NewIdent("sync.Once"))), opt.None[gg.Expr]())),

		gg.NewFunc(
			opt.None[gg.Param](),
			opt.Just(gg.NewIdent("_def_init")),
			nil,
			opt.None[gg.Type](),
			opt.Just(gg.NewBlock(
				gg.NewExprStmt(gg.NewCall(
					gg.NewSelect(gg.NewIdent("_def_init_once"), gg.NewIdent("Do")),
					gg.NewFuncExpr(doInit)))))),

		fg.decls)

	imps := fg.ti.imports()
	if len(imps) > 0 {
		fg.decls = slices.DeepFlatten[gg.Decl](gg.NewImports(imps...), fg.decls)
	}

	_, pn, _ := stru.LastCut(fg.ps.Name(), "/")
	fg.decls = slices.DeepFlatten[gg.Decl](gg.NewPackage(gg.NewIdent(pn)), fg.decls)

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
