/*
TODO:
 - reflective - tailscale deephash
 - adapter - HashEqImpl{foo}
*/
package types

//

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

//

type HashImpl[T any] func(T) uintptr
type EqImpl[T any] func(l, r T) bool

type HashEqImpl[T any] struct {
	Hash HashImpl[T]
	Eq   EqImpl[T]
}

func HashEqOf[T any](hash HashImpl[T], eq EqImpl[T]) HashEqImpl[T] {
	return HashEqImpl[T]{Hash: hash, Eq: eq}
}
