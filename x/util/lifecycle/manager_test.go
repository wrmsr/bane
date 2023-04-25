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

	check.Must(m.Construct())
	check.Must(m.Start())
	check.Must(m.Stop())
	check.Must(m.Destroy())
}
