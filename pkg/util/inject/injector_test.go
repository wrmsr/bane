package inject

import (
	"fmt"
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

type fooInt int

func TestInjector(t *testing.T) {
	numFoos := 0
	inj := NewInjector(Bind(
		420,
		Singleton(func() fooInt {
			numFoos++
			return 2
		}),
	))

	fmt.Println(inj.Debug())

	tu.AssertEqual(t, inj.Provide(Type[int]()).(int), 420)

	tu.AssertEqual(t, numFoos, 0)
	tu.AssertEqual(t, inj.Provide(Type[fooInt]()).(fooInt), 2)
	tu.AssertEqual(t, numFoos, 1)

	mul2 := func(i int, f fooInt) int {
		return i * int(f)
	}

	n := inj.InjectOne(mul2).(int)
	tu.AssertEqual(t, n, 840)

	tu.AssertEqual(t, numFoos, 1)
}

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
