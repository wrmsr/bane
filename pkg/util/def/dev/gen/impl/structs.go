//go:build !nodev

package impl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wrmsr/bane/pkg/util/def"
	opt "github.com/wrmsr/bane/pkg/util/optional"
)

func (fg *FileGen) genStruct(ss *def.StructSpec) {
	sName := gen.NewIdent(ss.Name())
	ssName := gen.NewIdent(fmt.Sprintf("struct_spec__%s", ss.Name()))

	fg.initStmts = append(fg.initStmts,
		gen.NewBlank(),

		gen.NewShortVar(ssName,
			gen.NewCall(
				gen.NewSelect(gen.NewIdent("spec"), gen.NewIdent("Struct")),
				gen.NewLit(fmt.Sprintf("\"%s\"", ss.Name())))),

		gen.NewAssign(gen.NewIdent("_"), ssName),
	)

	var sfs []gen.StructField
	var dflVds []gen.Var

	initStmts := []gen.Stmt{
		gen.NewExprStmt(gen.NewCall(gen.NewIdent("_def_init"))),
	}

	rcvr := ss.Receiver()
	if rcvr == "" {
		rcvr = strings.ToLower(string([]rune(strings.TrimLeft(ss.Name(), "_"))[0]))
	}

	for _, fs := range ss.Fields() {
		sfs = append(sfs, gen.NewStructField(gen.NewIdent(fs.Name()), fg.ti.importedType(fs.Type())))

		fsName := gen.NewIdent(fmt.Sprintf("field_spec__%s__%s", ss.Name(), fs.Name()))

		fg.initStmts = append(fg.initStmts,
			gen.NewBlank(),

			gen.NewShortVar(fsName,
				gen.NewCall(
					gen.NewSelect(ssName, gen.NewIdent("Field")),
					gen.NewLit(fmt.Sprintf("\"%s\"", fs.Name())))),

			gen.NewAssign(gen.NewIdent("_"), fsName),
		)

		if fs.Default() != nil {
			dflName := gen.NewIdent(fmt.Sprintf("_def_field_default__%s__%s", ss.Name(), fs.Name()))

			dflVds = append(dflVds,
				gen.NewVar(dflName, opt.Just[gen.Type](fg.ti.importedType(fs.Type())), opt.None[gen.Expr]()))

			fg.initStmts = append(fg.initStmts,
				gen.NewBlank(),

				gen.NewAssign(dflName,
					gen.NewTypeAssert(
						gen.NewCall(gen.NewSelect(fsName, gen.NewIdent("Default"))),
						fg.ti.importedType(fs.Type()))))

			initStmts = append(initStmts,
				gen.NewBlank(),

				gen.NewAssign(gen.NewSelect(gen.NewIdent("f"), gen.NewIdent(fs.Name())), dflName))
		}
	}

	fg.decls = append(fg.decls, gen.NewStruct(sName, sfs...))

	if len(dflVds) > 0 {
		fg.decls = append(fg.decls,
			gen.NewStmtDecl(gen.NewVars(dflVds)))
	}

	if len(ss.Inits()) > 0 {
		fg.initStmts = append(fg.initStmts, gen.NewBlank())
		initStmts = append(initStmts, gen.NewBlank())

		var initVds []gen.Var
		for i := range ss.Inits() {
			initName := gen.NewIdent(fmt.Sprintf("_def_struct_init__%s__%d", ss.Name(), i))

			initVds = append(initVds,
				gen.NewVar(initName,
					opt.Just[gen.Type](newPtrFuncType(gen.NewNameType(sName))), opt.None[gen.Expr]()))

			fg.initStmts = append(fg.initStmts,
				gen.NewAssign(initName,
					gen.NewTypeAssert(
						gen.NewIndex(gen.NewCall(gen.NewSelect(ssName, gen.NewIdent("Inits"))), gen.NewLit(strconv.Itoa(i))),
						newPtrFuncType(gen.NewNameType(sName)))))

			initStmts = append(initStmts,
				gen.NewExprStmt(gen.NewCall(initName, gen.NewIdent("f"))))
		}

		fg.decls = append(fg.decls,
			gen.NewStmtDecl(gen.NewVars(initVds)))
	}

	fg.decls = append(fg.decls,
		gen.NewFunc(
			opt.Just(gen.NewParam(opt.Just(gen.NewIdent("f")), gen.NewPtr(gen.NewNameType(sName)))),
			opt.Just(gen.NewIdent("init")),
			nil,
			opt.None[gen.Type](),
			opt.Just(gen.NewBlock(initStmts...))))
}

func (fg *FileGen) genStructs() {
	i := 0
	for _, ss := range fg.ps.Structs() {
		if i > 0 {
			fg.initStmts = append(fg.initStmts, gen.NewBlank())
		}
		i++
		fg.genStruct(ss)
	}
}
