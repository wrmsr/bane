//go:build !nodev

package main

import (
	"fmt"

	"github.com/wrmsr/bane/pkg/util/def"
	opt "github.com/wrmsr/bane/pkg/util/optional"
)

//go:generate go run github.com/wrmsr/bane/pkg/util/def/dev/gen

var _ = def.Struct("Foo",
	def.Field("bar", def.Type[int]()),
	def.Field("baz", def.Default(5)),
	def.Field("optInt", def.Type[opt.Optional[int]]()),
	def.Field("optOptInt", def.Type[opt.Optional[opt.Optional[int]]]()),
	def.Field("dflInt", def.Default(opt.Just(10))),
	def.Init(func(f *Foo) {
		fmt.Printf("new foo: %d %d\n", f.baz, f.bar)
	}),
)

func foof(s string, a ...any) {
	fmt.Printf(s, a...)
}

//

func main() {
	foof("hi %s\n", "there")

	var f Foo
	f.init()
	fmt.Printf("%+v\n", f)
}
