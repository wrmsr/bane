package structs

import (
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

type StructInfoCache struct {
	structInfos map[reflect.Type]*StructInfo
}

func NewStructInfoCache() *StructInfoCache {
	return &StructInfoCache{
		structInfos: make(map[reflect.Type]*StructInfo),
	}
}

func (sic *StructInfoCache) Info(ty any) *StructInfo {
	rty := rfl.AsType(ty)
	if si, ok := sic.structInfos[rty]; ok {
		return si
	}

	si := buildStructInfo(rty)
	sic.structInfos[rty] = si
	return si
}
