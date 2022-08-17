package inject

import (
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

type Injector struct {
	bs Bindings
	p  *Injector

	pfm providerFnMap
}

func NewInjector(bs Bindings) *Injector {
	return &Injector{
		bs: bs,

		pfm: makeProviderMap(bs).fns(),
	}
}

func (i *Injector) Bindings() Bindings { return i.bs }
func (i *Injector) Parent() *Injector  { return i.p }

func (i *Injector) NewChild(bs Bindings) *Injector {
	c := NewInjector(bs)
	c.p = i
	return c
}

func (i *Injector) TryProvide(o any) (any, bool) {
	k := AsKey(o)
	if p, ok := i.pfm[k]; ok {
		return p(i), true
	}
	return nil, false
}

func (i *Injector) Provide(o any) any {
	k := AsKey(o)
	if v, ok := i.TryProvide(k); ok {
		return v
	}
	panic(UnboundKeyError{Key: k})
}

func (i *Injector) ProvideArgs(fn any) []any {
	fnty := rfl.AsValue(fn).Type()
	ni := fnty.NumIn()
	as := make([]any, ni)
	for n := 0; n < ni; n++ {
		aty := fnty.In(n)
		ak := AsKey(aty)
		if v, ok := i.TryProvide(ak); ok {
			as[n] = v
			continue
		}
		panic(UnboundKeyError{Key: ak, Src: fn})
	}

	return as
}

func (i *Injector) Inject(fn any) []any {
	as := i.ProvideArgs(fn)
	avs := make([]reflect.Value, len(as))
	for n, a := range as {
		avs[n] = reflect.ValueOf(a)
	}

	rvs := rfl.AsValue(fn).Call(avs)
	rs := make([]any, len(rvs))
	for n, rv := range rvs {
		rs[n] = rv.Interface()
	}
	return rs
}

func (i *Injector) InjectOne(fn any) any {
	rs := i.Inject(fn)
	if len(rs) != 1 {
		panic(genericErrorf("expected 1 return value: %v", fn))
	}
	return rs[0]
}

func ProvideAs[T any](i *Injector) T {
	return i.Provide(KeyOf[T]()).(T)
}
