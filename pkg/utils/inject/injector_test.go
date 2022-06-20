package inject

import (
	"testing"

	tu "github.com/wrmsr/bane/pkg/utils/dev/testing"
	"github.com/wrmsr/bane/pkg/utils/log"
)

type fooInt int

func TestInjector(t *testing.T) {
	inj := NewInjector(
		bindings{
			bs: []Binding{
				toBinding(420),
				toBinding(func() fooInt { return 2 }),
			},
		},
	)
	tu.AssertEqual(t, inj.Provide(Key{ty: Type[int]()}).(int), 420)

	mul2 := func(i int, f fooInt) int {
		return i * int(f)
	}

	n := inj.InjectOne(mul2)
	log.Println(n)
}
