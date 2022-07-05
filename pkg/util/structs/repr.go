package structs

import (
	"reflect"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	"github.com/wrmsr/bane/pkg/util/slices"
)

func structFieldRepr(field reflect.StructField) ctr.Map[string, any] {
	return ctr.NewMapBuilder[string, any]().
		Put("pkg_path", field.PkgPath).
		Put("type", field.Type.Name()).
		Put("tag", string(field.Tag)).
		Put("offset", field.Offset).
		Put("index", field.Index).
		Put("anonymous", field.Anonymous).
		FilterValues(rfl.IsNotEmpty[any]).
		Build()
}

func fieldInfoRepr(fi *FieldInfo) ctr.Map[string, any] {
	return ctr.NewMapBuilder[string, any]().
		Put("name", fi.name.s).
		Put("field", structFieldRepr(fi.field)).
		Put("index", fi.index).
		Put("path", fi.path).
		Put("embedded", fi.embedded).
		Put("children", slices.Map(fieldInfoRepr, fi.children)).
		FilterValues(rfl.IsNotEmpty[any]).
		Build()
}

func structInfoRepr(si *StructInfo) ctr.Map[string, any] {
	return ctr.NewMapBuilder[string, any]().
		Put("name", si.name.String()).
		Put("type", si.ty.Name()).
		Put("fields", slices.Map(fieldInfoRepr, si.fields.root)).
		FilterValues(rfl.IsNotEmpty[any]).
		Build()
}
