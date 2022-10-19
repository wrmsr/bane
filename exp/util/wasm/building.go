package wasm

import (
	"fmt"
	"strings"

	"github.com/wrmsr/bane/pkg/util/check"
)

func BuildModule(root List) {
	if root.ps[0].(Atom).s != "module" {
		panic("invalid module")
	}

	for _, c := range root.ps[1:] {
		l := c.(List)
		switch k := l.ps[0].(Atom).s; k {

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

func BuildFunc(root List) Func {
	if root.ps[0].(Atom).s != "func" {
		panic("invalid func")
	}
	name := root.ps[1].(Atom).s

	i := 2
l:
	for ; i < len(root.ps); i++ {
		l := root.ps[i].(List)
		switch k := l.ps[0].(Atom).s; k {

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

	var bes []Expr
	for ; i < len(root.ps); i++ {
		l := root.ps[i].(List)

		be := BuildExpr(l)
		bes = append(bes, be)
	}

	var body Expr
	if len(bes) > 1 {
		body = Block{
			s: bes,
		}
	} else if len(bes) < 1 {
		body = Nop{}
	} else {
		body = bes[0]
	}

	return Func{
		name: name,

		body: body,
	}
}

func BuildExpr(root List) Expr {
	k := root.ps[0].(Atom).s

	var dot string
	if i := strings.Index(k, "."); i > 0 {
		dot = k[:i]
		k = k[i+1:]
	}

	switch k {
	case "block":
		id := root.ps[1].(Atom).s
		var s []Expr
		for i := 2; i < len(root.ps); i++ {
			l := root.ps[i].(List)
			s = append(s, BuildExpr(l))
		}
		return Block{
			id: id,
			s:  s,
		}

	case "set_local":
		check.Equal(len(root.ps), 3)
		n := root.ps[1].(Atom).s
		v := BuildExpr(root.ps[2].(List))
		return SetLocal{
			n: n,
			v: v,
		}

	case "const":
		check.Equal(len(root.ps), 2)
		ty := ParseType(dot)
		return Const{
			s:  root.ps[1].(Atom).s,
			ty: ty,
		}

	case "call":
		return Nop{}

	case "store":
		return Nop{}

	case "set":
		return Nop{}

	case "br":
		return Nop{}

	case "br_table":
		return Nop{}

	case "br_if":
		return Nop{}

	case "drop":
		return Nop{}

	case "loop":
		return Nop{}

	case "store8":
		return Nop{}

	case "get":
		return Nop{}

	case "lt_u":
		return Nop{}

	case "select":
		return Nop{}

	case "return":
		return Nop{}

	case "call_indirect":
		return Nop{}

	case "reinterpret_i64":
		return Nop{}

	case "reinterpret_f64":
		return Nop{}

	case "shl":
		return Nop{}

	case "sub":
		return Nop{}

	case "add":
		return Nop{}

	default:
		panic(k)
	}
}
