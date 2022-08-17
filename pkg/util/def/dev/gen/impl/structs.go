//go:build !nodev

package impl

import (
	"fmt"
	"go/ast"
	"strconv"
	"strings"

	"github.com/wrmsr/bane/pkg/util/def"
	gg "github.com/wrmsr/bane/pkg/util/go/gen"
)

func (fg *FileGen) genStruct(ss *def.StructSpec) {
	sn := ss.Type().(string)
	sName := gg.IdentOf(sn)
	ssName := gg.IdentOf(fmt.Sprintf("struct_spec__%s", sn))

	zn := fmt.Sprintf("zero_%s", sn)
	fg.initStmts = append(fg.initStmts,
		gg.Blank{},

		gg.VarOf(zn, gg.NameTypeOf(sn)),

		gg.ShortVarOf(
			ssName,
			gg.CallOf(
				gg.SelectOf(gg.IdentOf("spec"), gg.IdentOf("Struct")),
				gg.CallOf(
					gg.SelectOf(gg.IdentOf("reflect"), gg.IdentOf("TypeOf")),
					gg.IdentOf(zn)))),

		gg.AssignOf(gg.IdentOf("_"), ssName),
	)

	var sfs []gg.StructField
	var dflVds []gg.Var

	initStmts := []gg.Stmt{
		gg.ExprStmtOf(gg.CallOf(gg.IdentOf("_def_init"))),
	}

	rcvr := ss.Receiver()
	if rcvr == "" {
		rcvr = strings.ToLower(string([]rune(strings.TrimLeft(sn, "_"))[0]))
	}

	for _, fs := range ss.Fields() {
		sfs = append(sfs, gg.StructField{Name: gg.IdentOf(fs.Name()), Type: fg.ti.importedType(fs.Type())})

		fsName := gg.IdentOf(fmt.Sprintf("field_spec__%s__%s", sn, fs.Name()))

		fg.initStmts = append(fg.initStmts,
			gg.Blank{},

			gg.ShortVarOf(
				fsName,
				gg.CallOf(
					gg.SelectOf(ssName, gg.IdentOf("Field")),
					gg.LitOf(fmt.Sprintf("\"%s\"", fs.Name())))),

			gg.AssignOf(gg.IdentOf("_"), fsName),
		)

		if fs.Default() != nil {
			dflName := gg.IdentOf(fmt.Sprintf("_def_field_default__%s__%s", sn, fs.Name()))

			dflVds = append(dflVds,
				gg.VarOf(dflName, fg.ti.importedType(fs.Type())))

			fg.initStmts = append(fg.initStmts,
				gg.Blank{},

				gg.AssignOf(
					dflName,
					gg.TypeAssertOf(
						gg.CallOf(gg.SelectOf(fsName, gg.IdentOf("Default"))),
						fg.ti.importedType(fs.Type()))))

			initStmts = append(initStmts,
				gg.Blank{},

				gg.AssignOf(gg.SelectOf(gg.IdentOf("f"), gg.IdentOf(fs.Name())), dflName))
		}
	}

	fg.decls = append(fg.decls, gg.Struct{Name: sName, Fields: sfs})

	if len(dflVds) > 0 {
		fg.decls = append(fg.decls,
			gg.StmtDeclOf(gg.VarsOf(dflVds...)))
	}

	if len(ss.Inits()) > 0 {
		fg.initStmts = append(fg.initStmts, gg.Blank{})
		initStmts = append(initStmts, gg.Blank{})

		var initVds []gg.Var
		for i, o := range ss.Inits() {
			var initName gg.Ident
			switch o := o.(type) {
			case *ast.FuncLit:
				initName = gg.IdentOf(fmt.Sprintf("_def_struct_init__%s__%d", sn, i))

				initVds = append(initVds,
					gg.VarOf(initName, newPtrFuncType(gg.NameTypeOf(sName))))

				fg.initStmts = append(fg.initStmts,
					gg.AssignOf(
						initName,
						gg.TypeAssertOf(
							gg.IndexOf(gg.CallOf(gg.SelectOf(ssName, gg.IdentOf("Inits"))), gg.LitOf(strconv.Itoa(i))),
							newPtrFuncType(gg.NameTypeOf(sName)))))

				initStmts = append(initStmts,
					gg.ExprStmtOf(gg.CallOf(initName, gg.IdentOf("f"))))

			case *ast.FuncDecl:
				initStmts = append(initStmts,
					gg.ExprStmtOf(gg.CallOf(gg.SelectOf(gg.IdentOf("f"), gg.IdentOf(o.Name.Name)))))

			default:
				panic(i)
			}
		}

		fg.decls = append(fg.decls,
			gg.StmtDeclOf(gg.VarsOf(initVds...)))
	}

	fg.decls = append(fg.decls,
		gg.Func{
			Receiver: &gg.Param{Name: gg.NewIdent("f"), Type: gg.PtrOf(gg.NameTypeOf(sName))},
			Name:     gg.NewIdent("init"),
			Body:     gg.NewBlock(initStmts...)})
}

func (fg *FileGen) genStructs() {
	i := 0
	for _, ss := range fg.ps.Structs() {
		if i > 0 {
			fg.initStmts = append(fg.initStmts, gg.Blank{})
		}
		i++
		fg.genStruct(ss)
	}
}
