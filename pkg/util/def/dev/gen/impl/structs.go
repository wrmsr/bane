//go:build !nodev

package impl

import (
	"fmt"
	"go/ast"
	"strconv"
	"strings"

	"github.com/wrmsr/bane/pkg/util/def"
	gg "github.com/wrmsr/bane/pkg/util/go/gen"
	opt "github.com/wrmsr/bane/pkg/util/optional"
)

func (fg *FileGen) genStruct(ss *def.StructSpec) {
	sn := ss.Type().(string)
	sName := gg.NewIdent(sn)
	ssName := gg.NewIdent(fmt.Sprintf("struct_spec__%s", sn))

	zn := fmt.Sprintf("zero_%s", sn)
	fg.initStmts = append(fg.initStmts,
		gg.NewBlank(),

		gg.NewVar(
			gg.NewIdent(zn), opt.Just[gg.Type](gg.NewNameType(gg.NewIdent(sn))), opt.None[gg.Expr]()),

		gg.NewShortVar(ssName,
			gg.NewCall(
				gg.NewSelect(gg.NewIdent("spec"), gg.NewIdent("Struct")),
				gg.NewCall(
					gg.NewSelect(gg.NewIdent("reflect"), gg.NewIdent("TypeOf")),
					gg.NewIdent(zn)))),

		gg.NewAssign(gg.NewIdent("_"), ssName),
	)

	var sfs []gg.StructField
	var dflVds []gg.Var

	initStmts := []gg.Stmt{
		gg.NewExprStmt(gg.NewCall(gg.NewIdent("_def_init"))),
	}

	rcvr := ss.Receiver()
	if rcvr == "" {
		rcvr = strings.ToLower(string([]rune(strings.TrimLeft(sn, "_"))[0]))
	}

	for _, fs := range ss.Fields() {
		sfs = append(sfs, gg.NewStructField(gg.NewIdent(fs.Name()), fg.ti.importedType(fs.Type())))

		fsName := gg.NewIdent(fmt.Sprintf("field_spec__%s__%s", sn, fs.Name()))

		fg.initStmts = append(fg.initStmts,
			gg.NewBlank(),

			gg.NewShortVar(fsName,
				gg.NewCall(
					gg.NewSelect(ssName, gg.NewIdent("Field")),
					gg.NewLit(fmt.Sprintf("\"%s\"", fs.Name())))),

			gg.NewAssign(gg.NewIdent("_"), fsName),
		)

		if fs.Default() != nil {
			dflName := gg.NewIdent(fmt.Sprintf("_def_field_default__%s__%s", sn, fs.Name()))

			dflVds = append(dflVds,
				gg.NewVar(dflName, opt.Just[gg.Type](fg.ti.importedType(fs.Type())), opt.None[gg.Expr]()))

			fg.initStmts = append(fg.initStmts,
				gg.NewBlank(),

				gg.NewAssign(dflName,
					gg.NewTypeAssert(
						gg.NewCall(gg.NewSelect(fsName, gg.NewIdent("Default"))),
						fg.ti.importedType(fs.Type()))))

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
		for i, o := range ss.Inits() {
			var initName gg.Ident
			switch o := o.(type) {
			case *ast.FuncLit:
				initName = gg.NewIdent(fmt.Sprintf("_def_struct_init__%s__%d", sn, i))

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

			case *ast.FuncDecl:
				initStmts = append(initStmts,
					gg.NewExprStmt(gg.NewCall(gg.NewSelect(gg.NewIdent("f"), gg.NewIdent(o.Name.Name)))))

			default:
				panic(i)
			}
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

func (fg *FileGen) genStructs() {
	i := 0
	for _, ss := range fg.ps.Structs() {
		if i > 0 {
			fg.initStmts = append(fg.initStmts, gg.NewBlank())
		}
		i++
		fg.genStruct(ss)
	}
}
