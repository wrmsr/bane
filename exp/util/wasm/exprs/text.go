package exprs

import (
	"github.com/wrmsr/bane/exp/util/wasm/text"
)

func texts(es ...Expr) text.List {
	var l text.List
	for _, e := range es {
		l.Append(e.Text())
	}
	return l
}

func (e Block) Text() text.Element {
	l := text.L("block")
	if e.Id != "" {
		l.Append(e.Id)
	}
	for _, c := range e.S {
		l.Append(c.Text())
	}
	return l
}

func (e Const) Text() text.Element {
	return text.L(
		"const",
		e.Ty.String(),
		text.Q(e.S),
	)
}

func (e If) Text() text.Element {
	l := text.L(
		"if",
		texts(e.Then...),
	)
	if len(e.Else) > 0 {
		l.Append(texts(e.Else...))
	}
	return l
}

func (e Nop) Text() text.Element {
	return text.L("nop")
}

func (e Select) Text() text.Element {
	l := text.L(
		"select",
		e.C.Text(),
		e.T.Text(),
	)
	if e.F != nil {
		l.Append(e.F.Text())
	}
	return l
}

func (e SetLocal) Text() text.Element {
	return text.L(
		"set_local",
		e.N, e.V.Text(),
	)
}

func (e Unreachable) Text() text.Element {
	return text.L("unreachable")
}
