package reflect

import (
	"reflect"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

func PtrHashEq[T any]() bt.HashEqImpl[*T] {
	return bt.HashEqImpl[*T]{
		Hash: func(p *T) uintptr {
			if p == nil {
				return 0
			}
			return reflect.ValueOf(p).Pointer()
		},
		Eq: func(l, r *T) bool {
			return l == r
		},
	}
}

func AddrHashEq[T any]() bt.HashEqImpl[T] {
	return bt.HashEqImpl[T]{
		Hash: func(p T) uintptr {
			rv := reflect.ValueOf(p)
			if rv.IsNil() {
				return 0
			}
			return rv.Pointer()
		},
		Eq: func(l, r T) bool {
			return reflect.ValueOf(l).Pointer() == reflect.ValueOf(r).Pointer()
		},
	}
}
