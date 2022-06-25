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
