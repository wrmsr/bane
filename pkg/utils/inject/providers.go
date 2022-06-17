package inject

import "reflect"

type providerFn = func(injector any) any
type providerMap map[Key]providerFn

type Provider interface {
	providedTy() reflect.Type
	providerFn() providerFn
}
