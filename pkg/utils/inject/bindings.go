package inject

import (
	"reflect"

	its "github.com/wrmsr/bane/pkg/utils/iterators"
)

//

type Binding struct {
	key      Key
	provider Provider
}

func toBinding(o any) Binding {
	if o == nil {
		panic(genericErrorf("must explicitly bind nil"))
	}

	if o, ok := o.(Binding); ok {
		return o
	}

	if o, ok := o.(Provider); ok {
		return Binding{key: Key{ty: o.providedTy()}, provider: o}
	}

	rv := reflect.ValueOf(o)
	if rv.Kind() == reflect.Func {
		return toBinding(Func(rv))
	}

	return Binding{key: Key{ty: rv.Type()}, provider: Const(o)}
}

func toBindings(os []any) []Binding {
	bs := make([]Binding, len(os))
	for i, o := range os {
		bs[i] = toBinding(o)
	}
	return bs
}

//

type Bindings interface {
	its.CanForEach[Binding]
}

type bindings struct {
	bs []Binding
	ps []Bindings
}

var _ Bindings = bindings{}

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
		bs: toBindings(os),
	}
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
