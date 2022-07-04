package structs

import (
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	"github.com/wrmsr/bane/pkg/util/slices"
)

func getSimpleFieldInfo(
	ty reflect.Type,
	indexPrefix []int,
	pathPrefix string,
) []*FieldInfo {
	fis := make([]*FieldInfo, ty.NumField())
	for i := 0; i < ty.NumField(); i++ {
		field := ty.Field(i)

		nameBytes := []byte(field.Name)
		fi := &FieldInfo{
			field: field,

			name:      field.Name,
			nameBytes: nameBytes,

			equalFold: foldFunc(nameBytes),

			index: slices.MakeAppend(indexPrefix, i),
			path:  pathPrefix + field.Name,
		}
		fis[i] = fi

		fty := rfl.UnwrapPointerType(field.Type)
		if field.Anonymous && fty.Kind() == reflect.Struct {
			fi.children = getSimpleFieldInfo(
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
