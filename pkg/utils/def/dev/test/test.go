//go:build !nodev

package main

import (
	"fmt"

	"github.com/wrmsr/bane/pkg/utils/def"
)

//go:generate go run github.com/wrmsr/bane/pkg/utils/def/dev/gen

var _ = def.Struct("Foo",
	def.Field("bar", def.Type[int]()),
	def.Field("baz", def.Default(5)),
	def.Init(func(f *Foo) {
		fmt.Printf("new foo: %d %d\n", f.baz, f.bar)
	}),
)

func foof(s string, a ...any) {
	fmt.Printf(s, a...)
}

//
// TODO: codegen

//

func main() {
	foof("hi %s\n", "there")
}
