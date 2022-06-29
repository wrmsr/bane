//go:build !nodev

package test
import (
	"github.com/wrmsr/bane/pkg/util/optional"
	"sync"
	"github.com/wrmsr/bane/pkg/util/def"
)

var _def_init_once sync.Once

func _def_init() {
	_def_init_once.Do(func() {
		spec := def.X_getPackageSpec()
		
		struct_spec__Foo := spec.Struct("Foo")
		_ = struct_spec__Foo
		
		field_spec__Foo__dflInt := struct_spec__Foo.Field("dflInt")
		_ = field_spec__Foo__dflInt
		
		_def_field_default__Foo__dflInt = field_spec__Foo__dflInt.Default().(optional.Optional[int])
		
		field_spec__Foo__thing := struct_spec__Foo.Field("thing")
		_ = field_spec__Foo__thing
		
		field_spec__Foo__aThing := struct_spec__Foo.Field("aThing")
		_ = field_spec__Foo__aThing
		
		_def_field_default__Foo__aThing = field_spec__Foo__aThing.Default().(Thing)
		
		field_spec__Foo__bar := struct_spec__Foo.Field("bar")
		_ = field_spec__Foo__bar
		
		field_spec__Foo__baz := struct_spec__Foo.Field("baz")
		_ = field_spec__Foo__baz
		
		_def_field_default__Foo__baz = field_spec__Foo__baz.Default().(int)
		
		field_spec__Foo__optInt := struct_spec__Foo.Field("optInt")
		_ = field_spec__Foo__optInt
		
		field_spec__Foo__optOptInt := struct_spec__Foo.Field("optOptInt")
		_ = field_spec__Foo__optOptInt
		
		_def_struct_init__Foo__0 = struct_spec__Foo.Inits()[0].(func(*Foo))
	})
}

type Foo struct {
	dflInt optional.Optional[int]
	thing Thing
	aThing Thing
	bar int
	baz int
	optInt optional.Optional[int]
	optOptInt optional.Optional[optional.Optional[int]]
}

var (
	_def_field_default__Foo__dflInt optional.Optional[int]
	_def_field_default__Foo__aThing Thing
	_def_field_default__Foo__baz int
)

var (
	_def_struct_init__Foo__0 func(*Foo)
)

func (f *Foo) init() {
	_def_init()
	
	f.dflInt = _def_field_default__Foo__dflInt
	
	f.aThing = _def_field_default__Foo__aThing
	
	f.baz = _def_field_default__Foo__baz
	
	_def_struct_init__Foo__0(f)
}
