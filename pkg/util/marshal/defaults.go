package marshal

import (
	stu "github.com/wrmsr/bane/pkg/util/structs"
)

func NewDefaultManager(reg *Registry) *Manager {
	sic := stu.NewStructInfoCache()

	var mfs = []MarshalerFactory{
		NewRegistryMarshalerFactory(),
		NewPrimitiveMarshalerFactory(),
		NewConvertUserPrimitiveMarshalerFactory(),
		NewPointerMarshalerFactory(),
		NewIndexMarshalerFactory(),
		NewMapMarshalerFactory(),
		NewBase64MarshalerFactory(),
		NewTimeMarshalerFactory(DefaultTimeMarshalLayout),
		NewOptionalMarshalerFactory(),
		NewRegistryPolymorphismMarshalerFactory(),
		NewStructMarshalerFactory(sic),
	}

	mf := NewTypeCacheMarshalerFactory(
		NewRecursiveMarshalerFactory(
			NewCompositeMarshalerFactory(
				FirstComposite,
				mfs...,
			)))

	var ufs = []UnmarshalerFactory{
		NewRegistryUnmarshalerFactory(),
		NewConvertPrimitiveUnmarshalerFactory(),
		NewPointerUnmarshalerFactory(),
		NewIndexUnmarshalerFactory(),
		NewMapUnmarshalerFactory(),
		NewBase64UnmarshalerFactory(),
		NewTimeUnmarshalerFactory(DefaultTimeUnmarshalLayouts()),
		NewOptionalUnmarshalerFactory(),
		NewRegistryPolymorphismUnmarshalerFactory(),
		NewStructUnmarshalerFactory(sic),
	}

	uf := NewTypeCacheUnmarshalerFactory(
		NewRecursiveUnmarshalerFactory(
			NewCompositeUnmarshalerFactory(
				FirstComposite,
				ufs...,
			)))

	return &Manager{
		sic: sic,

		reg: reg,

		mf: mf,
		uf: uf,
	}
}
