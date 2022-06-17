package inject

import "reflect"

type Injector struct {
	bs Bindings
	pm providerMap
}

func NewInjector(bs Bindings) *Injector {
	pm := make(providerMap)
	for _, b := range bs.bs {
		if _, ok := pm[b.key]; ok {
			panic(DuplicateBindingError{b.key})
		}
		pm[b.key] = b.provider.providerFn()
	}

	return &Injector{
		bs: bs,
		pm: pm,
	}
}

func (i *Injector) TryProvide(o any) (any, bool) {
	k := ToKey(o)
	if p, ok := i.pm[k]; ok {
		return p(i), true
	}
	return nil, false
}

func (i *Injector) Provide(o any) any {
	k := ToKey(o)
	if v, ok := i.TryProvide(k); ok {
		return v
	}
	panic(UnboundKeyError{Key: k})
}

func (i *Injector) ProvideArgs(fn any) []any {
	fnty := reflect.TypeOf(fn)

	ni := fnty.NumIn()
	as := make([]any, ni)
	for n := 0; n < ni; n++ {
		aty := fnty.In(n)
		ak := ToKey(aty)
		if v, ok := i.TryProvide(ak); ok {
			as[n] = v
			continue
		}
		panic(UnboundKeyError{Key: ak, Src: fn})
	}

	return as
}
