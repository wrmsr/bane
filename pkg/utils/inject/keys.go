package inject

import (
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/utils/reflect"
)

type Key struct {
	ty  reflect.Type
	arr bool
	tag any
}

func ToKey(o any) Key {
	if k, ok := o.(Key); ok {
		return k
	}

	return Key{
		ty: rfl.AsType(o),
	}
}
