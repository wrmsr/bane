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
	sName := gg.IdentOf(sn)
	ssName := gg.IdentOf(fmt.Sprintf("struct_spec__%s", sn))

	zn := fmt.Sprintf("zero_%s", sn)
	fg.initStmts = append(fg.initStmts,
		gg.NewBlank(),

		gg.NewVar(
			gg.IdentOf(zn), opt.Just[gg.Type](gg.NewNameType(gg.IdentOf(sn))), opt.None[gg.Expr]()),

		gg.NewShortVar(ssName,
			gg.NewCall(
				gg.NewSelect(gg.IdentOf("spec"), gg.IdentOf("Struct")),
				gg.NewCall(
					gg.NewSelect(gg.IdentOf("reflect"), gg.IdentOf("TypeOf")),
					gg.IdentOf(zn)))),

		gg.NewAssign(gg.IdentOf("_"), ssName),
	)

	var sfs []gg.StructField
	var dflVds []gg.Var

	initStmts := []gg.Stmt{
		gg.NewExprStmt(gg.NewCall(gg.IdentOf("_def_init"))),
	}

	rcvr := ss.Receiver()
	if rcvr == "" {
		rcvr = strings.ToLower(string([]rune(strings.TrimLeft(sn, "_"))[0]))
	}

	for _, fs := range ss.Fields() {
		sfs = append(sfs, gg.StructField{Name: gg.IdentOf(fs.Name()), Type: fg.ti.importedType(fs.Type())})

		fsName := gg.IdentOf(fmt.Sprintf("field_spec__%s__%s", sn, fs.Name()))

		fg.initStmts = append(fg.initStmts,
			gg.NewBlank(),

			gg.NewShortVar(fsName,
				gg.NewCall(
					gg.NewSelect(ssName, gg.IdentOf("Field")),
					gg.NewLit(fmt.Sprintf("\"%s\"", fs.Name())))),

			gg.NewAssign(gg.IdentOf("_"), fsName),
		)

		if fs.Default() != nil {
			dflName := gg.IdentOf(fmt.Sprintf("_def_field_default__%s__%s", sn, fs.Name()))

			dflVds = append(dflVds,
				gg.NewVar(dflName, opt.Just[gg.Type](fg.ti.importedType(fs.Type())), opt.None[gg.Expr]()))

			fg.initStmts = append(fg.initStmts,
				gg.NewBlank(),

				gg.NewAssign(dflName,
					gg.NewTypeAssert(
						gg.NewCall(gg.NewSelect(fsName, gg.IdentOf("Default"))),
						fg.ti.importedType(fs.Type()))))

			initStmts = append(initStmts,
				gg.NewBlank(),

				gg.NewAssign(gg.NewSelect(gg.IdentOf("f"), gg.IdentOf(fs.Name())), dflName))
		}
	}

	fg.decls = append(fg.decls, gg.Struct{Name: sName, Fields: sfs})

	if len(dflVds) > 0 {
		fg.decls = append(fg.decls,
			gg.StmtDecl{Stmt: gg.NewVars(dflVds)})
	}

	if len(ss.Inits()) > 0 {
		fg.initStmts = append(fg.initStmts, gg.NewBlank())
		initStmts = append(initStmts, gg.NewBlank())

		var initVds []gg.Var
		for i, o := range ss.Inits() {
			var initName gg.Ident
			switch o := o.(type) {
			case *ast.FuncLit:
				initName = gg.IdentOf(fmt.Sprintf("_def_struct_init__%s__%d", sn, i))

				initVds = append(initVds,
					gg.NewVar(initName,
						opt.Just[gg.Type](newPtrFuncType(gg.NewNameType(sName))), opt.None[gg.Expr]()))

				fg.initStmts = append(fg.initStmts,
					gg.NewAssign(initName,
						gg.NewTypeAssert(
							gg.NewIndex(gg.NewCall(gg.NewSelect(ssName, gg.IdentOf("Inits"))), gg.NewLit(strconv.Itoa(i))),
							newPtrFuncType(gg.NewNameType(sName)))))

				initStmts = append(initStmts,
					gg.NewExprStmt(gg.NewCall(initName, gg.IdentOf("f"))))

			case *ast.FuncDecl:
				initStmts = append(initStmts,
					gg.NewExprStmt(gg.NewCall(gg.NewSelect(gg.IdentOf("f"), gg.IdentOf(o.Name.Name)))))

			default:
				panic(i)
			}
		}

		fg.decls = append(fg.decls,
			gg.StmtDecl{Stmt: gg.NewVars(initVds)})
	}

	fg.decls = append(fg.decls,
		gg.Func{
			Receiver: &gg.Param{Name: gg.NewIdent("f"), Type: gg.NewPtr(gg.NewNameType(sName))},
			Name:     gg.NewIdent("init"),
			Body:     gg.BlockOf(initStmts...)})
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
