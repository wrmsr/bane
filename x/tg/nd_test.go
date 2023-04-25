package tg

import (
	"fmt"
	"testing"

	nd "github.com/wrmsr/bane/core/ndarray"
)

func TestNdTensorDot(t *testing.T) {
	fmt.Println(NdTensorDot(
		nd.OfRange[float32](nd.ShapeOf(2, 3, 5)),
		nd.OfRange[float32](nd.ShapeOf(3, 2, 4)),
		nd.DimsOf(0, 1),
		nd.DimsOf(1, 0),
	))

	fmt.Println(NdTensorDot(
		nd.OfRange[float32](nd.ShapeOf(2, 3, 5)),
		nd.OfRange[float32](nd.ShapeOf(3, 2, 4)),
		nd.DimsOf(0),
		nd.DimsOf(1),
	))
}
