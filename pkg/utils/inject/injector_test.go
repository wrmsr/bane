package inject

import (
	"reflect"
	"testing"

	tu "github.com/wrmsr/bane/pkg/utils/dev/testing"
	"github.com/wrmsr/bane/pkg/utils/log"
)

func TestInjector(t *testing.T) {
	inj := NewInjector(
		Bindings{
			bs: []Binding{
				{
					key:      ToKey(reflect.TypeOf(0)),
					provider: Const(420),
				},
			},
		},
	)
	tu.AssertEqual(t, inj.Provide(Key{ty: reflect.TypeOf(0)}).(int), 420)

	mul2 := func(i int) int {
		return i * 2
	}

	n := inj.InjectOne(mul2)
	log.Println(n)
}
