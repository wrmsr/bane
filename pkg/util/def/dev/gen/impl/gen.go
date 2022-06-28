//go:build !nodev

package impl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wrmsr/bane/pkg/util/def"
	gg "github.com/wrmsr/bane/pkg/util/gogen"
	iou "github.com/wrmsr/bane/pkg/util/io"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	"github.com/wrmsr/bane/pkg/util/slices"
)

type FileGen struct {
	ps *def.PackageSpec

	decls     []gg.Decl
	initStmts []gg.Stmt
}

func NewFileGen(pkg *def.PackageSpec) *FileGen {
	return &FileGen{
		ps: pkg,
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

func (fg *FileGen) genStruct(ss *def.StructSpec) {
	sName := gg.NewIdent(ss.Name())
	ssName := gg.NewIdent(fmt.Sprintf("struct_spec__%s", ss.Name()))

	fg.initStmts = append(fg.initStmts,
		gg.NewBlank(),

		gg.NewShortVar(ssName,
			gg.NewCall(
				gg.NewSelect(gg.NewIdent("spec"), gg.NewIdent("Struct")),
				gg.NewLit(fmt.Sprintf("\"%s\"", ss.Name())))),

		gg.NewAssign(gg.NewIdent("_"), ssName),
	)

	var sfs []gg.StructField
	var dflVds []gg.Var

	initStmts := []gg.Stmt{
		gg.NewExprStmt(gg.NewCall(gg.NewIdent("_def_init"))),
	}

	rcvr := ss.Receiver()
	if rcvr == "" {
		rcvr = strings.ToLower(string([]rune(strings.TrimLeft(ss.Name(), "_"))[0]))
	}

	for _, fs := range ss.Fields() {
		sfs = append(sfs, gg.NewStructField(gg.NewIdent(fs.Name()), gg.NewNameType(gg.NewIdent("int"))))

		fsName := gg.NewIdent(fmt.Sprintf("field_spec__%s__%s", ss.Name(), fs.Name()))

		fg.initStmts = append(fg.initStmts,
			gg.NewBlank(),

			gg.NewShortVar(fsName,
				gg.NewCall(
					gg.NewSelect(ssName, gg.NewIdent("Field")),
					gg.NewLit(fmt.Sprintf("\"%s\"", fs.Name())))),

			gg.NewAssign(gg.NewIdent("_"), fsName),
		)

		if fs.Default() != nil {
			dflName := gg.NewIdent(fmt.Sprintf("_def_field_default__%s__%s", ss.Name(), fs.Name()))

			dflVds = append(dflVds,
				gg.NewVar(dflName, opt.Just[gg.Type](gg.NewNameType(gg.NewIdent("int"))), opt.None[gg.Expr]()))

			fg.initStmts = append(fg.initStmts,
				gg.NewBlank(),

				gg.NewAssign(dflName,
					gg.NewTypeAssert(
						gg.NewCall(gg.NewSelect(fsName, gg.NewIdent("Default"))),
						gg.NewNameType(gg.NewIdent("int")))))

			initStmts = append(initStmts,
				gg.NewBlank(),

				gg.NewAssign(gg.NewSelect(gg.NewIdent("f"), gg.NewIdent(fs.Name())), dflName))
		}
	}

	fg.decls = append(fg.decls, gg.NewStruct(sName, sfs...))

	if len(dflVds) > 0 {
		fg.decls = append(fg.decls,
			gg.NewStmtDecl(gg.NewVars(dflVds)))
	}

	if len(ss.Inits()) > 0 {
		fg.initStmts = append(fg.initStmts, gg.NewBlank())
		initStmts = append(initStmts, gg.NewBlank())

		var initVds []gg.Var
		for i, _ := range ss.Inits() {
			initName := gg.NewIdent(fmt.Sprintf("_def_struct_init__%s__%d", ss.Name(), i))

			initVds = append(initVds,
				gg.NewVar(initName,
					opt.Just[gg.Type](newPtrFuncType(gg.NewNameType(sName))), opt.None[gg.Expr]()))

			fg.initStmts = append(fg.initStmts,
				gg.NewAssign(initName,
					gg.NewTypeAssert(
						gg.NewIndex(gg.NewCall(gg.NewSelect(ssName, gg.NewIdent("Inits"))), gg.NewLit(strconv.Itoa(i))),
						newPtrFuncType(gg.NewNameType(sName)))))

			initStmts = append(initStmts,
				gg.NewExprStmt(gg.NewCall(initName, gg.NewIdent("f"))))
		}

		fg.decls = append(fg.decls,
			gg.NewStmtDecl(gg.NewVars(initVds)))
	}

	fg.decls = append(fg.decls,
		gg.NewFunc(
			opt.Just(gg.NewParam(opt.Just(gg.NewIdent("f")), gg.NewPtr(gg.NewNameType(sName)))),
			opt.Just(gg.NewIdent("init")),
			nil,
			opt.None[gg.Type](),
			opt.Just(gg.NewBlock(initStmts...))))
}

func (fg *FileGen) Gen() string {
	i := 0
	for _, ss := range fg.ps.Structs() {
		if i > 0 {
			fg.initStmts = append(fg.initStmts, gg.NewBlank())
		}
		i++
		fg.genStruct(ss)
	}

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

	var sb strings.Builder
	sw := iou.NewIndentWriter(iou.NewDiscardStringWriter(&sb), "\t")
	for _, n := range fg.decls {
		gg.Render(sw, n)
	}
	return sb.String()
}
