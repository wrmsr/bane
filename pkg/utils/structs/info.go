package structs

import (
	"errors"
	"reflect"
)

//

type FieldInfo struct {
	walkedField

	si *StructInfo
}

//

type StructInfo struct {
	ty reflect.Type

	fields       []FieldInfo
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
	fis := make([]FieldInfo, len(wfs.list))
	bn := make(map[string]*FieldInfo)
	for i, wf := range wfs.list {
		fis[i] = FieldInfo{
			walkedField: wf,
			si:          si,
		}
		if wf.name != "" {
			bn[wf.name] = &fis[i]
		}
	}

	si.fields = fis
	si.fieldsByName = bn

	return si
}
