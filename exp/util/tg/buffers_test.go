package tg

import (
	"fmt"
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TestBufferDot(t *testing.T) {
	arange := func(sh Shape) *Buffer {
		return BufferOf(sh, bt.RangeTo[float32](float32(sh.Dim())).Slice())
	}

	a := arange(Shape{3, 4, 5})
	b := arange(Shape{4, 3, 2})
	c := a.Dot(b, []int{1, 0}, []int{0, 1})
	fmt.Println(c)

	tu.AssertDeepEqual(t, c.shape, Shape{5, 2})
	tu.AssertDeepEqual(t, c.s, []float32{
		4400., 4730.,
		4532., 4874.,
		4664., 5018.,
		4796., 5162.,
		4928., 5306.,
	})
}
