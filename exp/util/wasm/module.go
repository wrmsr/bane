package wasm

import (
	"fmt"
	"strings"

	"github.com/wrmsr/bane/pkg/util/check"
)

type Module struct{}

type Memory struct{}

type Func struct{}

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
			BuildFunc(l)

		default:
			panic(k)
		}
	}
}

func BuildFunc(root List) {
	if root.ps[0].(Atom).s != "func" {
		panic("invalid func")
	}
	name := root.ps[1].(Atom).s
	fmt.Println(name)

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

	for ; i < len(root.ps); i++ {
		l := root.ps[i].(List)

		BuildExpr(l)
	}
}

func BuildExpr(root List) Expr {
	k := root.ps[0].(Atom).s
	fmt.Println(k)

	var dot string
	if i := strings.Index(k, "."); i > 0 {
		dot = k[:i]
		k = k[i+1:]
		fmt.Println(dot)
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
		return Const{
			s:  root.ps[1].(Atom).s,
			ty: dot,
		}

	default:
		panic(k)
	}
}
