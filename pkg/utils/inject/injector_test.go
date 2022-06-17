package inject

import (
	"reflect"
	"testing"

	bts "github.com/wrmsr/bane/pkg/dev/testing"
	"github.com/wrmsr/bane/pkg/utils/log"
)

func TestInjector(t *testing.T) {
	inj := NewInjector(
		Bindings{
			bs: []Binding{
				{
					key: Key{
						ty: reflect.TypeOf(0),
					},
					provider: Const(420),
				},
			},
		},
	)
	bts.AssertEqual(t, inj.Provide(Key{ty: reflect.TypeOf(0)}).(int), 420)

	mul2 := func(i int) int {
		return i * 2
	}

	a := inj.ProvideArgs(mul2)
	log.Println(a)
}
