package lisp

type Scope struct {
	parent *Scope
	defs   map[string]Value
}
