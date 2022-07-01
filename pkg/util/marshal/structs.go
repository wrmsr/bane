package marshal

import (
	"reflect"

	stu "github.com/wrmsr/bane/pkg/util/structs"
)

//

type StructMarshaler struct {
	si *stu.StructInfo
}

func NewStructMarshaler(si *stu.StructInfo) StructMarshaler {
	return StructMarshaler{si: si}
}

var _ Marshaler = StructMarshaler{}

func (s StructMarshaler) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	panic("implement me")
}
