package algo

import (
	"fmt"
	"testing"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	its "github.com/wrmsr/bane/pkg/util/iterators"
)

func TestTopo(t *testing.T) {
	data := map[int][]int{
		0: {1, 2},
		1: {4},
		2: {3, 4},
		3: {1},
		5: {2},
		6: {1, 3},
	}

	ts, err := TopoSort(
		ctr.NewMap(
			its.MapValues(
				its.OfMap(data),
				func(v []int) ctr.Set[int] {
					return ctr.NewSet(its.OfSlice(v))
				})))

	tu.AssertNoErr(t, err)
	fmt.Println(ts)
}
