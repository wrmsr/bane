package algo

import (
	"fmt"
	"testing"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	"github.com/wrmsr/bane/pkg/util/maps"
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

func TestUnify(t *testing.T) {
	l := []maps.Set[int]{
		maps.MakeSetOf([]int{1, 2}),
		maps.MakeSetOf([]int{1, 3}),
		maps.MakeSetOf([]int{4, 5}),
		maps.MakeSetOf([]int{5, 6}),
		maps.MakeSetOf([]int{7, 8}),
	}

	r := Unify(l)

	x := []maps.Set[int]{
		maps.MakeSetOf([]int{1, 2, 3}),
		maps.MakeSetOf([]int{4, 5, 6}),
		maps.MakeSetOf([]int{7, 8}),
	}

	// TODO: FrozenSet...
	tu.AssertEqual(t, len(r), len(x))
	ne := 0
	for _, c0 := range r {
		for _, c1 := range x {
			if c0.Equals(c1) {
				ne++
			}
		}
	}
	tu.AssertEqual(t, ne, 3)
}
