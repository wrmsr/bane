package marshal

import (
	"reflect"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	its "github.com/wrmsr/bane/pkg/util/iterators"
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

type MarshalStructContext struct {
	Struct reflect.Value
	Field  *stu.FieldInfo
	Map    map[string]any
	Opts   ctr.TypeMap[MarshalOpt]
}

type StructMarshaller interface {
	ExplodeStruct(ctx MarshalStructContext)
}

//

type Marshalling struct {
	sic *stu.StructInfoCache
}

func (m *Marshalling) Explode(v any, o ...MarshalOpt) map[string]any {
	opts := ctr.NewTypeMap[MarshalOpt](its.Of(o...))

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

		if fv, ok := fv.(StructMarshaller); ok {
			fv.ExplodeStruct(MarshalStructContext{
				Struct: rv,
				Field:  fi,
				Map:    r,
				Opts:   opts,
			})
			continue
		}

		if frv.Kind() == reflect.Struct {

		}

		r[fi.Name()] = fv
	}

	return r
}
