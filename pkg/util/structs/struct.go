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

func (si *StructInfo) GetPath(name string) *FieldInfo { return si.fields.byPath.Get(name) }
func (si *StructInfo) GetFlat(name string) *FieldInfo { return si.fields.byFlat.Get(name) }

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

	si.fields.all.ForEach(func(fi *FieldInfo) bool {
		fi.si = si
		return true
	})

	return si
}
