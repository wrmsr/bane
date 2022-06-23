//go:build !nodev

package main

import (
	"reflect"
	"sync"

	"github.com/wrmsr/bane/pkg/util/def"
)

var _ = func() any {
	def.X_addPackageDef(
		def.X_packageExpect{
			Structs: map[string]def.X_structExpect{
				"Foo": {
					Fields: map[string]def.X_fieldExpect{
						"bar": {
							Ty: def.Type[int]().Ty.(reflect.Type),
						},
						"baz": {
							HasDefault: true,
						},
					},
					NumInits: 1,
				},
			},
		},
	)
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

		field_spec__Foo__baz := struct_spec__Foo.Field("baz")
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
	for _, fn := range _bane_def_struct_inits__Foo {
		fn(f)
	}
}
