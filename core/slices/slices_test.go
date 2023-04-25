package slices

import (
	"testing"

	tu "github.com/wrmsr/bane/core/dev/testing"
)

func TestReverse(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	Reverse(s)
	tu.AssertDeepEqual(t, s, []int{5, 4, 3, 2, 1})
}

func TestDeepFlatten(t *testing.T) {
	tu.AssertDeepEqual(t, DeepFlatten[int](1, []int{2, 3}, [][][]int{{{4, 5}, {6, 7}}, {{8, 9}}}), []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}

func TestMakeAppend(t *testing.T) {
	s := []int{0, 1, 2, 3}
	r := MakeAppend(s, 4, 5, 6)
	tu.AssertDeepEqual(t, r, []int{0, 1, 2, 3, 4, 5, 6})
	tu.AssertDeepEqual(t, s, []int{0, 1, 2, 3})
}

func TestMakePrepend(t *testing.T) {
	s := []int{0, 1, 2, 3}
	r := MakePrepend(s, 4, 5, 6)
	tu.AssertDeepEqual(t, r, []int{4, 5, 6, 0, 1, 2, 3})
	tu.AssertDeepEqual(t, s, []int{0, 1, 2, 3})
}
