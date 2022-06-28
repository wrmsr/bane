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
						Ty:         def.Type[int]().Ty.(reflect.Type),
						HasDefault: true,
					},
				},
				NumInits: 1,
			},
		},
	})
	return nil
}()

var _def_init_once sync.Once

func _def_init() {
	_def_init_once.Do(func() {
		spec := def.X_getPackageSpec()

		struct_spec__Foo := spec.Struct("Foo")
		_ = struct_spec__Foo

		field_spec__Foo__baz := struct_spec__Foo.Field("Foo")
		_ = struct_spec__Foo

		_def_field_default__Foo__baz = field_spec__Foo__baz.Default().(int)

		_def_struct_init__Foo__0 = struct_spec__Foo.Inits()[0].(func(*Foo))
	})
}

type Foo struct {
	bar int
	baz int
}

var (
	_def_field_default__Foo__baz int
	_def_struct_init__Foo__0     func(*Foo)
)

func (f *Foo) init() {
	_def_init()

	f.baz = _def_field_default__Foo__baz

	_def_struct_init__Foo__0(f)
}
