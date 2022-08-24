//go:build !nodev

package test

import (
	"fmt"

	"github.com/wrmsr/bane/pkg/util/def"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

type Thing struct {
	x int
}

var aThing = Thing{420}

var _ = def.Struct[Foo](
	def.Field("bar", def.Type[int]()),
	def.Field("baz", def.Default(5)),
	def.Field("optInt", def.Type[bt.Optional[int]]()),
	def.Field("optOptInt", def.Type[bt.Optional[bt.Optional[int]]]()),
	def.Field("dflInt", def.Default(bt.Just(10))),
	def.Field("thing", def.Type[Thing]()),
	def.Field("aThing", def.Default(aThing)),
	def.Init(func(f *Foo) {
		fmt.Printf("new foo: %d %d\n", f.baz, f.bar)
	}),
)

func (f *Foo) _def_init_barf() {

}

//func (f *Foo) _def_lazy_Things() []int {
//	return []int{1, 2, 3}
//}

func foof(s string, a ...any) {
	fmt.Printf(s, a...)
}
