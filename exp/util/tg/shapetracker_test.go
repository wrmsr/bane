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
	st := NewShapeTracker(Shape{4, 4})
	st.Permute(1, 0)
	st.Reshape(Shape{1, 4, 1, 4, 1})
	tu.AssertEqual(t, st.Contiguous(), false)
	st.Permute(0, 3, 2, 1, 4)
	tu.AssertEqual(t, st.Contiguous(), true)
}
