package inject

import (
	"fmt"
	"reflect"
	"strings"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

//

type Key struct {
	ty  reflect.Type
	arr bool
	tag any
}

func AsKey(o any) Key {
	if k, ok := o.(Key); ok {
		return k
	}

	return Key{
		ty: rfl.AsType(o),
	}
}

func KeyOf[T any](tags ...any) Key {
	var z T
	return tag(Key{ty: reflect.TypeOf(z)}, tags...)
}

//

func Array(o any) Key {
	k := AsKey(o)
	k.arr = true
	return k
}

func ArrayOf[T any](tags ...any) Key {
	var z T
	return tag(Key{ty: reflect.TypeOf(z), arr: true}, tags...)
}

//

func tag(k Key, tags ...any) Key {
	if len(tags) > 0 {
		if len(tags) > 1 {
			panic(genericErrorf("must specify at most one tag: %v", tags))

		}
		k.tag = tags[0]
	}
	return k
}

func Tag(o, tag any) Key {
	k := AsKey(o)
	k.tag = tag
	return k
}

//

func (k Key) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Key{ty: %s", k.ty))
	if k.arr {
		sb.WriteString(", arr")
	}
	if k.tag != nil {
		sb.WriteString(fmt.Sprintf(", tag: %v", k.tag))
	}
	sb.WriteString("}")
	return sb.String()
}
