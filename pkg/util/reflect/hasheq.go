package reflect

import (
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func PtrHashEq[T any]() bt.HashEqImpl[*T] {
	return bt.HashEqImpl[*T]{
		Hash: func(p *T) uintptr {
			return PointerOf(p)
		},
		Eq: func(l, r *T) bool {
			return l == r
		},
	}
}

func AddrHashEq[T any]() bt.HashEqImpl[T] {
	return bt.HashEqImpl[T]{
		Hash: func(v T) uintptr {
			return PointerOf(v)
		},
		Eq: func(l, r T) bool {
			return PointerOf(l) == PointerOf(r)
		},
	}
}
