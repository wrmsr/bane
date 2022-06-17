package structs

import (
	"reflect"

	brfl "github.com/wrmsr/bane/pkg/utils/reflect"
)

//

type ToMapOpts struct{}

func (o ToMapOpts) Coalesce(opts ...ToMapOpts) ToMapOpts {
	return o
}

//

type ToMapContext struct {
	Struct reflect.Value
	Field  *FieldInfo
	Map    map[string]any
	Opts   ToMapOpts
}

type StructToMapper interface {
	StructToMap(tmc ToMapContext)
}

//

func (st *StructTool) ToMap(v any, opts ...ToMapOpts) map[string]any {
	o := ToMapOpts{}.Coalesce(opts...)

	rv, ok := brfl.UnwrapPointerValue(brfl.AsValue(v))
	if !ok {
		return nil
	}

	si := st.Info(rv.Type())

	m := make(map[string]any)
	for _, fi := range si.fields {
		if fi.name == "" {
			continue
		}

		frv, ok := fi.GetValue(v)
		if !ok {
			continue
		}
		fv := frv.Interface()

		if fv, ok := fv.(StructToMapper); ok {
			fv.StructToMap(ToMapContext{
				Struct: rv,
				Field:  fi,
				Map:    m,
				Opts:   o,
			})
			continue
		}

		if frv.Kind() == reflect.Struct {

		}

		m[fi.name] = fv
	}

	return m
}
