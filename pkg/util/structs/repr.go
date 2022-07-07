package structs

import (
	"reflect"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	"github.com/wrmsr/bane/pkg/util/slices"
)

func typeRepr(ty reflect.Type) ctr.Map[string, any] {
	return ctr.NewOrderedMapBuilder[string, any]().
		Put("name", rfl.TypeName(ty)).
		Put("kind", ty.Kind().String()).
		FilterValues(rfl.IsNotEmpty[any]).
		Build()
}

func structFieldRepr(field reflect.StructField) ctr.Map[string, any] {
	return ctr.NewOrderedMapBuilder[string, any]().
		Put("pkg_path", field.PkgPath).
		Put("type", typeRepr(field.Type)).
		Put("tag", string(field.Tag)).
		Put("offset", field.Offset).
		Put("index", field.Index).
		Put("anonymous", field.Anonymous).
		FilterValues(rfl.IsNotEmpty[any]).
		Build()
}

func fieldInfoRepr(fi *FieldInfo) ctr.Map[string, any] {
	return ctr.NewOrderedMapBuilder[string, any]().
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
	return ctr.NewOrderedMapBuilder[string, any]().
		Put("name", si.name.String()).
		Put("type", typeRepr(si.ty)).
		Put("fields", its.Seq(its.Map[*FieldInfo](si.fields.root, fieldInfoRepr))).
		FilterValues(rfl.IsNotEmpty[any]).
		Build()
}
