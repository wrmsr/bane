package container

import (
	"fmt"
	"testing"

	its "github.com/wrmsr/bane/pkg/util/iterators"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TestSliceMapIt(t *testing.T) {
	m := NewMutSliceMap[int, string](nil)
	m.Put(1, "one")
	m.Put(2, "two")
	m.Put(3, "three")
	m.Put(4, "four")

	fmt.Println(its.Seq[bt.Kv[int, string]](m))
	fmt.Println(its.Seq[bt.Kv[int, string]](m.ReverseIterate()))

	fmt.Println(its.Seq[bt.Kv[int, string]](m.IterateFrom(2)))
	fmt.Println(its.Seq[bt.Kv[int, string]](m.ReverseIterateFrom(2)))

	fmt.Println(its.Seq[bt.Kv[int, string]](m.IterateFrom(8)))
}
