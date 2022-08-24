package slices

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TestReshape(t *testing.T) {
	s := bt.RangeTo(10).Slice()
	tu.AssertDeepEqual(t, Reshape(s, 8, -1, 1), []int{8})
	tu.AssertDeepEqual(t, Reshape(s, 1, 6, 2), []int{1, 3, 5})
	tu.AssertDeepEqual(t, Reshape(s, 6, 1, -2), []int{6, 4, 2})
	tu.AssertDeepEqual(t, Reshape(s, 3, -1, 1), []int{3, 4, 5, 6, 7, 8})
	tu.AssertDeepEqual(t, Reshape(s, 3, -1, 3), []int{3, 6})
	tu.AssertDeepEqual(t, Reshape(s, -2, 2, -3), []int{8, 5})
}
