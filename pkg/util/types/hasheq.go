package types

import (
	"encoding/binary"
	"hash/fnv"

	"golang.org/x/exp/constraints"
)

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

func MethodHashEqImpl[T HashEq[T]]() HashEqImpl[T] {
	return HashEqImpl[T]{
		Hash: T.Hash,
		Eq:   T.Equals,
	}
}

func IntHashEq[T constraints.Integer]() HashEqImpl[T] {
	return HashEqImpl[T]{
		Hash: func(i T) uintptr { return uintptr(i) },
		Eq:   func(l, r T) bool { return l == r },
	}
}

func HashStr(s string) uintptr {
	h := fnv.New64a()
	var ub [4]byte
	u := ub[:]
	for _, b := range s {
		binary.LittleEndian.PutUint32(u, uint32(b))
		_, _ = h.Write(u)
	}
	return uintptr(h.Sum64())
}

func StrHashEq() HashEqImpl[string] {
	return HashEqImpl[string]{
		Hash: HashStr,
		Eq:   Eq[string],
	}
}

func DefaultHashEqImpl[T any]() HashEqImpl[T] {
	if !CanAssign[T, HashEq[T]]() {
		panic("no default HashEq")
	}
	return HashEqImpl[T]{
		Hash: func(i T) uintptr { return As[T, Hasher](i).Hash() },
		Eq:   func(l, r T) bool { return As[T, Equaler[T]](l).Equals(r) },
	}
}
