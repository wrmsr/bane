package explode

import (
	"reflect"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	stu "github.com/wrmsr/bane/pkg/util/structs"
)

//

type ExplodeOpt interface {
	isExplodeOpt()
}

type BaseExplodeOpt struct{}

func (o BaseExplodeOpt) isExplodeOpt() {}

//

type ExplodeStructContext struct {
	Struct reflect.Value
	Field  *stu.FieldInfo
	Map    map[string]any
	Opts   ctr.TypeMap[ExplodeOpt]
}

type StructExploder interface {
	ExplodeStruct(ctx ExplodeStructContext)
}

//

type ExplosionManager struct {
	sic *stu.StructInfoCache
}

func (em *ExplosionManager) Explode(v any, o ...ExplodeOpt) map[string]any {
	opts := ctr.NewTypeMap[ExplodeOpt](its.Of(o...))

	rv, ok := rfl.UnwrapPointerValue(rfl.AsValue(v))
	if !ok {
		return nil
	}

	si := em.sic.Info(rv.Type())

	m := make(map[string]any)
	for _, fi := range si.Fields() {
		if fi.Name() == "" {
			continue
		}

		frv, ok := fi.GetValue(v)
		if !ok {
			continue
		}
		fv := frv.Interface()

		if fv, ok := fv.(StructExploder); ok {
			fv.ExplodeStruct(ExplodeStructContext{
				Struct: rv,
				Field:  fi,
				Map:    m,
				Opts:   opts,
			})
			continue
		}

		if frv.Kind() == reflect.Struct {

		}

		m[fi.Name()] = fv
	}

	return m
}
