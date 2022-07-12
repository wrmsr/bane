package marshal

import (
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	stu "github.com/wrmsr/bane/pkg/util/structs"
)

//

type Manager struct {
	sic *stu.StructInfoCache

	reg *Registry

	mf MarshalerFactory
	uf UnmarshalerFactory
}

func NewManager() *Manager {
	sic := stu.NewStructInfoCache()

	mf := NewTypeCacheMarshalerFactory(NewCompositeMarshalerFactory(
		FirstComposite,
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

func (m *Manager) MarshalRfl(rv reflect.Value, o ...MarshalOpt) (Value, error) {
	ty := rv.Type()
	mc := MarshalContext{
		Make: m.mf.Make,
		Reg:  m.reg,
	}
	mi, err := m.mf.Make(mc, ty)
	if err != nil {
		return nil, err
	}
	return mi.Marshal(mc, rv)
}

func (m *Manager) Marshal(v any, o ...MarshalOpt) (Value, error) {
	return m.MarshalRfl(reflect.ValueOf(v), o...)
}

func (m *Manager) UnmarshalRfl(mv Value, ty reflect.Type, o ...UnmarshalOpt) (reflect.Value, error) {
	uc := UnmarshalContext{
		Make: m.uf.Make,
		Reg:  m.reg,
	}
	ui, err := m.uf.Make(uc, ty)
	if err != nil {
		return rfl.Invalid(), err
	}
	return ui.Unmarshal(uc, mv)
}

func (m *Manager) Unmarshal(mv Value, v any, o ...UnmarshalOpt) error {
	rv := reflect.ValueOf(v).Elem()
	uv, err := m.UnmarshalRfl(mv, rv.Type(), o...)
	if err != nil {
		return err
	}
	rv.Set(uv)
	return nil
}
