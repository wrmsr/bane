package jmespath

type Function struct {
	Name string
	Fn   func([]Arg[any]) any
}

var Functions = []Function{
	{"sum", func(arg []Arg[any]) any {
		panic("fixme")
	}},
}
