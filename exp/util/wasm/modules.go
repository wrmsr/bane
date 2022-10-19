package wasm

import (
	we "github.com/wrmsr/bane/exp/util/wasm/exprs"
	wt "github.com/wrmsr/bane/exp/util/wasm/types"
)

type Module struct {
	funcs []Func
}

type Local struct {
	name string
	ty   wt.Type
}

type Func struct {
	name string

	res    wt.Type
	params []wt.Type
	vars   []wt.Type

	body we.Expr
}
