package inject

import (
	"fmt"
	"reflect"
	"strings"

	rfl "github.com/wrmsr/bane/pkg/utils/reflect"
)

type Key struct {
	ty  reflect.Type
	arr bool
	tag any
}

func (k Key) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Key{ty: %s", k.ty.Name()))
	if k.arr {
		sb.WriteString(", arr: true")
	}
	if k.tag != nil {
		sb.WriteString(fmt.Sprintf(", tag: %v", k.tag))
	}
	sb.WriteString("}")
	return sb.String()
}

func ToKey(o any) Key {
	if k, ok := o.(Key); ok {
		return k
	}

	return Key{
		ty: rfl.AsType(o),
	}
}
