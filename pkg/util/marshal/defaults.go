package marshal

import (
	stu "github.com/wrmsr/bane/pkg/util/structs"
)

func NewDefaultManager() *Manager {
	sic := stu.NewStructInfoCache()

	mf := NewTypeCacheMarshalerFactory(NewCompositeMarshalerFactory(
		FirstComposite,
		NewRegistryMarshalerFactory(),
		NewPrimitiveMarshalerFactory(),
		NewPointerMarshalerFactory(),
		NewIndexMarshalerFactory(),
		NewMapMarshalerFactory(),
		NewBase64MarshalerFactory(),
		NewTimeMarshalerFactory(DefaultTimeMarshalLayout),
		NewOptionalMarshalerFactory(),
		NewRegistryPolymorphismMarshalerFactory(),
		NewStructMarshalerFactory(sic),
	))

	uf := NewTypeCacheUnmarshalerFactory(NewCompositeUnmarshalerFactory(
		FirstComposite,
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
	))

	return &Manager{
		sic: sic,

		reg: globalRegistry,

		mf: mf,
		uf: uf,
	}
}
