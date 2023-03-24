package types

import (
	"fmt"
	"reflect"

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

func HashEqAs[F, T any](impl HashEqImpl[T]) HashEqImpl[F] {
	return HashEqImpl[F]{
		Hash: func(o F) uintptr { return impl.Hash(As[F, T](o)) },
		Eq:   func(l, r F) bool { return impl.Eq(As[F, T](l), As[F, T](r)) },
	}
}

func BoolHashEq() HashEqImpl[bool] {
	return HashEqImpl[bool]{
		Hash: func(i bool) uintptr { return BoolNum[uintptr](i) },
		Eq:   func(l, r bool) bool { return l == r },
	}
}

func IntHashEq[T constraints.Integer]() HashEqImpl[T] {
	return HashEqImpl[T]{
		Hash: func(i T) uintptr { return uintptr(i) }, // FIXME: barf
		Eq:   func(l, r T) bool { return l == r },
	}
}

func StrHashEq() HashEqImpl[string] {
	return HashEqImpl[string]{
		Hash: HashStr,
		Eq:   Eq[string],
	}
}

func DefaultHashEqImpl[T any]() HashEqImpl[T] {
	if CanAssign[T, HashEq[T]]() {
		return HashEqImpl[T]{
			Hash: func(i T) uintptr { return As[T, Hasher](i).Hash() },
			Eq:   func(l, r T) bool { return As[T, Equaler[T]](l).Equals(r) },
		}
	}

	if CanAssign[T, bool]() {
		return HashEqAs[T, bool](BoolHashEq())
	}

	if CanAssign[T, int]() {
		return HashEqAs[T, int](IntHashEq[int]())
	}
	if CanAssign[T, int8]() {
		return HashEqAs[T, int8](IntHashEq[int8]())
	}
	if CanAssign[T, int16]() {
		return HashEqAs[T, int16](IntHashEq[int16]())
	}
	if CanAssign[T, int32]() {
		return HashEqAs[T, int32](IntHashEq[int32]())
	}
	if CanAssign[T, int64]() {
		return HashEqAs[T, int64](IntHashEq[int64]())
	}

	if CanAssign[T, uint]() {
		return HashEqAs[T, uint](IntHashEq[uint]())
	}
	if CanAssign[T, uint8]() {
		return HashEqAs[T, uint8](IntHashEq[uint8]())
	}
	if CanAssign[T, uint16]() {
		return HashEqAs[T, uint16](IntHashEq[uint16]())
	}
	if CanAssign[T, uint32]() {
		return HashEqAs[T, uint32](IntHashEq[uint32]())
	}
	if CanAssign[T, uint64]() {
		return HashEqAs[T, uint64](IntHashEq[uint64]())
	}
	if CanAssign[T, uintptr]() {
		return HashEqAs[T, uintptr](IntHashEq[uintptr]())
	}

	if CanAssign[T, string]() {
		return HashEqAs[T, string](StrHashEq())
	}

	var z T
	panic(fmt.Errorf("no default HashEq for %s", reflect.TypeOf(z)))
}
