package lisp

import "fmt"

type Scope struct {
	parent *Scope
	defs   map[string]Value
}

func NewScope(parent *Scope) *Scope {
	return &Scope{
		parent: parent,
		defs:   make(map[string]Value),
	}
}

func (sc *Scope) Get(key string) (v Value, ok bool) {
	for p := sc; !ok && p != nil; p = p.parent {
		v, ok = p.defs[key]
	}
	return
}

func (sc *Scope) Set(key string, val Value) {
	sc.defs[key] = val
}

func (sc *Scope) Merge(proc *Proc, vals []Value) {
	argv := len(vals)
	argc := len(proc.Args)

	if argv != argc {
		panic(fmt.Errorf("proc %s takes %d arguments, got %d", proc.Name, argc, argv))
	}

	for i, v := range proc.Args {
		sc.Set(v, vals[i])
	}
}

func (sc *Scope) Derive(proc *Proc, vals []Value) (ret *Scope) {
	ret = new(Scope)
	ret.parent = sc
	ret.defs = make(map[string]Value, len(proc.Args))
	ret.Merge(proc, vals)
	return
}

type valueRef struct {
	refs *Scope
	name string
}

func (r valueRef) update(v Value) {
	if r.refs == nil {
		panic("undefined reference: " + r.name)
	} else {
		r.refs.defs[r.name] = v
	}
}

func (sc *Scope) resolve(name string) valueRef {
	var ok bool
	var nsc *Scope

	for nsc = sc; nsc != nil; nsc = nsc.parent {
		if _, ok = nsc.defs[name]; ok {
			break
		}
	}

	return valueRef{
		refs: nsc,
		name: name,
	}
}
