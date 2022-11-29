package main

import (
	"fmt"
	"unsafe"

	ctr "github.com/wrmsr/bane/pkg/util/container"
)

type foo struct{ s string }

//

type _ctr_SliceList[T any] struct{ s []T }

func (l _ctr_SliceList[T]) Len() int { return len(l.s) }

type _ctr_MutSliceList[T any] struct{ l _ctr_SliceList[T] }

func (l *_ctr_MutSliceList[T]) Len() int { return l.l.Len() }

//

type _ctr_foo_SliceList struct{ s []foo }

func (l _ctr_foo_SliceList) Len() int { return len(l.s) }

type _ctr_foo_MutSliceList struct{ l _ctr_foo_SliceList }

func (l *_ctr_foo_MutSliceList) Len() int { return l.l.Len() }

//

func main() {
	lst := ctr.NewMutSliceList[foo](nil)
	lst.Append(foo{s: "hi"})

	fmt.Println(lst.Len())

	p := unsafe.Pointer(&(*lst))
	plst2 := (*_ctr_MutSliceList[foo])(p)
	fmt.Println(plst2.Len())

	p = unsafe.Pointer(&(*lst))
	plst3 := (*_ctr_foo_MutSliceList)(p)
	fmt.Println(plst3.Len())
}
