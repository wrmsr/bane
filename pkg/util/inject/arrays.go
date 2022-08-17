package inject

import "reflect"

//

type arrayProvider struct {
	ty reflect.Type
	ps []Provider

	sty reflect.Type
}

func newArrayProvider(ty reflect.Type, ps []Provider) arrayProvider {
	return arrayProvider{
		ty: ty,
		ps: ps,

		sty: reflect.SliceOf(ty),
	}
}

var _ Provider = arrayProvider{}

func (p arrayProvider) providedTy(rec func(Key) reflect.Type) reflect.Type {
	return p.sty
}

func (p arrayProvider) providerFn() providerFn {
	l := len(p.ps)
	ps := make([]providerFn, l, l)
	for n, p := range p.ps {
		ps[n] = p.providerFn()
	}
	return func(i any) any {
		rv := reflect.MakeSlice(p.sty, l, l)
		for n, c := range ps {
			rv.Index(n).Set(reflect.ValueOf(c(i)))
		}
		return rv.Interface()
	}
}

//

// TODO: BindEmptyArrayOf
