package iterators

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/utils/dev/testing"
)

func TestTransform(t *testing.T) {
	tu.AssertDeepEqual(t, Take[int](Transform(Count(), func(v int) int { return v * 2 }), 5), []int{0, 2, 4, 6, 8})
}

func TestFilter(t *testing.T) {
	tu.AssertDeepEqual(t, Take[int](Filter(Count(), func(v int) bool { return v%3 == 0 }), 5), []int{0, 3, 6, 9, 12})
}
