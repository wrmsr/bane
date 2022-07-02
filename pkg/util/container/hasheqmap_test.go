package container

import (
	"fmt"
	"testing"

	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TestHashEqMap(t *testing.T) {
	he := bt.HashEqOf(
		func(v int) uintptr { return uintptr(v % 3) },
		func(l, r int) bool { return l == r },
	)
	m := newHashEqMapImpl[int, int](he, its.OfSlice(bt.KvsOf[int, int](
		0, 10,
		11, 20,
		12, 21,
		13, 22,
		14, 23,
		15, 24,
		16, 25,
		20, 30,
	)))
	fmt.Println(m)
}
