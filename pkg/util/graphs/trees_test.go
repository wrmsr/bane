package graphs

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	its "github.com/wrmsr/bane/pkg/util/iterators"
)

type Calc interface {
	isCalc()
}

type Num struct{}
type Add struct{ L, R Calc }

func (c *Num) isCalc() {}
func (c *Add) isCalc() {}

func TestTrees(t *testing.T) {
	root := &Add{
		&Add{
			&Num{},
			&Num{},
		},
		&Num{},
	}

	walk := func(n Calc) its.Iterable[Calc] {
		switch n := n.(type) {
		case *Num:
			return its.Of[Calc]()
		case *Add:
			return its.Of(n.L, n.R)
		}
		panic("unreachable")
	}

	tree := check.Must1(NewTree[Calc](root, walk))
	fmt.Println(tree)
}
