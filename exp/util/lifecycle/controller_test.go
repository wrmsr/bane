package lifecycles

import (
	"fmt"
	"testing"
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

	c := &Controller{
		lc: lc,
	}

	c.advance(&construct)
}
