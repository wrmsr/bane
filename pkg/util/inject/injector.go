package inject

import (
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

type Injector struct {
	bs Bindings
	pm providerMap
}

func NewInjector(bs Bindings) *Injector {
	return &Injector{
		bs: bs,
		pm: makeProviderMap(bs),
	}
}

func (i *Injector) TryProvide(o any) (any, bool) {
	k := AsKey(o)
	if p, ok := i.pm[k]; ok {
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
