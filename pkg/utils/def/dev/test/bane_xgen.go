//go:build !nodev

package main

import (
	"github.com/wrmsr/bane/pkg/utils/def"
)

var _ = def.XStructExpectsDef{
	Name: "Foo",

	FieldNames: []string{
		"bar",
		"baz",
	},
	NumInits: 1,
}.Register()

type Foo struct {
	bar, baz int
}

func NewFoo(bar, baz int) Foo {
	return Foo{
		bar: bar,
		baz: baz,
	}
}

func (f *Foo) init() {
	f.baz = 5
}
