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
	varDecls  []gg.Var
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

	var sfs []gg.StructField

	initStmts := []gg.Stmt{
		gg.NewExprStmt(gg.NewCall(gg.NewIdent("_def_init"))),
	}

	for _, fs := range ss.Fields() {
		sfs = append(sfs, gg.NewStructField(gg.NewIdent(fs.Name()), gg.NewNameType(gg.NewIdent("int"))))

		fg.initStmts = append(fg.initStmts,
			gg.NewBlank(),
			gg.NewShortVar(ssName,
				gg.NewCall(
					gg.NewField(gg.NewIdent("spec"), gg.NewIdent("Struct")),
					gg.NewLit(fmt.Sprintf("\"%s\"", ss.Name())))),
			gg.NewAssign(gg.NewIdent("_"), ssName),
		)

		if fs.Default() != nil {
			dflName := gg.NewIdent(fmt.Sprintf("_def_field_default__%s__%s", ss.Name(), fs.Name()))

			fg.varDecls = append(fg.varDecls,
				gg.NewVar(dflName, opt.Just[gg.Type](gg.NewNameType(gg.NewIdent("int"))), opt.None[gg.Expr]()))

			fg.initStmts = append(fg.initStmts,
				gg.NewBlank(),
				gg.NewAssign(dflName,
					gg.NewTypeAssert(
						gg.NewCall(gg.NewField(ssName, gg.NewIdent("Default"))),
						gg.NewNameType(gg.NewIdent("int")))))

			initStmts = append(initStmts,
				gg.NewBlank(),
				gg.NewAssign(gg.NewField(gg.NewIdent("f"), gg.NewIdent(fs.Name())), dflName))
		}
	}

	fg.decls = append(fg.decls, gg.NewStruct(sName, sfs...))

	if len(ss.Inits()) > 0 {
		fg.initStmts = append(fg.initStmts, gg.NewBlank())
		initStmts = append(initStmts, gg.NewBlank())
	}
	for i, _ := range ss.Inits() {
		initName := gg.NewIdent(fmt.Sprintf("_def_struct_init__%s__%d", ss.Name(), i))

		fg.varDecls = append(fg.varDecls,
			gg.NewVar(initName,
				opt.Just[gg.Type](newPtrFuncType(gg.NewNameType(sName))), opt.None[gg.Expr]()))

		fg.initStmts = append(fg.initStmts,
			gg.NewAssign(initName,
				gg.NewTypeAssert(
					gg.NewIndex(gg.NewCall(gg.NewField(ssName, gg.NewIdent("Inits"))), gg.NewLit(strconv.Itoa(i))),
					newPtrFuncType(gg.NewNameType(sName)))))

		initStmts = append(initStmts,
			gg.NewExprStmt(gg.NewCall(initName, gg.NewIdent("f"))))
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
				gg.NewCall(gg.NewField(gg.NewIdent("def"), gg.NewIdent("X_getPackageSpec")))),
			fg.initStmts,
		)...)))

	fg.decls = append(fg.decls,
		gg.NewFunc(
			opt.None[gg.Param](),
			opt.Just(gg.NewIdent("_def_init")),
			nil,
			opt.None[gg.Type](),
			opt.Just(gg.NewBlock(
				gg.NewExprStmt(gg.NewCall(
					gg.NewField(gg.NewIdent("_def_init_once"), gg.NewIdent("Do")),
					gg.NewFuncExpr(doInit)))))))

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
