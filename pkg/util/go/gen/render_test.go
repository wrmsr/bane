package gen

import (
	"fmt"
	"testing"

	opt "github.com/wrmsr/bane/pkg/util/optional"
)

func TestBuilder(t *testing.T) {
	n := Func{
		Name: IdentOf("foo"),
		Params: []Param{
			NewParam(
				opt.Just(NewIdent("x")),
				NewNameType(NewIdent("int"))),
		},
		Body: BlockOf(
			NewIf(
				NewLit("foo"),
				NewBlock(
					NewExprStmt(NewLit("bar")))))}

	fmt.Println(RenderString(n))
}

func TestDefBuilder(t *testing.T) {
	/*
		var _ = func() any {
			def.X_addPackageDef(def.X_packageExpect{
				Structs: map[string]def.X_structExpect{
					"Foo": {
						Fields: map[string]def.X_fieldExpect{
							"bar": {
								Ty: def.Type[int]().Ty.(reflect.Type),
							},
							"baz": {
								Ty: def.Type[int]().Ty.(reflect.Type),
								HasDefault: true,
							},
						},
						NumInits: 1,
					},
				},
			})
			return nil
		}()
	*/
	expectsNode := NewVar(NewIdent("_"), opt.None[Type](), opt.Just[Expr](NewCall(NewFuncExpr(Func{
		Type: NewNameType(NewIdent("any")),
		Body: BlockOf(),
	}))))
	fmt.Println(RenderString(expectsNode))

	newPtrFuncType := func(elem Type) FuncType {
		return NewFuncType(
			Func{
				Params: []Param{
					NewParam(opt.None[Ident](), NewPtr(elem)),
				},
			},
		)
	}

	varsNode := NewVars([]Var{
		NewVar(NewIdent("_def_init_once"), opt.Just[Type](NewNameType(NewIdent("sync.Once"))), opt.None[Expr]()),
		NewVar(NewIdent("_def_field_default__Foo__baz"), opt.Just[Type](NewNameType(NewIdent("int"))), opt.None[Expr]()),
		NewVar(NewIdent("_def_struct_init__Foo__0"), opt.Just[Type](newPtrFuncType(NewNameType(NewIdent("Foo")))), opt.None[Expr]()),
	})
	fmt.Println(RenderString(varsNode))

	defInitNode := Func{
		Name: IdentOf("_def_init"),
		Body: BlockOf(
			NewExprStmt(NewCall(
				NewSelect(NewIdent("_def_init_once"), NewIdent("Do")),
				NewFuncExpr(Func{
					Body: BlockOf(
						NewShortVar(
							NewIdent("spec"),
							NewCall(NewSelect(NewIdent("def"), NewIdent("X_getPackageSpec")))),

						NewShortVar(
							NewIdent("struct_spec__Foo"),
							NewCall(NewSelect(NewIdent("spec"), NewIdent("Struct")), NewLit("\"foo\""))),
						NewAssign(NewIdent("_"), NewIdent("struct_spec__Foo")),

						NewAssign(
							NewIdent("_def_field_default__Foo__baz"),
							NewTypeAssert(
								NewCall(NewSelect(NewIdent("struct_spec__Foo"), NewIdent("Default"))),
								NewNameType(NewIdent("int")))),

						NewAssign(
							NewIdent("_def_struct_init__Foo__0"),
							NewTypeAssert(
								NewIndex(NewCall(NewSelect(NewIdent("struct_spec__Foo"), NewIdent("Inits"))), NewLit("0")),
								newPtrFuncType(NewNameType(NewIdent("Foo"))))),
					),
				}),
			)),
		),
	}
	fmt.Println(RenderString(defInitNode))

	structNode := NewStruct(
		NewIdent("Foo"),
		NewStructField(NewIdent("bar"), NewNameType(NewIdent("int"))),
		NewStructField(NewIdent("baz"), NewNameType(NewIdent("int"))),
	)
	fmt.Println(RenderString(structNode))

	fooInitNode := Func{
		Receiver: ParamOf(opt.Just(NewIdent("f")), NewPtr(NewNameType(NewIdent("Foo")))),
		Name:     IdentOf("init"),
		Body: BlockOf(
			NewExprStmt(NewCall(NewIdent("_def_init"))),
			NewAssign(NewSelect(NewIdent("f"), NewIdent("baz")), NewIdent("_def_field_default__Foo__baz")),
			NewExprStmt(NewCall(NewIdent("_def_struct_init__Foo__0"), NewIdent("f"))),
		),
	}
	fmt.Println(RenderString(fooInitNode))
}
