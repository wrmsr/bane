//go:build !nodev

package impl

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/wrmsr/bane/pkg/util/def"
	gg "github.com/wrmsr/bane/pkg/util/go/gen"
)

func (fg *FileGen) genStruct(ss *def.StructSpec) {
	sn := ss.Type().(string)
	sName := gg.IdentOf(sn)
	ssName := gg.IdentOf(fmt.Sprintf("struct_spec__%s", sn))

	zn := fmt.Sprintf("zero_%s", sn)
	fg.initStmts.Append(
		gg.Blank{},

		gg.VarOf(zn, gg.TypeOf(sn)),

		gg.ShortVarOf(
			ssName,
			gg.CallOf(
				gg.SelectOf("spec", "Struct"),
				gg.CallOf(
					gg.SelectOf("reflect", "TypeOf"),
					zn))),

		gg.AssignOf("_", ssName),
	)

	var sfs []gg.StructField
	var dflVds []gg.Var

	initStmts := gg.Stmts{
		gg.ExprStmtOf(gg.CallOf("_def_init")),
	}

	rcvr := ss.Receiver()
	if rcvr == "" {
		rcvr = strings.ToLower(string([]rune(strings.TrimLeft(sn, "_"))[0]))
	}

	for _, fs := range ss.Fields() {
		sfs = append(sfs, gg.StructField{Name: gg.IdentOf(fs.Name()), Type: fg.ti.importedType(fs.Type())})

		fsName := gg.IdentOf(fmt.Sprintf("field_spec__%s__%s", sn, fs.Name()))

		fg.initStmts.Append(
			gg.Blank{},

			gg.ShortVarOf(
				fsName,
				gg.CallOf(
					gg.SelectOf(ssName, "Field"),
					gg.LitOf(fmt.Sprintf("\"%s\"", fs.Name())))),

			gg.AssignOf("_", fsName),
		)

		if fs.Default() != nil {
			dflName := gg.IdentOf(fmt.Sprintf("_def_field_default__%s__%s", sn, fs.Name()))

			dflVds = append(dflVds,
				gg.VarOf(dflName, fg.ti.importedType(fs.Type())))

			fg.initStmts.Append(
				gg.Blank{},

				gg.AssignOf(
					dflName,
					gg.TypeAssertOf(
						gg.CallOf(gg.SelectOf(fsName, "Default")),
						fg.ti.importedType(fs.Type()))))

			initStmts.Append(
				gg.Blank{},

				gg.AssignOf(gg.SelectOf("f", fs.Name()), dflName))
		}
	}

	fg.decls.Append(gg.Struct{Name: sName, Fields: sfs})

	if len(dflVds) > 0 {
		fg.decls.Append(gg.VarsOf(dflVds...))
	}

	if len(ss.Inits()) > 0 {
		fg.initStmts.Append(gg.Blank{})
		initStmts.Append(gg.Blank{})

		var initVds []gg.Var
		for i, o := range ss.Inits() {
			var initName gg.Ident
			switch o := o.(type) {
			case *ast.FuncLit:
				initName = gg.IdentOf(fmt.Sprintf("_def_struct_init__%s__%d", sn, i))

				initVds = append(initVds,
					gg.VarOf(initName, newPtrFuncType(sName)))

				fg.initStmts.Append(
					gg.AssignOf(
						initName,
						gg.TypeAssertOf(
							gg.IndexOf(gg.CallOf(gg.SelectOf(ssName, "Inits")), i),
							newPtrFuncType(sName))))

				initStmts.Append(gg.CallOf(initName, "f"))

			case *ast.FuncDecl:
				initStmts.Append(gg.CallOf(gg.SelectOf("f", o.Name.Name)))

			default:
				panic(i)
			}
		}

		fg.decls.Append(gg.VarsOf(initVds...))
	}

	fg.decls.Append(
		gg.Func{
			Receiver: &gg.Param{Name: gg.NewIdent("f"), Type: gg.PtrOf(sName)},
			Name:     gg.NewIdent("init"),
			Body:     &gg.Block{Body: initStmts}})
}

func (fg *FileGen) genStructs() {
	i := 0
	for _, ss := range fg.ps.Structs() {
		if i > 0 {
			fg.initStmts.Append(gg.Blank{})
		}
		i++
		fg.genStruct(ss)
	}
}
