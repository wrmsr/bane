package structs

import (
	"errors"
	"reflect"
)

type StructInfo struct {
	name Name
	ty   reflect.Type

	fields StructFields
}

func (si *StructInfo) Name() Name            { return si.name }
func (si *StructInfo) Type() reflect.Type    { return si.ty }
func (si *StructInfo) Fields() *StructFields { return &si.fields }

func buildStructInfo(
	ty reflect.Type,
) *StructInfo {
	if ty.Kind() != reflect.Struct {
		panic(errors.New("must be struct"))
	}

	rootFields := buildFieldInfo(
		ty,
		nil,
		"",
	)

	si := &StructInfo{
		name: buildName(ty.Name()),
		ty:   ty,

		fields: buildStructFields(rootFields),
	}

	for _, fi := range si.fields.all {
		fi.si = si
	}

	return si
}
