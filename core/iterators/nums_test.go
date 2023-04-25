package iterators

import (
	"testing"

	tu "github.com/wrmsr/bane/core/dev/testing"
)

func TestCount(t *testing.T) {
	tu.AssertDeepEqual(t, Take[int](Count(), 4), []int{0, 1, 2, 3})
}

func TestRange(t *testing.T) {
	tu.AssertDeepEqual(t, Seq[int](Range(0, 20, 3)), []int{0, 3, 6, 9, 12, 15, 18})
	tu.AssertDeepEqual(t, Seq[int](Range(10, -10, -3)), []int{10, 7, 4, 1, -2, -5, -8})
}
