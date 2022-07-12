package inject

import (
	"fmt"
	"reflect"

	its "github.com/wrmsr/bane/pkg/util/iterators"
)

//

type Binding struct {
	key      Key
	provider Provider
}

func asBinding(o any) Binding {
	if o == nil {
		panic(genericErrorf("must explicitly bind nil"))
	}

	if o, ok := o.(Binding); ok {
		return o
	}

	if o, ok := o.(Provider); ok {
		return Binding{key: Key{ty: o.providedTy(func(Key) reflect.Type {
			panic(genericError(fmt.Errorf("can't determine binding type: %s", o)))
		})}, provider: o}
	}

	rv := reflect.ValueOf(o)
	if rv.Kind() == reflect.Func {
		return asBinding(Func(rv))
	}

	return Binding{key: Key{ty: rv.Type()}, provider: Const(o)}
}

func asBindings(os []any) []Binding {
	bs := make([]Binding, len(os))
	for i, o := range os {
		bs[i] = asBinding(o)
	}
	return bs
}

func As(k, p any) Binding {
	return Binding{AsKey(k), asProvider(p)}
}

//

type Binder = func() Bindings

type Bindings interface {
	its.Traversable[Binding]
	isBindings()
}

//

type bindings struct {
	bs []Binding
	ps []Bindings
}

var _ Bindings = bindings{}

func (bs bindings) isBindings() {}

func (bs bindings) ForEach(fn func(Binding) bool) bool {
	for _, p := range bs.ps {
		if !p.ForEach(fn) {
			return false
		}
	}
	for _, b := range bs.bs {
		if !fn(b) {
			return false
		}
	}
	return true
}

func Bind(os ...any) Bindings {
	return &bindings{
		bs: asBindings(os),
	}
}

func Append(ps ...Bindings) Bindings {
	var rps []Bindings
	for _, p := range ps {
		if p != nil {
			rps = append(rps, p)
		}
	}
	return &bindings{
		ps: rps,
	}
}

//

type overrides struct {
	p Bindings
	m map[Key]Binding
}

func Override(p Bindings, a ...any) Bindings {
	m := make(map[Key]Binding)
	Bind(a...).ForEach(func(b Binding) bool {
		if _, ok := m[b.key]; ok {
			panic(DuplicateBindingError{b.key})
		}
		m[b.key] = b
		return true
	})
	return overrides{
		p: p,
		m: m,
	}
}

var _ Bindings = overrides{}

func (bs overrides) isBindings() {}

func (bs overrides) ForEach(fn func(Binding) bool) bool {
	return bs.p.ForEach(func(b Binding) bool {
		if ob, ok := bs.m[b.key]; ok {
			b = ob
		}
		return fn(b)
	})
}

//

func makeProviderMap(bs Bindings) providerMap {
	m := make(providerMap)
	bs.ForEach(func(b Binding) bool {
		if _, ok := m[b.key]; ok {
			panic(DuplicateBindingError{b.key})
		}
		m[b.key] = b.provider.providerFn()
		return true
	})
	return m
}
