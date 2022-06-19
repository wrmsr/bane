//go:build !nodev

package test

import (
	"fmt"

	stu "github.com/wrmsr/bane/pkg/utils/structs"
)

//go:generate go run github.com/wrmsr/bane/pkg/utils/structs/dev/gen

var _ = stu.Def(
	stu.Struct("Foo",
		stu.Field("bar", stu.Type[int]()),
		stu.Field("baz", stu.Type[int]()),
	),
)

func foof(s string, a ...any) {
	fmt.Printf(s, a...)
}

func Foo() {
	foof("hi %s\n", "there")
}
