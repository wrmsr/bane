package dev

import (
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

type devInjector struct {
	mtx  sync.Mutex
	once sync.Once

	injector *inj.Injector
	closer   func() error
}

func (i *devInjector) get() *inj.Injector {

}

var (
	_injector = au.NewLazy(func() *inj.Injector {
		return inj.New

	})
)

func Injector() *inj.Injector {
	return _injector.Get()
}

//

func TryProvide(o any) (any, bool) { return Injector().TryProvide(o) }
func Provide(o any) any            { return Injector().Provide(o) }
func ProvideArgs(fn any) []any     { return Injector().ProvideArgs(fn) }
func Inject(fn any) []any          { return Injector().Inject(fn) }
func InjectOne(fn any) any         { return Injector().InjectOne(fn) }
