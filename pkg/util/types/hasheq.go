/*
TODO:
 - reflective - tailscale deephash
 - adapter - HashEqImpl{foo}
*/
package types

type Hasher interface {
	Hash() uintptr
}

type Equaler[T any] interface {
	Equals(o T) bool
}

type HashEq[T any] interface {
	Hasher
	Equaler[T]
}
