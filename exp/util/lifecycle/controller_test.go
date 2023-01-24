package lifecycles

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestController(t *testing.T) {
	lc := &Lifecycle{
		Fn: func(st State) error {
			fmt.Println(st)
			return nil
		},
	}

	c := &controller{}

	check.Must(c.advance(lc, &construct))
}
