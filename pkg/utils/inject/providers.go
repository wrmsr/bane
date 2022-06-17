package inject

import "reflect"

//

type providerFn = func(injector any) any
type providerMap map[Key]providerFn

type Provider interface {
	providedTy() reflect.Type
	providerFn() providerFn
}

//

type constProvider struct {
	v any
}

func Const(v any) Provider {
	return constProvider{v}
}

var _ Provider = constProvider{}

func (p constProvider) providedTy() reflect.Type {
	return reflect.TypeOf(p.v)
}

func (p constProvider) providerFn() providerFn {
	return func(_ any) any {
		return p.v
	}
}
