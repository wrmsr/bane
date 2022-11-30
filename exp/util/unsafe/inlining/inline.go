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

func (l _ctr_SliceList[T]) Get(i int) T { return l.s[i] }

type _ctr_MutSliceList[T any] struct{ l _ctr_SliceList[T] }

func (l *_ctr_MutSliceList[T]) Len() int { return l.l.Len() }

func (l *_ctr_MutSliceList[T]) Get(i int) T { return l.l.Get(i) }

//

type _ctr_foo_SliceList struct{ s []foo }

func (l _ctr_foo_SliceList) Len() int { return len(l.s) }

func (l _ctr_foo_SliceList) Get(i int) foo { return l.s[i] }

type _ctr_foo_MutSliceList struct{ l _ctr_foo_SliceList }

func (l *_ctr_foo_MutSliceList) Len() int { return l.l.Len() }

func (l *_ctr_foo_MutSliceList) Get(i int) foo { return l.l.Get(i) }

//

func main() {
	lst := ctr.NewMutSliceList[foo](nil)
	lst.Append(foo{s: "hi"})

	// no inline, cross package and generic
	fmt.Println(lst.Len())
	fmt.Println(lst.Get(0))

	// no inline, generic
	p := unsafe.Pointer(&(*lst))
	plst2 := (*_ctr_MutSliceList[foo])(p)
	fmt.Println(plst2.Len())
	fmt.Println(plst2.Get(0))

	// inlines
	p = unsafe.Pointer(&(*lst))
	plst3 := (*_ctr_foo_MutSliceList)(p)
	fmt.Println(plst3.Len())
	fmt.Println(plst3.Get(0))
}
