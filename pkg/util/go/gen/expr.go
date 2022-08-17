package gen

import (
	"fmt"

	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/slices"
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

	switch o.(type) {
	case string:
		return IdentOf(o)
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
		return Ident{
			Name: check.NotZero(o),
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

func IndexOf(value, index Expr) Index {
	return Index{
		Value: value,
		Index: index,
	}
}

//

type InfixOp int8

const (
	InvalidInfixOp InfixOp = iota

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

	AndOp
	OrOp
)

func (o InfixOp) String() string {
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

	case AndOp:
		return "&&"
	case OrOp:
		return "||"

	}
	panic(fmt.Errorf("invalid infix op: %d", o))
}

func (o InfixOp) IsBoolean() bool {
	switch o {
	case
		AndOp,
		OrOp:
		return true
	}
	return false
}

type Infix struct {
	Op   InfixOp
	Args []Expr

	expr
}

func InfixOf(op InfixOp, args ...Expr) Infix {
	return Infix{
		Op:   op,
		Args: check.NotEmptySlice(args),
	}
}

func Add(args ...Expr) Expr { return InfixOf(AddOp, args...) }
func Sub(args ...Expr) Expr { return InfixOf(SubOp, args...) }
func Mul(args ...Expr) Expr { return InfixOf(MulOp, args...) }
func Div(args ...Expr) Expr { return InfixOf(DivOp, args...) }
func Mod(args ...Expr) Expr { return InfixOf(ModOp, args...) }

func Lt(args ...Expr) Expr  { return InfixOf(LtOp, args...) }
func Lte(args ...Expr) Expr { return InfixOf(LteOp, args...) }
func Eq(args ...Expr) Expr  { return InfixOf(EqOp, args...) }
func Ne(args ...Expr) Expr  { return InfixOf(NeOp, args...) }
func Gt(args ...Expr) Expr  { return InfixOf(GtOp, args...) }
func Gte(args ...Expr) Expr { return InfixOf(GteOp, args...) }

func And(args ...Expr) Expr { return InfixOf(AndOp, args...) }
func Or(args ...Expr) Expr  { return InfixOf(OrOp, args...) }

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

func ParenOf(value Expr) Paren {
	return Paren{
		Value: value,
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

func TypeAssertOf(value Expr, type_ Type) TypeAssert {
	return TypeAssert{
		Value: value,
		Type:  type_,
	}
}

//

type UnaryOp int8

const (
	InvalidUnaryOp UnaryOp = iota

	NotOp
)

func (o UnaryOp) String() string {
	switch o {

	case NotOp:
		return "!"

	}
	panic(fmt.Errorf("invalid unary op: %d", o))
}

type Unary struct {
	Op  UnaryOp
	Arg Expr

	expr
}

func UnaryOf(op UnaryOp, arg Expr) Unary {
	return Unary{
		Op:  op,
		Arg: check.NotNil(arg).(Expr),
	}
}

func Not(arg Expr) Expr { return UnaryOf(NotOp, arg) }
