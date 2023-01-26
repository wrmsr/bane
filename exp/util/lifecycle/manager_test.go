package lifecycles

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestManager(t *testing.T) {
	m := NewManager()

	check.Must(m.AddObject("a", func(st State) error {
		fmt.Printf("a: %v\n", st)
		return nil
	}, nil))
	check.Must(m.AddObject("b", func(st State) error {
		fmt.Printf("b: %v\n", st)
		return nil
	}, []any{"a"}))

	check.Must(m.advance(&construct))
	check.Must(m.advance(&start))
	check.Must(m.advance(&stop))
	check.Must(m.advance(&destroy))
}
