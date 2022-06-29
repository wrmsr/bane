package marshal

import (
	"reflect"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	stu "github.com/wrmsr/bane/pkg/util/structs"
)

//

type MarshalOpt interface {
	isMarshalOpt()
}

type BaseMarshalOpt struct{}

func (o BaseMarshalOpt) isMarshalOpt() {}

//

type MarshalContext struct {
	Opts ctr.TypeMap[MarshalOpt]
}

type Marshaller interface {
	Marshal(ctx MarshalContext, v reflect.Value)
}

//

type Marshalling struct {
	sic *stu.StructInfoCache

	mm map[reflect.Type]Marshaller
}

func NewMarshalling() *Marshalling {
	return &Marshalling{
		sic: stu.NewStructInfoCache(),

		mm: make(map[reflect.Type]Marshaller),
	}
}

func (m *Marshalling) Marshal(v any, o ...MarshalOpt) any {
	//opts := ctr.NewTypeMap[MarshalOpt](its.Of(o...))

	rv, ok := rfl.UnwrapPointerValue(rfl.AsValue(v))
	if !ok {
		return nil
	}

	si := m.sic.Info(rv.Type())

	r := make(map[string]any)
	for _, fi := range si.Fields() {
		if fi.Name() == "" {
			continue
		}

		frv, ok := fi.GetValue(v)
		if !ok {
			continue
		}
		fv := frv.Interface()

		if frv.Kind() == reflect.Struct {

		}

		r[fi.Name()] = fv
	}

	return r
}
