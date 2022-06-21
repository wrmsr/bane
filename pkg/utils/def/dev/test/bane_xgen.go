//go:build !nodev

package main

import (
	"sync"

	"github.com/wrmsr/bane/pkg/utils/def"
)

var _ = def.X_structExpectsDef{
	Name: "Foo",
	Fields: []string{
		"bar",
		"baz",
	},
	Defaults: []string{
		"baz",
	},
	Inits: 1,
}.Register()

var (
	_bane_init_once sync.Once

	_bane_init__Foo__default_baz int
	_bane_init__Foo__inits       []func(*Foo)
)

func _bane_init() {
	_bane_init_once.Do(func() {
		init := def.X_getPackageInit()

		init_Foo := init.Structs["Foo"]
		_bane_init__Foo__default_baz = init_Foo.Fields["baz"].Default.(int)
		_bane_init__Foo__inits = make([]func(*Foo), len(init_Foo.Inits))
		if len(_bane_init__Foo__inits) != 1 {
			panic("Foo")
		}
		for i, f := range init_Foo.Inits {
			var ok bool
			if _bane_init__Foo__inits[i], ok = f.(func(*Foo)); !ok {
				panic("Foo")
			}
		}
	})
}

type Foo struct {
	bar, baz int
}

func (f *Foo) init() {
	_bane_init()
	f.baz = _bane_init__Foo__default_baz
	for _, fn := range _bane_init__Foo__inits {
		fn(f)
	}
}
