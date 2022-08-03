package tg

import (
	"fmt"
	"testing"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TestOps(t *testing.T) {
	//a := -.5
	//b := 20.
	//rand.Seed(0)

	sh := Shape{3, 3}

	xs := BufferOf(sh, bt.RangeTo[float32](9.).Slice())
	ys := BufferOf(sh, bt.RangeOf[float32](10., 19., 1.).Slice())

	xt := NewTensor(MakeLoadBuffer(xs, sh), true)
	yt := NewTensor(MakeLoadBuffer(ys, sh), true)

	zt := xt.Add(yt)
	fmt.Println(zt)

	fmt.Println(zt.Data().Realize())

	zt.Mean(nil, false).Backward()

	zg := zt.grad.Data().Realize()

	fmt.Println(zg)
}
