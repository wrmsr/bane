package inject

import (
	"reflect"
	"sync"
	"sync/atomic"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

//

type providerFn = func(injector any) any
type providerMap map[Key]providerFn

type Provider interface {
	providedTy() reflect.Type
	providerFn() providerFn
}

//

func toProvider(o any) Provider {
	if _, ok := o.(Binding); ok {
		panic(genericErrorf("must not use bindings as providers"))
	}
	if _, ok := o.(Bindings); ok {
		panic(genericErrorf("must not use bindings as providers"))
	}

	if o == nil {
		return Const(nil)
	}

	rv := reflect.ValueOf(o)
	rt := rv.Type()
	if rt.Kind() == reflect.Func {
		return Func(o)
	}

	return Const(o)
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

//

type funcProvider struct {
	fn reflect.Value
	ty reflect.Type
}

func Func(fn any) Provider {
	rv := rfl.AsValue(fn)
	ty := rv.Type()
	if ty.Kind() != reflect.Func {
		panic(genericErrorf("must be func: %v", rv))
	}
	if ty.NumOut() != 1 {
		panic(genericErrorf("func must have one output: %v", ty))
	}
	return funcProvider{fn: rv, ty: ty}
}

var _ Provider = funcProvider{}

func (p funcProvider) providedTy() reflect.Type {
	return p.ty.Out(0)
}

func (p funcProvider) providerFn() providerFn {
	return func(i any) any {
		return i.(*Injector).InjectOne(p.fn)
	}
}

//

type singletonProvider struct {
	p Provider
}

func Singleton(p any) Provider {
	return singletonProvider{p: toProvider(p)}
}

var _ Provider = singletonProvider{}

func (p singletonProvider) providedTy() reflect.Type {
	return p.p.providedTy()
}

func (p singletonProvider) providerFn() providerFn {
	var o sync.Once
	var v atomic.Value
	type box struct{ v any }
	fn := p.p.providerFn()
	return func(i any) any {
		o.Do(func() {
			v.Store(box{fn(i)})
		})
		return v.Load().(box).v
	}
}
