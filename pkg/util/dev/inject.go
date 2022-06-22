//go:build !nodev

package dev

import (
	"errors"
	"sync"

	au "github.com/wrmsr/bane/pkg/util/atomic"
	inj "github.com/wrmsr/bane/pkg/util/inject"
)

//

var binders = inj.NewBinderRegistry()

func Binders() *inj.BinderRegistry {
	return binders
}

func Register(bs ...inj.Binder) *inj.BinderRegistry { return binders.Register(bs...) }
func Bind() inj.Bindings                            { return binders.Bind() }

//

type injectorAndCloser struct {
	injector *inj.Injector
	closer   func() error
}

type devInjector struct {
	mtx sync.Mutex
	val au.Value[*injectorAndCloser]
}

func (i *devInjector) setup(bs inj.Bindings) (*inj.Injector, func() error) {
	i.mtx.Lock()
	defer i.mtx.Unlock()

	if i.val.Load() != nil {
		panic(errors.New("dev injector already setup"))
	}

	injector, closer := inj.NewDeferInjector(bs)
	i.val.Store(&injectorAndCloser{
		injector: injector,
		closer:   closer,
	})
	return injector, closer
}

func (i *devInjector) get() *inj.Injector {
	if val := i.val.Load(); val != nil {
		return val.injector
	}
	panic(errors.New("dev injector not setup"))
}

var _injector devInjector

func SetupInjector() (*inj.Injector, func() error) {
	return _injector.setup(Bind())
}

func Injector() *inj.Injector {
	return _injector.get()
}

//

func TryProvide(o any) (any, bool) { return Injector().TryProvide(o) }
func Provide(o any) any            { return Injector().Provide(o) }
func ProvideArgs(fn any) []any     { return Injector().ProvideArgs(fn) }
func Inject(fn any) []any          { return Injector().Inject(fn) }
func InjectOne(fn any) any         { return Injector().InjectOne(fn) }
