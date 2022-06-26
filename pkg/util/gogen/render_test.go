package gogen

import (
	"fmt"
	"testing"

	opt "github.com/wrmsr/bane/pkg/util/optional"
)

func TestBuilder(t *testing.T) {
	n := NewFunc(
		opt.None[Param](),
		opt.Just(NewIdent("foo")),
		[]Param{
			NewParam(
				opt.Just(NewIdent("x")),
				NewNameType(NewIdent("int"))),
		},
		opt.None[Type](),
		opt.Just(NewBlock(
			NewIf(
				NewLit("foo"),
				NewBlock(
					NewExprStmt(NewLit("bar")))))))

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

	/*
		var (
			_def_init_once sync.Once

			_def_field_default__Foo__baz int

			_def_struct_inits__Foo__0 func(*Foo)
		)
	*/

	newPtrFuncType := func(elem Type) FuncType {
		return NewFuncType(
			NewFunc(
				opt.None[Param](),
				opt.None[Ident](),
				[]Param{
					NewParam(opt.None[Ident](), NewPtr(elem)),
				},
				opt.None[Type](),
				opt.None[Block](),
			),
		)
	}

	varsNode := NewVars([]Var{
		NewVar(NewIdent("_def_init_once"), NewNameType(NewIdent("sync.Once"))),
		NewVar(NewIdent("_def_field_default__Foo__baz"), NewNameType(NewIdent("int"))),
		NewVar(NewIdent("_def_struct_init__Foo__0"), newPtrFuncType(NewNameType(NewIdent("Foo")))),
	})
	fmt.Println(RenderString(varsNode))

	/*
		func _def_init() {
			_def_init_once.Do(func() {
				spec := def.X_getPackageSpec()

				struct_spec__Foo := spec.Struct("Foo")
				_ = struct_spec__Foo

				_def_field_default__Foo__baz = field_spec__Foo__baz.Default().(int)

				_def_struct_inits__Foo[0] = struct_spec__Foo.Inits()[0].(func(*Foo))
			})
		}
	*/

	//defInitOnceNode := NewFunc

	defInitNode := NewFunc(
		opt.None[Param](),
		opt.Just(NewIdent("_def_init")),
		nil,
		opt.None[Type](),
		opt.Just(NewBlock(
			NewExprStmt(NewCall(
				NewField(NewIdent("_def_init_once"), NewIdent("Do")),
				NewFuncExpr(NewFunc(
					opt.None[Param](),
					opt.None[Ident](),
					nil,
					opt.None[Type](),
					opt.Just(NewBlock(
						NewShortVar(
							NewIdent("spec"),
							NewCall(NewField(NewIdent("def"), NewIdent("X_getPackageSpec")))),

						NewShortVar(
							NewIdent("struct_spec__Foo"),
							NewCall(NewField(NewIdent("spec"), NewIdent("Struct")), NewLit("\"foo\""))),
						NewAssign(NewIdent("_"), NewIdent("struct_spec__Foo")),

						NewAssign(
							NewIdent("_def_field_default__Foo__baz"),
							NewTypeAssert(
								NewCall(NewField(NewIdent("struct_spec__Foo"), NewIdent("Default"))),
								NewNameType(NewIdent("int")))),

						NewAssign(
							NewIdent("_def_struct_init__Foo__0"),
							NewTypeAssert(
								NewIndex(NewCall(NewField(NewIdent("struct_spec__Foo"), NewIdent("Inits"))), NewLit("0")),
								newPtrFuncType(NewNameType(NewIdent("Foo"))))),
					)),
				)),
			)),
		)),
	)
	fmt.Println(RenderString(defInitNode))

	/*
		type Foo struct {
			bar int

			baz int
		}
	*/

	structNode := NewStruct(
		NewIdent("Foo"),
		NewStructField(NewIdent("bar"), NewNameType(NewIdent("int"))),
		NewStructField(NewIdent("baz"), NewNameType(NewIdent("int"))),
	)
	fmt.Println(RenderString(structNode))

	/*
		func (f *Foo) init() {
			_def_init()

			f.baz = _def_field_default__Foo__baz

			_def_struct_inits__Foo[0](f)
		}
	*/

	fooInitNode := NewFunc(
		opt.Just(NewParam(opt.Just(NewIdent("f")), NewPtr(NewNameType(NewIdent("Foo"))))),
		opt.Just(NewIdent("init")),
		nil,
		opt.None[Type](),
		opt.Just(NewBlock(
			NewExprStmt(NewCall(NewIdent("_def_init"))),
			NewAssign(NewField(NewIdent("f"), NewIdent("baz")), NewIdent("_def_field_default__Foo__baz")),
			NewExprStmt(NewCall(NewIdent("_def_struct_init__Foo__0"), NewIdent("f"))),
		)),
	)
	fmt.Println(RenderString(fooInitNode))
}
