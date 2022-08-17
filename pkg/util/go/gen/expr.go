package gen

import (
	"fmt"

	"github.com/wrmsr/bane/pkg/util/check"
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

func CallOf(func_ Expr, args ...Expr) Call {
	return Call{
		Func: func_,
		Args: args,
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

func IdentOf(name string) Ident {
	return Ident{
		Name: check.NotZero(name),
	}
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

type InfixExpr struct {
	Op   InfixOp
	Args []Expr

	expr
}

func InfixExprOf(op InfixOp, args ...Expr) InfixExpr {
	return InfixExpr{
		Op:   op,
		Args: check.NotEmptySlice(args),
	}
}

func InfixExprOrSelf(op InfixOp, args ...Expr) Expr {
	if len(args) == 1 {
		return args[0]
	}
	return InfixExprOf(
		op,
		args...,
	)
}

func Add(args ...Expr) Expr { return InfixExprOrSelf(AddOp, args...) }
func Sub(args ...Expr) Expr { return InfixExprOrSelf(SubOp, args...) }
func Mul(args ...Expr) Expr { return InfixExprOrSelf(MulOp, args...) }
func Div(args ...Expr) Expr { return InfixExprOrSelf(DivOp, args...) }
func Mod(args ...Expr) Expr { return InfixExprOrSelf(ModOp, args...) }

func Lt(args ...Expr) Expr  { return InfixExprOrSelf(LtOp, args...) }
func Lte(args ...Expr) Expr { return InfixExprOrSelf(LteOp, args...) }
func Eq(args ...Expr) Expr  { return InfixExprOrSelf(EqOp, args...) }
func Ne(args ...Expr) Expr  { return InfixExprOrSelf(NeOp, args...) }
func Gt(args ...Expr) Expr  { return InfixExprOrSelf(GtOp, args...) }
func Gte(args ...Expr) Expr { return InfixExprOrSelf(GteOp, args...) }

func And(args ...Expr) Expr { return InfixExprOrSelf(AndOp, args...) }
func Or(args ...Expr) Expr  { return InfixExprOrSelf(OrOp, args...) }

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

func SelectOf(value Expr, names ...Ident) Select {
	return Select{
		Value: value,
		Names: names,
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

type UnaryExpr struct {
	Op  UnaryOp
	Arg Expr

	expr
}

func UnaryExprOf(op UnaryOp, arg Expr) UnaryExpr {
	return UnaryExpr{
		Op:  op,
		Arg: check.NotNil(arg).(Expr),
	}
}

func Not(arg Expr) Expr { return UnaryExprOf(NotOp, arg) }
