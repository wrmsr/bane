package gen

import (
	"fmt"
	"testing"

	opt "github.com/wrmsr/bane/pkg/util/optional"
)

func TestBuilder(t *testing.T) {
	n := Func{
		Name: NewIdent("foo"),
		Params: []Param{
			{
				Name: NewIdent("x"),
				Type: NewNameType(IdentOf("int"))},
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
	expectsNode := NewVar(IdentOf("_"), opt.None[Type](), opt.Just[Expr](NewCall(NewFuncExpr(Func{
		Type: NewNameType(IdentOf("any")),
		Body: BlockOf(),
	}))))
	fmt.Println(RenderString(expectsNode))

	newPtrFuncType := func(elem Type) FuncType {
		return NewFuncType(
			Func{
				Params: []Param{
					{Type: NewPtr(elem)},
				},
			},
		)
	}

	varsNode := NewVars([]Var{
		NewVar(IdentOf("_def_init_once"), opt.Just[Type](NewNameType(IdentOf("sync.Once"))), opt.None[Expr]()),
		NewVar(IdentOf("_def_field_default__Foo__baz"), opt.Just[Type](NewNameType(IdentOf("int"))), opt.None[Expr]()),
		NewVar(IdentOf("_def_struct_init__Foo__0"), opt.Just[Type](newPtrFuncType(NewNameType(IdentOf("Foo")))), opt.None[Expr]()),
	})
	fmt.Println(RenderString(varsNode))

	defInitNode := Func{
		Name: NewIdent("_def_init"),
		Body: BlockOf(
			NewExprStmt(NewCall(
				NewSelect(IdentOf("_def_init_once"), IdentOf("Do")),
				NewFuncExpr(Func{
					Body: BlockOf(
						NewShortVar(
							IdentOf("spec"),
							NewCall(NewSelect(IdentOf("def"), IdentOf("X_getPackageSpec")))),

						NewShortVar(
							IdentOf("struct_spec__Foo"),
							NewCall(NewSelect(IdentOf("spec"), IdentOf("Struct")), NewLit("\"foo\""))),
						NewAssign(IdentOf("_"), IdentOf("struct_spec__Foo")),

						NewAssign(
							IdentOf("_def_field_default__Foo__baz"),
							NewTypeAssert(
								NewCall(NewSelect(IdentOf("struct_spec__Foo"), IdentOf("Default"))),
								NewNameType(IdentOf("int")))),

						NewAssign(
							IdentOf("_def_struct_init__Foo__0"),
							NewTypeAssert(
								NewIndex(NewCall(NewSelect(IdentOf("struct_spec__Foo"), IdentOf("Inits"))), NewLit("0")),
								newPtrFuncType(NewNameType(IdentOf("Foo"))))),
					),
				}),
			)),
		),
	}
	fmt.Println(RenderString(defInitNode))

	structNode := Struct{
		Name: IdentOf("Foo"),
		Fields: []StructField{
			{Name: IdentOf("bar"), Type: NewNameType(IdentOf("int"))},
			{Name: IdentOf("baz"), Type: NewNameType(IdentOf("int"))},
		},
	}
	fmt.Println(RenderString(structNode))

	fooInitNode := Func{
		Receiver: &Param{Name: NewIdent("f"), Type: NewPtr(NewNameType(IdentOf("Foo")))},
		Name:     NewIdent("init"),
		Body: BlockOf(
			NewExprStmt(NewCall(IdentOf("_def_init"))),
			NewAssign(NewSelect(IdentOf("f"), IdentOf("baz")), IdentOf("_def_field_default__Foo__baz")),
			NewExprStmt(NewCall(IdentOf("_def_struct_init__Foo__0"), IdentOf("f"))),
		),
	}
	fmt.Println(RenderString(fooInitNode))
}
