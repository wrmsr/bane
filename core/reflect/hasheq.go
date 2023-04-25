package reflect

import (
	bt "github.com/wrmsr/bane/core/types"
)

func PtrHashEq[T any]() bt.HashEqImpl[*T] {
	return bt.HashEqImpl[*T]{
		Hash: func(p *T) uintptr {
			return AddressOf(p)
		},
		Eq: func(l, r *T) bool {
			return l == r
		},
	}
}

func AddrHashEq[T any]() bt.HashEqImpl[T] {
	return bt.HashEqImpl[T]{
		Hash: func(v T) uintptr {
			return AddressOf(v)
		},
		Eq: func(l, r T) bool {
			return AddressOf(l) == AddressOf(r)
		},
	}
}
