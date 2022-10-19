package wasm

type Module struct {
	funcs []Func
}

type Local struct {
	name string
	ty   Type
}

type Func struct {
	name string

	res    Type
	params []Type
	vars   []Type

	body Expr
}
