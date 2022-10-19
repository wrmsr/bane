package wasm

import (
	"fmt"
	"strings"
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

func BuildExpr(root List) {
	k := root.ps[0].(Atom).s
	fmt.Println(k)

	var dot string
	if i := strings.Index(k, "."); i > 0 {
		dot = k[i+i:]
		k = k[:i]
		fmt.Println(dot)
	}

	switch k {
	case "block":
		//

	default:
		panic(k)
	}
}
