package inject

import (
	"reflect"
	"testing"

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

	f := func(i int) {
		log.Println(i)
	}

	a := inj.ProvideArgs(f)
	log.Println(a)
}
