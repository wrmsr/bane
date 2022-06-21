package inject

import eu "github.com/wrmsr/bane/pkg/util/errors"

func NewDeferInjector(bs Bindings) (*Injector, func() error) {
	inj := NewInjector(
		Append(
			Bind(Singleton(eu.NewDeferStack)),
			bs,
		),
	)
	ds := Provide[*eu.DeferStack](inj)
	return inj, ds.Call
}
