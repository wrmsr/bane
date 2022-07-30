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

func TestReshapeAdd1s(t *testing.T) {
	// self.st = ShapeTracker((4, 4))
	// self.st.permute(1, 0)
	// self.st.reshape(1, 4, 1, 4, 1)
	// assert not self.st.contiguous
	// self.st.permute(0, 3, 2, 1, 4)
	// assert self.st.contiguous
}
