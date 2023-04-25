package consteval

import (
	"fmt"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type flow int8

const (
	noFlow flow = iota
	returnFlow
	breakFlow
	gotoFlow
	panicFlow
)

func (f flow) String() string {
	switch f {
	case noFlow:
		return "no"
	case returnFlow:
		return "return"
	case breakFlow:
		return "break"
	case gotoFlow:
		return "goto"
	case panicFlow:
		return "panic"
	default:
		panic(f)
	}
}

//

type effect struct {
	bt.Result[Value]
	flow flow
}

func valEffect(v Value) effect {
	return effect{Result: bt.Ok(v)}
}

func errEffect(err error) effect {
	return effect{Result: bt.Err[Value](err)}
}

func (e effect) expr() effect {
	if e.Err == nil {
		if e.flow != noFlow {
			e.Err = fmt.Errorf("expr has flow: %#v", e)
		} else if e.Val == nil && e.Err != nil {
			e.Err = fmt.Errorf("empty expr: %#v", e)
		}
	}
	return e
}

func (e effect) mustVal() Value {
	if e.Val == nil {
		panic(fmt.Errorf("must have val: %#v", e))
	}
	return e.Val
}
