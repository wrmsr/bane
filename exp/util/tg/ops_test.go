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

	xs := &Buffer{bt.RangeTo[float32](10.).Slice()}
	ys := &Buffer{bt.RangeOf[float32](10., 20., 1.).Slice()}

	mkLoadBuffer := func(data *Buffer, shape Shape) *LazyBuffer {
		return NewLazyBuffer(
			NewShapeTracker(shape),
			LoadOpType,
			&LazyOp{
				op:   FromCpuOp,
				srcs: nil,
				arg:  data,
			},
		)
	}

	xb := mkLoadBuffer(xs, Shape{3, 3})
	yb := mkLoadBuffer(ys, Shape{3, 3})

	xt := NewTensor(xb, true)
	yt := NewTensor(yb, true)

	zt := xt.Add(yt)
	fmt.Println(zt)
}
