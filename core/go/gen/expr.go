package gen

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wrmsr/bane/core/check"
	"github.com/wrmsr/bane/core/slices"
)

//

type Expr interface {
	Node
	isExpr()
}

type expr struct {
	node
}

func (e expr) isExpr() {}

//

func ExprOf(o any) Expr {
	if o, ok := o.(Expr); ok {
		return o
	}

	switch o := o.(type) {
	case Func:
		return FuncExprOf(o)
	case string:
		if strings.Contains(o, ".") {
			is := slices.Map(func(p string) any { return Ident{Name: p} }, strings.Split(o, "."))
			return SelectOf(is[0], is[1:]...)
		}
		return IdentOf(o)
	case int:
		return LitOf(strconv.Itoa(o))
	}

	panic(o)
}

//

type Exprs []Expr

func (s *Exprs) Append(os ...any) {
	for _, o := range os {
		*s = append(*s, ExprOf(o))
	}
}

//

type Addr struct {
	Value Expr

	expr
}

//

type BinaryOp int8

const (
	InvalidBinaryOp BinaryOp = iota

	AddOp
	SubOp
	MulOp
	DivOp
	ModOp

	LtOp
	LteOp
	EqOp
	NeOp
	GtOp
	GteOp

	BitAndOp
	BitOrOp
	BitXorOp

	AndOp
	OrOp
)

func (o BinaryOp) String() string {
	switch o {

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

	case LtOp:
		return "<"
	case LteOp:
		return "<="
	case EqOp:
		return "="
	case NeOp:
		return "!="
	case GtOp:
		return ">"
	case GteOp:
		return ">="

	case BitAndOp:
		return "&"
	case BitOrOp:
		return "|"
	case BitXorOp:
		return "^"

	case AndOp:
		return "&&"
	case OrOp:
		return "||"

	}
	panic(fmt.Errorf("invalid binary op: %d", o))
}

func (o BinaryOp) IsBit() bool {
	switch o {
	case
		BitAndOp,
		BitOrOp,
		BitXorOp:
		return true
	}
	return false
}

func (o BinaryOp) IsBoolean() bool {
	switch o {
	case
		AndOp,
		OrOp:
		return true
	}
	return false
}

type Binary struct {
	Op   BinaryOp
	Args []Expr

	expr
}

func BinaryOf(op BinaryOp, args ...Expr) Binary {
	return Binary{
		Op:   op,
		Args: check.NotEmptySlice(args),
	}
}

func Add(args ...Expr) Expr { return BinaryOf(AddOp, args...) }
func Sub(args ...Expr) Expr { return BinaryOf(SubOp, args...) }
func Mul(args ...Expr) Expr { return BinaryOf(MulOp, args...) }
func Div(args ...Expr) Expr { return BinaryOf(DivOp, args...) }
func Mod(args ...Expr) Expr { return BinaryOf(ModOp, args...) }

func Lt(args ...Expr) Expr  { return BinaryOf(LtOp, args...) }
func Lte(args ...Expr) Expr { return BinaryOf(LteOp, args...) }
func Eq(args ...Expr) Expr  { return BinaryOf(EqOp, args...) }
func Ne(args ...Expr) Expr  { return BinaryOf(NeOp, args...) }
func Gt(args ...Expr) Expr  { return BinaryOf(GtOp, args...) }
func Gte(args ...Expr) Expr { return BinaryOf(GteOp, args...) }

func BitAnd(args ...Expr) Expr { return BinaryOf(BitAndOp, args...) }
func BitOr(args ...Expr) Expr  { return BinaryOf(BitOrOp, args...) }
func BitXor(args ...Expr) Expr { return BinaryOf(BitXorOp, args...) }

func And(args ...Expr) Expr { return BinaryOf(AndOp, args...) }
func Or(args ...Expr) Expr  { return BinaryOf(OrOp, args...) }

//

type Call struct {
	Func Expr
	Args []Expr

	expr
}

func CallOf(func_ any, args ...any) Call {
	return Call{
		Func: ExprOf(func_),
		Args: slices.Map(ExprOf, args),
	}
}

//

type Deref struct {
	Value Expr

	expr
}

//

type FuncExpr struct {
	Func Func

	expr
}

func FuncExprOf(func_ Func) FuncExpr {
	return FuncExpr{
		Func: func_,
	}
}

//

type Ident struct {
	Name string

	expr
}

func NewIdent(name string) *Ident {
	return &Ident{
		Name: check.NotZero(name),
	}
}

func IdentOf(o any) Ident {
	switch o := o.(type) {
	case Ident:
		return o
	case *Ident:
		return *o
	case string:
		check.NotZero(o)
		// FIXME:
		//check.Condition(!strings.Contains(o, "."))
		return Ident{
			Name: o,
		}
	}
	panic(o)
}

//

type Index struct {
	Value Expr
	Index Expr

	expr
}

func IndexOf(value, index any) Index {
	return Index{
		Value: ExprOf(value),
		Index: ExprOf(index),
	}
}

//

type Lit struct {
	String string

	expr
}

func LitOf(s string) Lit {
	return Lit{
		String: check.NotZero(s),
	}
}

//

type Paren struct {
	Value Expr

	expr
}

func ParenOf(value any) Paren {
	return Paren{
		Value: ExprOf(value),
	}
}

//

type Select struct {
	Value Expr
	Names []Ident

	expr
}

func SelectOf(value any, names ...any) Select {
	return Select{
		Value: ExprOf(value),
		Names: slices.Map(IdentOf, names),
	}
}

//

type TypeAssert struct {
	Value Expr
	Type  Type

	expr
}

func TypeAssertOf(value, type_ any) TypeAssert {
	return TypeAssert{
		Value: ExprOf(value),
		Type:  TypeOf(type_),
	}
}

//

type UnaryOp int8

const (
	InvalidUnaryOp UnaryOp = iota

	NotOp
	NegOp
	InvOp

	IncOp
	DecOp
)

func (o UnaryOp) String() string {
	switch o {

	case NotOp:
		return "!"
	case NegOp:
		return "-"
	case InvOp:
		return "^"

	case IncOp:
		return "++"
	case DecOp:
		return "--"

	}
	panic(fmt.Errorf("invalid unary op: %d", o))
}

func (o UnaryOp) IsPrefix() bool {
	switch o {

	case
		NotOp,
		NegOp,
		InvOp:
		return true

	case
		IncOp,
		DecOp:
		return false

	}
	panic(o)
}

type Unary struct {
	Op  UnaryOp
	Arg Expr

	expr
}

func UnaryOf(op UnaryOp, arg any) Unary {
	return Unary{
		Op:  op,
		Arg: ExprOf(arg),
	}
}

func Not(arg Expr) Expr { return UnaryOf(NotOp, arg) }
func Neg(arg Expr) Expr { return UnaryOf(NegOp, arg) }
func Inv(arg Expr) Expr { return UnaryOf(InvOp, arg) }

func Inc(arg Expr) Expr { return UnaryOf(IncOp, arg) }
func Dec(arg Expr) Expr { return UnaryOf(DecOp, arg) }
