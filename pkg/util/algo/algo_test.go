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

	var m ctr.Map[int, ctr.Set[int]] = ctr.NewStdMap(
		its.MapValues(
			its.OfMap(data),
			func(v []int) ctr.Set[int] {
				return ctr.NewStdSet(its.OfSlice(v))
			}))

	ts, err := TopoSort(m)

	tu.AssertNoErr(t, err)
	fmt.Println(ts)
}
