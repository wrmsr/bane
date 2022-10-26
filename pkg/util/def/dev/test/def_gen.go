package test

import (
	"reflect"
	"sync"

	"github.com/wrmsr/bane/pkg/util/def"
	"github.com/wrmsr/bane/pkg/util/types"
)

var _def_init_once sync.Once

func _def_init() {
	_def_init_once.Do(func() {
		spec := def.X_getPackageSpec()

		var zero_Foo Foo
		struct_spec__Foo := spec.Struct(reflect.TypeOf(zero_Foo))
		_ = struct_spec__Foo

		field_spec__Foo__bar := struct_spec__Foo.Field("bar")
		_ = field_spec__Foo__bar

		field_spec__Foo__baz := struct_spec__Foo.Field("baz")
		_ = field_spec__Foo__baz

		_def_field_default__Foo__baz = field_spec__Foo__baz.Default().(int)

		field_spec__Foo__optInt := struct_spec__Foo.Field("optInt")
		_ = field_spec__Foo__optInt

		field_spec__Foo__optOptInt := struct_spec__Foo.Field("optOptInt")
		_ = field_spec__Foo__optOptInt

		field_spec__Foo__dflInt := struct_spec__Foo.Field("dflInt")
		_ = field_spec__Foo__dflInt

		_def_field_default__Foo__dflInt = field_spec__Foo__dflInt.Default().(types.Optional[int])

		field_spec__Foo__thing := struct_spec__Foo.Field("thing")
		_ = field_spec__Foo__thing

		field_spec__Foo__aThing := struct_spec__Foo.Field("aThing")
		_ = field_spec__Foo__aThing

		_def_field_default__Foo__aThing = field_spec__Foo__aThing.Default().(Thing)

		_def_struct_init__Foo__0 = struct_spec__Foo.Inits()[0].(func(*Foo))
	})
}

type Foo struct {
	bar       int
	baz       int
	optInt    types.Optional[int]
	optOptInt types.Optional[types.Optional[int]]
	dflInt    types.Optional[int]
	thing     Thing
	aThing    Thing
}

var (
	_def_field_default__Foo__baz    int
	_def_field_default__Foo__dflInt types.Optional[int]
	_def_field_default__Foo__aThing Thing
)

var (
	_def_struct_init__Foo__0 func(*Foo)
)

func (f *Foo) init() {
	_def_init()

	f.baz = _def_field_default__Foo__baz

	f.dflInt = _def_field_default__Foo__dflInt

	f.aThing = _def_field_default__Foo__aThing

	_def_struct_init__Foo__0(f)
	f._def_init_barf()
}

func _def_inl_Baz(t InlineThing, y int) int {
	var __def_inl_0 int
	__def_inl_1 := y
	{
		{
			__def_inl_0 = t.x + __def_inl_1
			goto __def_inl_2
		}
	__def_inl_2:
	}

	return __def_inl_0 + 1
}
func _def_inl_Bar(x, y int) int {
	var __def_inl_0 int
	__def_inl_1 := x
	__def_inl_2 := y
	{
		{
			__def_inl_0 = __def_inl_1 + __def_inl_2
			goto __def_inl_3
		}
	__def_inl_3:
	}
	var __def_inl_4 int
	__def_inl_5 := __def_inl_0
	__def_inl_6 := y
	{
		{
			__def_inl_4 = __def_inl_5 * __def_inl_6
			goto __def_inl_7
		}
	__def_inl_7:
	}

	return __def_inl_4
}
