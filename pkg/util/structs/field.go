package structs

import (
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	"github.com/wrmsr/bane/pkg/util/slices"
)

type FieldInfo struct {
	field reflect.StructField

	name Name
	flat bool

	index []int
	path  string

	embedded bool

	parent   *FieldInfo
	children []*FieldInfo

	si *StructInfo
}

func (fi FieldInfo) Field() reflect.StructField { return fi.field }
func (fi FieldInfo) Name() Name                 { return fi.name }
func (fi FieldInfo) Flat() bool                 { return fi.flat }
func (fi FieldInfo) Index() []int               { return fi.index }
func (fi FieldInfo) Struct() *StructInfo        { return fi.si }

func (fi FieldInfo) Type() reflect.Type { return fi.field.Type }
func (fi FieldInfo) Anonymous() bool    { return fi.field.Anonymous }
func (fi FieldInfo) Dupe() bool         { return !fi.flat }
func (fi FieldInfo) Depth() int         { return len(fi.index) }

func buildFieldInfo(
	ty reflect.Type,
	indexPrefix []int,
	pathPrefix string,
) []*FieldInfo {
	fis := make([]*FieldInfo, ty.NumField())
	for i := 0; i < ty.NumField(); i++ {
		field := ty.Field(i)

		fi := &FieldInfo{
			field: field,

			name: buildName(field.Name),

			index: slices.MakeAppend(indexPrefix, i),
			path:  pathPrefix + field.Name,
		}
		fis[i] = fi

		fty := rfl.UnwrapPointerType(field.Type)
		if field.Anonymous && fty.Kind() == reflect.Struct {
			fi.children = buildFieldInfo(
				fty,
				fi.index,
				fi.path+".",
			)

			for _, c := range fi.children {
				c.embedded = true
				c.parent = fi
			}
		}
	}
	return fis
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

func (fi FieldInfo) SetValue(s, v any) bool {
	subv := reflect.ValueOf(s)
	for _, i := range fi.index {
		if subv.Kind() == reflect.Pointer {
			if subv.IsNil() {
				if !subv.CanSet() {
					return false
				}
				subv.Set(reflect.New(subv.Type().Elem()))
			}
			subv = subv.Elem()
		}
		subv = subv.Field(i)
	}
	subv.Set(reflect.ValueOf(v))
	return true
}
