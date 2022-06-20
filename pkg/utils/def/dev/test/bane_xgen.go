//go:build !nodev

package main

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

func init() {

}
