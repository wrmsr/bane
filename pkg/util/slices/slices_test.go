package slices

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

func TestReverse(t *testing.T) {
	s := []int{1, 2, 3, 4, 5}
	Reverse(s)
	tu.AssertDeepEqual(t, s, []int{5, 4, 3, 2, 1})
}

func TestDeepFlatten(t *testing.T) {
	tu.AssertDeepEqual(t, DeepFlatten[int](1, []int{2, 3}, [][][]int{{{4, 5}, {6, 7}}, {{8, 9}}}), []int{1, 2, 3, 4, 5, 6, 7, 8, 9})
}
