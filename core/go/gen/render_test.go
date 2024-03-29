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
				Type: TypeOf("int")},
		},
		Body: NewBlock(
			IfOf(
				LitOf("foo"),
				*NewBlock(
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
	expectsNode := VarOf(
		IdentOf("_"),
		CallOf(FuncExpr{Func: Func{
			Type: TypeOf("any"),
			Body: NewBlock()}}))

	fmt.Println(RenderString(expectsNode))

	newPtrFuncType := func(elem any) FuncType {
		return FuncTypeOf(
			Func{
				Params: []Param{
					{Type: PtrOf(elem)},
				},
			},
		)
	}

	varsNode := VarsOf(
		VarOf("_def_init_once", TypeOf("sync.Once")),
		VarOf("_def_field_default__Foo__baz", TypeOf("int")),
		VarOf("_def_struct_init__Foo__0", newPtrFuncType("Foo")),
	)
	fmt.Println(RenderString(varsNode))

	defInitNode := Func{
		Name: NewIdent("_def_init"),
		Body: NewBlock(
			CallOf(
				SelectOf("_def_init_once", "Do"),
				FuncExprOf(Func{
					Body: NewBlock(
						ShortVarOf(
							"spec",
							CallOf(SelectOf("def", "X_getPackageSpec"))),

						ShortVarOf(
							"struct_spec__Foo",
							CallOf(SelectOf("spec", "Struct"), LitOf("\"foo\""))),
						AssignOf("_", "struct_spec__Foo"),

						AssignOf(
							"_def_field_default__Foo__baz",
							TypeAssertOf(
								CallOf(SelectOf("struct_spec__Foo", "Default")),
								"int")),

						AssignOf(
							"_def_struct_init__Foo__0",
							TypeAssertOf(
								IndexOf(CallOf(SelectOf("struct_spec__Foo", "Inits")), 0),
								newPtrFuncType("Foo"))),
					),
				}),
			),
		),
	}
	fmt.Println(RenderString(defInitNode))

	structNode := Struct{
		Name: IdentOf("Foo"),
		Fields: []StructField{
			{Name: IdentOf("bar"), Type: TypeOf("int")},
			{Name: IdentOf("baz"), Type: TypeOf("int")},
		},
	}
	fmt.Println(RenderString(structNode))

	fooInitNode := Func{
		Receiver: &Param{Name: NewIdent("f"), Type: PtrOf("Foo")},
		Name:     NewIdent("init"),
		Body: NewBlock(
			CallOf("_def_init"),
			AssignOf(SelectOf("f", "baz"), "_def_field_default__Foo__baz"),
			CallOf("_def_struct_init__Foo__0", "f"),
		),
	}
	fmt.Println(RenderString(fooInitNode))
}
