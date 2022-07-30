package numexpr

import (
	"fmt"
	"strings"

	iou "github.com/wrmsr/bane/pkg/util/io"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type Node[T bt.Rational] interface {
	isNode()

	Eval() T
	Render(sb *strings.Builder)
}

type node[T bt.Rational] struct{}

func (n node[T]) isNode() {}

//

type Const[T bt.Rational] struct {
	node[T]
	V T
}

var _ Node[int] = Const[int]{}

func NewConst[T bt.Rational](v T) Const[T] {
	return Const[T]{V: v}
}

//

type OpKind int8

const (
	InvalidOp OpKind = iota
	AddOp
	SubOp
	MulOp
	DivOp
	ModOp
)

func (k OpKind) String() string {
	switch k {
	case AddOp:
		return "+"
	case SubOp:
		return "-"
	case MulOp:
		return "*"
	case DivOp:
		return "/"
	case ModOp:
		return "%"
	}
	panic(k)
}

type Op[T bt.Rational] struct {
	node[T]
	K  OpKind
	Vs []Node[T]
}

var _ Node[int] = Op[int]{}

func Add[T bt.Rational](vs ...Node[T]) Op[T] { return Op[T]{K: AddOp, Vs: vs} }
func Sub[T bt.Rational](vs ...Node[T]) Op[T] { return Op[T]{K: SubOp, Vs: vs} }
func Mul[T bt.Rational](vs ...Node[T]) Op[T] { return Op[T]{K: MulOp, Vs: vs} }
func Div[T bt.Rational](vs ...Node[T]) Op[T] { return Op[T]{K: DivOp, Vs: vs} }
func Mod[T bt.Rational](vs ...Node[T]) Op[T] { return Op[T]{K: ModOp, Vs: vs} }

//

func (n Const[T]) Eval() T {
	return n.V
}

func (n Op[T]) Eval() T {
	switch n.K {
	case AddOp:
		var r T
		for _, v := range n.Vs {
			r += v.Eval()
		}
		return r
	case SubOp:
		var r T
		for _, v := range n.Vs {
			r -= v.Eval()
		}
		return r
	case MulOp:
		var r T
		for _, v := range n.Vs {
			r *= v.Eval()
		}
		return r
	case DivOp:
		var r T
		for _, v := range n.Vs {
			r /= v.Eval()
		}
		return r
	case ModOp:
		panic("nyi")
	}
	panic(n.K)
}

//

func (n Const[T]) Render(sb *strings.Builder) {
	iou.Discard(fmt.Fprintf(sb, "%v", n.V))
}

func (n Op[T]) Render(sb *strings.Builder) {
	for i, v := range n.Vs {
		if i > 0 {
			iou.Discard(sb.WriteString(n.K.String()))
		}
		v.Render(sb)
	}
}
