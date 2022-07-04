package structs

import (
	"reflect"

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
		fis[i] = &FieldInfo{
			field: field,

			name:      field.Name,
			nameBytes: nameBytes,

			equalFold: foldFunc(nameBytes),

			index: slices.MakeAppend(indexPrefix, i),
			path:  pathPrefix + field.Name,
		}
	}
	return fis
}
