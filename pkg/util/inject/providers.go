package inject

import (
	"fmt"
	"reflect"
	"sync"
	"sync/atomic"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	rtu "github.com/wrmsr/bane/pkg/util/runtime"
)

//

type Provider interface {
	String() string

	providedTy(rec func(Key) reflect.Type) reflect.Type
	providerFn() providerFn
}

//

type providerMap map[Key]Provider

type providerFn = func(injector any) any
type providerFnMap map[Key]providerFn

func (pm providerMap) fns() providerFnMap {
	pfm := make(providerFnMap, len(pm))
	for k, p := range pm {
		pfm[k] = p.providerFn()
	}
	return pfm
}

//

func asProvider(o any) Provider {
	if _, ok := o.(Binding); ok {
		panic(genericErrorf("must not use bindings as providers"))
	}
	if _, ok := o.(Bindings); ok {
		panic(genericErrorf("must not use bindings as providers"))
	}

	if o == nil {
		return Const(nil)
	}

	if o, ok := o.(Provider); ok {
		return o
	}

	if o, ok := o.(Key); ok {
		return Link(o)
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

func (p constProvider) String() string {
	return fmt.Sprintf("Const{%v}", p.v)
}

func (p constProvider) providedTy(rec func(Key) reflect.Type) reflect.Type {
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

	fty reflect.Type
	ty  reflect.Type
}

func Func(fn any) Provider {
	rv := rfl.AsValue(fn)
	fty := rv.Type()
	if fty.Kind() != reflect.Func {
		panic(genericErrorf("must be func: %v", rv))
	}
	if fty.NumOut() != 1 {
		panic(genericErrorf("func must have one output: %v", fty))
	}
	return funcProvider{fn: rv, fty: fty, ty: fty.Out(0)}
}

var _ Provider = funcProvider{}

func (p funcProvider) String() string {
	return fmt.Sprintf("Func{%v, %s}", rtu.GetFuncName(p.fn), p.ty)
}

func (p funcProvider) providedTy(rec func(Key) reflect.Type) reflect.Type {
	return p.ty
}

func (p funcProvider) providerFn() providerFn {
	return func(i any) any {
		return i.(*Injector).InjectOne(p.fn)
	}
}

//

type linkProvider struct {
	k Key
}

func Link(k any) Provider {
	return linkProvider{k: AsKey(k)}
}

var _ Provider = linkProvider{}

func (p linkProvider) String() string {
	return fmt.Sprintf("Link{%s}", p.k)
}

func (p linkProvider) providedTy(rec func(Key) reflect.Type) reflect.Type {
	return rec(p.k)
}

func (p linkProvider) providerFn() providerFn {
	return func(i any) any {
		return i.(*Injector).Provide(p.k)
	}
}

//

type singletonProvider struct {
	p Provider
}

func Singleton(p any) Provider {
	return singletonProvider{p: asProvider(p)}
}

var _ Provider = singletonProvider{}

func (p singletonProvider) String() string {
	return fmt.Sprintf("Singleton{%s}", p.p)
}

func (p singletonProvider) providedTy(rec func(Key) reflect.Type) reflect.Type {
	return p.p.providedTy(rec)
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
