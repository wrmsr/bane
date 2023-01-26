package lifecycles

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestManager(t *testing.T) {
	m := NewManager()

	afn := func(st State) error {
		fmt.Printf("a: %v\n", st)
		return nil
	}

	bfn := func(st State) error {
		fmt.Printf("b: %v\n", st)
		return nil
	}

	check.Must(m.Add("a", afn, nil))
	check.Must(m.Add("b", bfn, []any{"a"}))

	check.Must(m.advance(&construct))
	check.Must(m.advance(&start))
	check.Must(m.advance(&stop))
	check.Must(m.advance(&destroy))
}
