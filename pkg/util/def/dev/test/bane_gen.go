//go:build !nodev

package main

import (
	"reflect"
	"sync"

	"github.com/wrmsr/bane/pkg/util/def"
)

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

var (
	_bane_def_init_once sync.Once

	_bane_def_field_default__Foo__baz int

	_bane_def_struct_inits__Foo [1]func(*Foo)
)

func _bane_def_init() {
	_bane_def_init_once.Do(func() {
		spec := def.X_getPackageSpec()

		struct_spec__Foo := spec.Struct("Foo")
		_ = struct_spec__Foo

		field_spec__Foo__bar := struct_spec__Foo.Field("bar")
		_ = field_spec__Foo__bar

		field_spec__Foo__baz := struct_spec__Foo.Field("baz")
		_ = field_spec__Foo__baz

		_bane_def_field_default__Foo__baz = field_spec__Foo__baz.Default().(int)

		_bane_def_struct_inits__Foo[0] = struct_spec__Foo.Inits()[0].(func(*Foo))

	})
}

type Foo struct {
	bar int

	baz int
}

func (f *Foo) init() {
	_bane_def_init()

	f.baz = _bane_def_field_default__Foo__baz

	_bane_def_struct_inits__Foo[0](f)

}
