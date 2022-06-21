package structs

import (
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

type StructTool struct {
	structInfos map[reflect.Type]*StructInfo
}

func NewStructTool() *StructTool {
	return &StructTool{
		structInfos: make(map[reflect.Type]*StructInfo),
	}
}

func (st *StructTool) Info(ty any) *StructInfo {
	rty := rfl.AsType(ty)
	if si, ok := st.structInfos[rty]; ok {
		return si
	}

	si := newStructInfo(rty)
	st.structInfos[rty] = si
	return si
}
