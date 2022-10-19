package wasm

import (
	"fmt"
	"strings"

	we "github.com/wrmsr/bane/exp/util/wasm/exprs"
	wp "github.com/wrmsr/bane/exp/util/wasm/parsing"
	wt "github.com/wrmsr/bane/exp/util/wasm/types"
	"github.com/wrmsr/bane/pkg/util/check"
)

func BuildModule(root wp.List) {
	if root.Ps[0].(wp.Atom).S != "module" {
		panic("invalid module")
	}

	for _, c := range root.Ps[1:] {
		l := c.(wp.List)
		switch k := l.Ps[0].(wp.Atom).S; k {

		case "memory":
			//

		case "func":
			fn := BuildFunc(l)
			fmt.Printf("%#v\n", fn)

		case "type":
			//

		case "import":
			//

		case "global":
			//

		case "data":
			//

		case "table":
			//

		case "elem":
			//

		case "export":
			//

		default:
			panic(k)
		}
	}
}

func BuildFunc(root wp.List) Func {
	if root.Ps[0].(wp.Atom).S != "func" {
		panic("invalid func")
	}
	name := root.Ps[1].(wp.Atom).S

	i := 2
l:
	for ; i < len(root.Ps); i++ {
		l := root.Ps[i].(wp.List)
		switch k := l.Ps[0].(wp.Atom).S; k {

		case "param":
			//

		case "local":
			//

		case "result":
			//

		case "type":
			//

		default:
			break l
		}
	}

	var bes []we.Expr
	for ; i < len(root.Ps); i++ {
		l := root.Ps[i].(wp.List)

		be := BuildExpr(l)
		bes = append(bes, be)
	}

	var body we.Expr
	if len(bes) > 1 {
		body = we.Block{
			S: bes,
		}
	} else if len(bes) < 1 {
		body = we.Nop{}
	} else {
		body = bes[0]
	}

	return Func{
		name: name,

		body: body,
	}
}

func BuildExpr(root wp.List) we.Expr {
	k := root.Ps[0].(wp.Atom).S

	var dot string
	if i := strings.Index(k, "."); i > 0 {
		dot = k[:i]
		k = k[i+1:]
	}

	switch k {

	case "block":
		id := root.Ps[1].(wp.Atom).S
		var s []we.Expr
		for i := 2; i < len(root.Ps); i++ {
			l := root.Ps[i].(wp.List)
			s = append(s, BuildExpr(l))
		}
		return we.Block{
			Id: id,
			S:  s,
		}

	case "const":
		check.Equal(len(root.Ps), 2)
		ty := wt.ParseType(dot)
		return we.Const{
			S:  root.Ps[1].(wp.Atom).S,
			Ty: ty,
		}

	case "store":
		return we.Nop{}

	case "set":
		return we.Nop{}

	case "drop":
		return we.Nop{}

	case "store8":
		return we.Nop{}

	case "get":
		return we.Nop{}

	case "set_local":
		check.Equal(len(root.Ps), 3)
		n := root.Ps[1].(wp.Atom).S
		v := BuildExpr(root.Ps[2].(wp.List))
		return we.SetLocal{
			N: n,
			V: v,
		}

	//

	case "select":
		c := BuildExpr(root.Ps[1].(wp.List))
		t := BuildExpr(root.Ps[2].(wp.List))
		var f we.Expr
		if len(root.Ps) > 3 {
			f = BuildExpr(root.Ps[3].(wp.List))
		}
		check.Condition(len(root.Ps) < 5)
		return we.Select{
			C: c,
			T: t,
			F: f,
		}

	case "call":
		return we.Nop{}

	case "br":
		return we.Nop{}

	case "br_table":
		return we.Nop{}

	case "br_if":
		return we.Nop{}

	case "loop":
		return we.Nop{}

	case "return":
		return we.Nop{}

	case "call_indirect":
		return we.Nop{}

	//

	case "reinterpret_i64":
		return we.Nop{}

	case "reinterpret_f64":
		return we.Nop{}

	//

	case "add":
		return we.Nop{}

	case "sub":
		return we.Nop{}

	case "shl":
		return we.Nop{}

	case "lt_u":
		return we.Nop{}

	//

	default:
		panic(k)
	}
}
