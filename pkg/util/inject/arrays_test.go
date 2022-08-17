package inject

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

func TestArrays(t *testing.T) {
	inj := NewInjector(Bind(
		As(KeyOf[string](), "hi"),
		As(Array(KeyOf[int]()), 420),
		As(Array(KeyOf[int]()), 421),
	))

	tu.AssertDeepEqual(t, inj.Provide(KeyOf[string]()), "hi")
	tu.AssertDeepEqual(t, inj.Provide(Array(KeyOf[int]())), []int{420, 421})

	inj2 := inj.NewChild(Bind(
		As(KeyOf[[]int](), Link(Array(KeyOf[int]()))),
	))
	tu.AssertDeepEqual(t, inj2.Provide(KeyOf[[]int]()), []int{420, 421})
}

func TestEmptyArray(t *testing.T) {
	inj := NewInjector(Bind(
		BindArrayOf[int],
	))

	tu.AssertDeepEqual(t, inj.Provide(Array(KeyOf[int]())), []int{})

	inj = NewInjector(Bind(
		As(Array(KeyOf[int]()), EmptyOf[int]()),
		As(Array(KeyOf[int]()), 420),
		As(Array(KeyOf[int]()), EmptyOf[int]()),
	))

	tu.AssertDeepEqual(t, inj.Provide(Array(KeyOf[int]())), []int{420})
}
