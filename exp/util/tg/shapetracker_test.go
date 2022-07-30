package tg

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TestStridesForShape(t *testing.T) {
	strides := StridesForShape(bt.RangeTo[Dim](10).Slice())
	tu.AssertDeepEqual(t, strides, Strides{362880, 362880, 181440, 60480, 15120, 3024, 504, 72, 9, 1})
}
