package inject

import (
	"fmt"
	"reflect"
	"strings"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

type Key struct {
	ty  reflect.Type
	arr bool
	tag any
}

func KeyOf[T any]() Key {
	var z T
	return Key{ty: reflect.TypeOf(z)}
}

func AsKey(o any) Key {
	if k, ok := o.(Key); ok {
		return k
	}

	return Key{
		ty: rfl.AsType(o),
	}
}

func Tag(o, tag any) Key {
	k := AsKey(o)
	k.tag = tag
	return k
}

func (k Key) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Key{ty: %s", k.ty))
	if k.arr {
		sb.WriteString(", arr: true")
	}
	if k.tag != nil {
		sb.WriteString(fmt.Sprintf(", tag: %v", k.tag))
	}
	sb.WriteString("}")
	return sb.String()
}
