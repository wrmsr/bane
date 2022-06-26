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

			_def_struct_inits__Foo [1]func(*Foo)
		)
	*/

	varsNode := NewVars([]Var{
		NewVar(NewIdent("_def_init_once"), NewNameType(NewIdent("sync.Once"))),
		NewVar(NewIdent("_def_field_default__Foo__baz"), NewNameType(NewIdent("int"))),
		NewVar(NewIdent("_def_struct_inits__Foo"), NewArray(opt.Just(1), NewFuncType(
			NewFunc(
				opt.None[Param](),
				opt.None[Ident](),
				[]Param{
					NewParam(opt.None[Ident](), NewPtr(NewNameType(NewIdent("Foo")))),
				},
				opt.None[Type](),
				opt.None[Block](),
			),
		))),
	})
	fmt.Println(RenderString(varsNode))

	/*
		func _def_init() {
			_def_init_once.Do(func() {
				spec := def.X_getPackageSpec()

				struct_spec__Foo := spec.Struct("Foo")
				_ = struct_spec__Foo

				field_spec__Foo__bar := struct_spec__Foo.Field("bar")
				_ = field_spec__Foo__bar

				field_spec__Foo__baz := struct_spec__Foo.Field("baz")
				_ = field_spec__Foo__baz

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
						NewShortVar(NewIdent("spec"), NewField(NewIdent("def"), NewIdent("X_getPackageSpec"))),
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

	/*
		func (f *Foo) init() {
			_def_init()

			f.baz = _def_field_default__Foo__baz

			_def_struct_inits__Foo[0](f)
		}
	*/
}
