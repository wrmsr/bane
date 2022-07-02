package graphs

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

type Calc interface {
	isCalc()
}

type Num struct{ N int }
type Add struct{ L, R Calc }

func (c *Num) isCalc() {}
func (c *Add) isCalc() {}

func (c *Num) String() string { return fmt.Sprintf("Num@%p{%d}", c, c.N) }
func (c *Add) String() string { return fmt.Sprintf("Add@%p{%s, %s}", c, c.L, c.R) }

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

	fmt.Println(root.String())

	tree := check.Must1(NewTree[Calc](root, walk, rfl.AddrHashEq[Calc]()))
	fmt.Println(tree)
}
