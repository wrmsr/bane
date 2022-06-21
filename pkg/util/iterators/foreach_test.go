package iterators

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

type CanForEachInt interface {
	Traversable[int]
}

type FooForEach struct{}

var _ Traversable[int] = FooForEach{}
var _ CanForEachInt = FooForEach{}

func (f FooForEach) ForEach(fn func(v int) bool) bool {
	return fn(1) && fn(2) && fn(3)
}

func TestForEach(t *testing.T) {
	var a []int
	FooForEach{}.ForEach(func(v int) bool {
		a = append(a, v)
		return true
	})
	tu.AssertDeepEqual(t, a, []int{1, 2, 3})
}
