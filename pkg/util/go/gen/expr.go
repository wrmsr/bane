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

func NewAddr(value Expr) Addr {
	return Addr{
		Value: value,
	}
}

//

type Call struct {
	Func Expr
	Args []Expr

	expr
}

func NewCall(func_ Expr, args ...Expr) Call {
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

func NewDeref(value Expr) Deref {
	return Deref{
		Value: value,
	}
}

//

type FuncExpr struct {
	Func Func

	expr
}

func NewFuncExpr(func_ Func) FuncExpr {
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

func NewIndex(value, index Expr) Index {
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

func NewInfixExpr(op InfixOp, args ...Expr) InfixExpr {
	return InfixExpr{
		Op:   op,
		Args: check.NotEmptySlice(args),
	}
}

func NewInfixExprOrSelf(op InfixOp, args ...Expr) Expr {
	if len(args) == 1 {
		return args[0]
	}
	return NewInfixExpr(
		op,
		args...,
	)
}

func Add(args ...Expr) Expr { return NewInfixExprOrSelf(AddOp, args...) }
func Sub(args ...Expr) Expr { return NewInfixExprOrSelf(SubOp, args...) }
func Mul(args ...Expr) Expr { return NewInfixExprOrSelf(MulOp, args...) }
func Div(args ...Expr) Expr { return NewInfixExprOrSelf(DivOp, args...) }
func Mod(args ...Expr) Expr { return NewInfixExprOrSelf(ModOp, args...) }

func Lt(args ...Expr) Expr  { return NewInfixExprOrSelf(LtOp, args...) }
func Lte(args ...Expr) Expr { return NewInfixExprOrSelf(LteOp, args...) }
func Eq(args ...Expr) Expr  { return NewInfixExprOrSelf(EqOp, args...) }
func Ne(args ...Expr) Expr  { return NewInfixExprOrSelf(NeOp, args...) }
func Gt(args ...Expr) Expr  { return NewInfixExprOrSelf(GtOp, args...) }
func Gte(args ...Expr) Expr { return NewInfixExprOrSelf(GteOp, args...) }

func And(args ...Expr) Expr { return NewInfixExprOrSelf(AndOp, args...) }
func Or(args ...Expr) Expr  { return NewInfixExprOrSelf(OrOp, args...) }

//

type Lit struct {
	String string

	expr
}

func NewLit(s string) Lit {
	return Lit{
		String: check.NotZero(s),
	}
}

//

type Paren struct {
	Value Expr

	expr
}

func NewParen(value Expr) Paren {
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

func NewSelect(value Expr, names ...Ident) Select {
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

func NewTypeAssert(value Expr, type_ Type) TypeAssert {
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

func NewUnaryExpr(op UnaryOp, arg Expr) UnaryExpr {
	return UnaryExpr{
		Op:  op,
		Arg: check.NotNil(arg).(Expr),
	}
}

func Not(arg Expr) Expr { return NewUnaryExpr(NotOp, arg) }
