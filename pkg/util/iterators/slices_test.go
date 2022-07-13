package iterators

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TestSlices(t *testing.T) {
	s := bt.RangeTo(10).Slice()
	tu.AssertDeepEqual(t, Seq(OfSliceRange(s, bt.RangeOf(2, 7, 2))), []int{2, 4, 6})
	tu.AssertDeepEqual(t, Seq(OfSliceRange(s, bt.RangeOf(7, 2, -2))), []int{7, 5, 3})
}
