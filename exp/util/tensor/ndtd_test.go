package tensor

import (
	"fmt"
	"testing"

	nd "github.com/wrmsr/bane/pkg/util/ndarray"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TestNdTd(t *testing.T) {
	// [[[[[[0.]]],,,   [[[3.]]],,,   [[[6.]]]],,,,  [[[[1.]]],,,   [[[4.]]],,,   [[[7.]]]],,,,  [[[[2.]]],,,   [[[5.]]],,,   [[[8.]]]]]]
	txg := nd.Of[float32](
		nd.ShapeOf(1, 3, 3, 1, 1, 1),
		nd.Strides{},
		0,
		bt.RangeTo[float32](9).Slice(),
	)

	fmt.Println(txg)

	// [[[[10.]],,  [[13.]],,  [[16.]]],,, [[[11.]],,  [[14.]],,  [[17.]]],,, [[[12.]],,  [[15.]],,  [[18.]]]]
	twg := nd.Of[float32](
		nd.ShapeOf(3, 3, 1, 1),
		nd.Strides{},
		0,
		[]float32{10, 13, 16, 11, 14, 17, 12, 15, 18},
	)

	fmt.Println(twg)

	// [[[[ 45.  48.  51.]],,  [[162. 174. 186.]],,  [[279. 300. 321.]]]]
	want := nd.Of[float32](
		nd.ShapeOf(1, 3, 1, 3),
		nd.Strides{},
		0,
		[]float32{45, 48, 51, 162, 174, 186, 279, 300, 321},
	)

	fmt.Println(want)

	// np.tensordot(txg, twg, ((1, 4, 5), (1, 2, 3)))
}
