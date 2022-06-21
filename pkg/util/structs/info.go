package structs

import (
	"errors"
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

//

type FieldInfo struct {
	walkedField

	si *StructInfo
}

func (fi FieldInfo) GetValue(v any) (reflect.Value, bool) {
	rv, ok := rfl.UnwrapPointerValue(rfl.AsValue(v))
	if !ok {
		return reflect.Value{}, false
	}

	fv := rv
	for _, i := range fi.index {
		if fv.Kind() == reflect.Pointer {
			if fv.IsNil() {
				return reflect.Value{}, false
			}
			fv = fv.Elem()
		}
		fv = fv.Field(i)
	}

	return fv, true
}

//

type StructInfo struct {
	ty reflect.Type

	fields       []*FieldInfo
	fieldsByName map[string]*FieldInfo
}

func newStructInfo(ty reflect.Type) *StructInfo {
	if ty.Kind() != reflect.Struct {
		panic(errors.New("must be struct"))
	}

	si := &StructInfo{
		ty: ty,
	}

	wfs := walkFields(ty)
	fis := make([]*FieldInfo, len(wfs.list))
	bn := make(map[string]*FieldInfo)
	for i, wf := range wfs.list {
		fis[i] = &FieldInfo{
			walkedField: wf,
			si:          si,
		}
		if wf.name != "" {
			bn[wf.name] = fis[i]
		}
	}

	si.fields = fis
	si.fieldsByName = bn

	return si
}
