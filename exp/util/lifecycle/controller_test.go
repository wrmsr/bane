package lifecycles

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestController(t *testing.T) {
	lc := &Lifecycle{
		Construct: func() error {
			fmt.Println("constructing")
			return nil
		},
		Start: func() error {
			fmt.Println("starting")
			return nil
		},
		Stop: func() error {
			fmt.Println("stopping")
			return nil
		},
		Destroy: func() error {
			fmt.Println("destroying")
			return nil
		},
	}

	c := &controller{
		lc: lc,
	}

	check.Must(c.advance(&construct))
}
