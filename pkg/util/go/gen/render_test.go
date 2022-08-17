package gen

import (
	"fmt"
	"testing"
)

func TestBuilder(t *testing.T) {
	n := Func{
		Name: NewIdent("foo"),
		Params: []Param{
			{
				Name: NewIdent("x"),
				Type: NameTypeOf(IdentOf("int"))},
		},
		Body: NewBlock(
			IfOf(
				LitOf("foo"),
				BlockOf(
					ExprStmtOf(LitOf("bar")))))}

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
	expectsNode := Var{
		Name: IdentOf("_"),
		Value: CallOf(FuncExpr{Func: Func{
			Type: NameTypeOf(IdentOf("any")),
			Body: NewBlock()}})}

	fmt.Println(RenderString(expectsNode))

	newPtrFuncType := func(elem Type) FuncType {
		return FuncTypeOf(
			Func{
				Params: []Param{
					{Type: PtrOf(elem)},
				},
			},
		)
	}

	varsNode := VarsOf(
		Var{Name: IdentOf("_def_init_once"), Type: NameTypeOf(IdentOf("sync.Once"))},
		Var{Name: IdentOf("_def_field_default__Foo__baz"), Type: NameTypeOf(IdentOf("int"))},
		Var{Name: IdentOf("_def_struct_init__Foo__0"), Type: newPtrFuncType(NameTypeOf(IdentOf("Foo")))},
	)
	fmt.Println(RenderString(varsNode))

	defInitNode := Func{
		Name: NewIdent("_def_init"),
		Body: NewBlock(
			ExprStmtOf(CallOf(
				SelectOf(IdentOf("_def_init_once"), IdentOf("Do")),
				FuncExprOf(Func{
					Body: NewBlock(
						ShortVarOf(
							IdentOf("spec"),
							CallOf(SelectOf(IdentOf("def"), IdentOf("X_getPackageSpec")))),

						ShortVarOf(
							IdentOf("struct_spec__Foo"),
							CallOf(SelectOf(IdentOf("spec"), IdentOf("Struct")), LitOf("\"foo\""))),
						AssignOf(IdentOf("_"), IdentOf("struct_spec__Foo")),

						AssignOf(
							IdentOf("_def_field_default__Foo__baz"),
							TypeAssertOf(
								CallOf(SelectOf(IdentOf("struct_spec__Foo"), IdentOf("Default"))),
								NameTypeOf(IdentOf("int")))),

						AssignOf(
							IdentOf("_def_struct_init__Foo__0"),
							TypeAssertOf(
								IndexOf(CallOf(SelectOf(IdentOf("struct_spec__Foo"), IdentOf("Inits"))), LitOf("0")),
								newPtrFuncType(NameTypeOf(IdentOf("Foo"))))),
					),
				}),
			)),
		),
	}
	fmt.Println(RenderString(defInitNode))

	structNode := Struct{
		Name: IdentOf("Foo"),
		Fields: []StructField{
			{Name: IdentOf("bar"), Type: NameTypeOf(IdentOf("int"))},
			{Name: IdentOf("baz"), Type: NameTypeOf(IdentOf("int"))},
		},
	}
	fmt.Println(RenderString(structNode))

	fooInitNode := Func{
		Receiver: &Param{Name: NewIdent("f"), Type: PtrOf(NameTypeOf(IdentOf("Foo")))},
		Name:     NewIdent("init"),
		Body: NewBlock(
			ExprStmtOf(CallOf(IdentOf("_def_init"))),
			AssignOf(SelectOf(IdentOf("f"), IdentOf("baz")), IdentOf("_def_field_default__Foo__baz")),
			ExprStmtOf(CallOf(IdentOf("_def_struct_init__Foo__0"), IdentOf("f"))),
		),
	}
	fmt.Println(RenderString(fooInitNode))
}
