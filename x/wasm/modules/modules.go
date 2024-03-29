package modules

import (
	we "github.com/wrmsr/bane/x/wasm/exprs"
	wt "github.com/wrmsr/bane/x/wasm/types"
)

type Module struct {
	Funcs []Func
}

type Local struct {
	Name string
	Ty   wt.Type
}

type Func struct {
	Name string

	Res    wt.Type
	Params []wt.Type
	Vars   []wt.Type

	Body we.Expr
}
