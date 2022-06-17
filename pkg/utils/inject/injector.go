package inject

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

func (i *Injector) ProvideArgs(fn any) []any {
	return nil
}
