package graphs

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/core/check"
	its "github.com/wrmsr/bane/core/iterators"
	rfl "github.com/wrmsr/bane/core/reflect"
	bt "github.com/wrmsr/bane/core/types"
)

///

type Calc interface {
	Accept(v CalcVisitor)
	Children() []Calc

	isCalc()
}

type Num struct{ N int }
type Add struct{ L, R Calc }
type Sub struct{ L, R Calc }

func (c *Num) isCalc() {}
func (c *Add) isCalc() {}
func (c *Sub) isCalc() {}

var testCalc = &Add{
	&Sub{
		&Add{
			&Num{1},
			&Num{2},
		},
		&Num{1},
	},
	&Num{3},
}

//

func (c *Num) String() string { return fmt.Sprintf("Num@%p{%d}", c, c.N) }
func (c *Add) String() string { return fmt.Sprintf("Add@%p{%s, %s}", c, c.L, c.R) }
func (c *Sub) String() string { return fmt.Sprintf("Sub@%p{%s, %s}", c, c.L, c.R) }

func (c *Num) Children() []Calc { return nil }
func (c *Add) Children() []Calc { return []Calc{c.L, c.R} }
func (c *Sub) Children() []Calc { return []Calc{c.L, c.R} }

///

func TestTrees(t *testing.T) {
	walk := func(n Calc) bt.Iterable[Calc] {
		switch n := n.(type) {
		case *Num:
			return its.Of[Calc]()
		case *Add:
			return its.Of(n.L, n.R)
		case *Sub:
			return its.Of(n.L, n.R)
		}
		panic("unreachable")
	}

	fmt.Println(testCalc.String())

	tree := check.Must1(NewTree[Calc](testCalc, walk, rfl.AddrHashEq[Calc]()))
	fmt.Println(tree)
}

///

type CalcVisitor interface {
	VisitNum(c *Num)
	VisitAdd(c *Add)
	VisitSub(c *Sub)
}

func (c *Num) Accept(v CalcVisitor) { v.VisitNum(c) }
func (c *Add) Accept(v CalcVisitor) { v.VisitAdd(c) }
func (c *Sub) Accept(v CalcVisitor) { v.VisitSub(c) }

//

type SuperCalcVisitor struct {
	Super CalcVisitor
}

var _ CalcVisitor = &SuperCalcVisitor{}

func (v SuperCalcVisitor) SuperVisitChildren(c Calc) {
	for _, r := range c.Children() {
		r.Accept(v.Super)
	}
}

func (v SuperCalcVisitor) VisitNum(c *Num) { v.SuperVisitChildren(c) }
func (v SuperCalcVisitor) VisitAdd(c *Add) { v.SuperVisitChildren(c) }
func (v SuperCalcVisitor) VisitSub(c *Sub) { v.SuperVisitChildren(c) }

//

type AddPrintingCalcVisitor struct {
	SuperCalcVisitor
}

func (v AddPrintingCalcVisitor) VisitAdd(c *Add) {
	fmt.Printf("add!: %v\n", c)
	v.SuperVisitChildren(c)
}

func TestSuperVisitor(t *testing.T) {
	v := AddPrintingCalcVisitor{}
	v.Super = &v

	testCalc.Accept(&v)
}
