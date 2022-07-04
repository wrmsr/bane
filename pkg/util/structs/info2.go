package structs

import "reflect"

func getSimpleFieldInfo(ty reflect.Type) []*FieldInfo {
	fis := make([]*FieldInfo, ty.NumField())
	for i := 0; i < ty.NumField(); i++ {
		field := ty.Field(i)
		nameBytes := []byte(field.Name)
		fis[i] = &FieldInfo{
			field: field,

			name:      field.Name,
			nameBytes: nameBytes,

			equalFold: foldFunc(nameBytes),
		}
	}
	return fis
}
