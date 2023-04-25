package tg

import (
	"fmt"
	"testing"

	tu "github.com/wrmsr/bane/core/dev/testing"
	bt "github.com/wrmsr/bane/core/types"
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

func TestPermute1sSimple(t *testing.T) {
	st := NewShapeTracker(Shape{1, 16, 9, 9})
	st.Permute(1, 0, 2, 3)
	tu.AssertEqual(t, st.Contiguous(), true)
	st = NewShapeTracker(Shape{2, 16, 9, 9})
	st.Permute(1, 0, 2, 3)
	tu.AssertEqual(t, st.Contiguous(), false)
}

func TestRemove1sSimple(t *testing.T) {
	st := NewShapeTracker(Shape{1, 16, 1, 1})
	st.Reshape(Shape{16})
	tu.AssertEqual(t, st.Contiguous(), true)
}

func TestRemove1s(t *testing.T) {
	st := NewShapeTracker(Shape{1, 4, 1, 4, 1})
	st.Permute(0, 3, 2, 1, 4)
	st.Reshape(Shape{4, 4})
	tu.AssertEqual(t, st.Contiguous(), false)
	st.Permute(1, 0)
	tu.AssertEqual(t, st.Contiguous(), true)
}

func TestWork(t *testing.T) {
	st := NewShapeTracker(Shape{64, 1024, 4})
	st.Reshape(Shape{1, 64, 128, 32})
	st.Permute(0, 3, 1, 2)
	st.Reshape(Shape{1, 32, 1, 64, 128})
	st.Permute(0, 3, 4, 1, 2)
	tu.AssertEqual(t, st.Contiguous(), true)
}

func TestWork2(t *testing.T) {
	st := NewShapeTracker(Shape{64, 1024, 4})
	st.Reshape(Shape{1, 64, 128, 32})
	st.Permute(0, 3, 1, 2)
	st.Reshape(Shape{1, 1, 32, 64, 128})
	st.Permute(0, 3, 4, 1, 2)
	st.Reshape(Shape{64, 1024, 4})
	fmt.Println(st.views)
	tu.AssertEqual(t, st.Contiguous(), true)
}

func TestDoubleStride(t *testing.T) {
	st := NewShapeTracker(Shape{7, 4})
	st.Stride(1, 2)
	st.Stride(2, 1)
	tu.AssertDeepEqual(t, st.Shape(), Shape{4, 2})
	tu.AssertDeepEqual(t, st.Strides(), Strides{8, 2})
}
