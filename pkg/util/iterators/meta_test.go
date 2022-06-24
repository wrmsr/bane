package iterators

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

func TestMap(t *testing.T) {
	tu.AssertDeepEqual(t, Take[int](Map(Count(), func(v int) int { return v * 2 }), 5), []int{0, 2, 4, 6, 8})
}

func TestFilter(t *testing.T) {
	tu.AssertDeepEqual(t, Take[int](Filter(Count(), func(v int) bool { return v%3 == 0 }), 5), []int{0, 3, 6, 9, 12})
}

func TestFlatten(t *testing.T) {
	tu.AssertDeepEqual(t, Take[int](Flatten(Map(Count(), func(i int) Iterable[int] { return Of[int](i, i*2) })), 7), []int{0, 0, 1, 2, 2, 4, 3})
}

func TestChunk(t *testing.T) {
	tu.AssertDeepEqual(t, Seq[[]int](Chunk(Range(0, 10, 1), 3)), [][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, {9}})
}
