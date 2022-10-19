package wasm

func BuildModule(root Element) {
	l := root.(List)
	if l.ps[0].(Atom).s != "module" {
		panic("invalid module")
	}
}
