package inject

import (
	"fmt"
	"testing"

	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
)

func TestBinderRegistry(t *testing.T) {
	reg := NewBinderRegistry().
		Register(func() Bindings {
			return Bind(
				420,
				func(i int) string {
					return fmt.Sprintf("%d!", i)
				},
			)
		})

	inj := NewInjector(reg.Bind())
	tu.AssertEqual(t, inj.Provide((*int)(nil)).(int), 420)
	tu.AssertEqual(t, inj.Provide((*string)(nil)).(string), "420!")
}
