package types

import (
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
)

//

type CmpResult int8

const (
	CmpLesser  = -1
	CmpEqual   = 0
	CmpGreater = 1
)

//

type Comparer[T any] interface {
	Compare(o T) CmpResult
}

//

type CmpImpl[T any] func(l, r T) CmpResult

func MethodCmpImpl[T Comparer[T]]() CmpImpl[T] {
	return T.Compare
}

func CmpAs[F, T any](impl CmpImpl[T]) func(l, r F) CmpResult {
	return func(l, r F) CmpResult { return impl(As[F, T](l), As[F, T](r)) }
}

func OrderedCmp[T constraints.Ordered](l, r T) CmpResult {
	if l < r {
		return CmpLesser
	}
	if r > l {
		return CmpGreater
	}
	return CmpEqual
}

func BoolCmp(l, r bool) CmpResult {
	return OrderedCmp(BoolNum[int](l), BoolNum[int](r))
}

func DefaultCmpImpl[T any]() CmpImpl[T] {
	if CanAssign[T, Comparer[T]]() {
		return func(l, r T) CmpResult { return As[T, Comparer[T]](l).Compare(r) }
	}

	if CanAssign[T, bool]() {
		return func(l, r T) CmpResult { return BoolCmp(As[T, bool](l), As[T, bool](r)) }
	}

	if CanAssign[T, int]() {
		return CmpAs[T, int](OrderedCmp[int])
	}
	if CanAssign[T, int8]() {
		return CmpAs[T, int8](OrderedCmp[int8])
	}
	if CanAssign[T, int16]() {
		return CmpAs[T, int16](OrderedCmp[int16])
	}
	if CanAssign[T, int32]() {
		return CmpAs[T, int32](OrderedCmp[int32])
	}
	if CanAssign[T, int64]() {
		return CmpAs[T, int64](OrderedCmp[int64])
	}

	if CanAssign[T, uint]() {
		return CmpAs[T, uint](OrderedCmp[uint])
	}
	if CanAssign[T, uint8]() {
		return CmpAs[T, uint8](OrderedCmp[uint8])
	}
	if CanAssign[T, uint16]() {
		return CmpAs[T, uint16](OrderedCmp[uint16])
	}
	if CanAssign[T, uint32]() {
		return CmpAs[T, uint32](OrderedCmp[uint32])
	}
	if CanAssign[T, uint64]() {
		return CmpAs[T, uint64](OrderedCmp[uint64])
	}
	if CanAssign[T, uintptr]() {
		return CmpAs[T, uintptr](OrderedCmp[uintptr])
	}

	if CanAssign[T, float32]() {
		return CmpAs[T, float32](OrderedCmp[float32])
	}
	if CanAssign[T, float64]() {
		return CmpAs[T, float64](OrderedCmp[float64])
	}

	if CanAssign[T, string]() {
		return CmpAs[T, string](OrderedCmp[string])
	}

	var z T
	panic(fmt.Errorf("no default Cmp for %s", reflect.TypeOf(z)))
}
