//go:build !nodev

package main

import (
	"fmt"

	"github.com/wrmsr/bane/pkg/utils/structs/def"
)

//go:generate go run github.com/wrmsr/bane/pkg/utils/structs/dev/gen

var _ = def.Def(
	def.Struct("Foo",
		def.Field("bar", def.Type[int]()),
		def.Field("baz", def.Default(5)),
	),
)

func foof(s string, a ...any) {
	fmt.Printf(s, a...)
}

func main() {
	foof("hi %s\n", "there")
}
