package structs

import (
	"reflect"
	"sync"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

type StructInfoCache struct {
	mtx         sync.Mutex
	structInfos map[reflect.Type]*StructInfo
}

func NewStructInfoCache() *StructInfoCache {
	return &StructInfoCache{
		structInfos: make(map[reflect.Type]*StructInfo),
	}
}

func (sic *StructInfoCache) Info(ty any) *StructInfo {
	sic.mtx.Lock()
	defer sic.mtx.Unlock()

	rty := rfl.AsType(ty)
	if si, ok := sic.structInfos[rty]; ok {
		return si
	}

	si := buildStructInfo(rty)
	sic.structInfos[rty] = si
	return si
}
