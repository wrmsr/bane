package types

import "golang.org/x/exp/constraints"

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

func OrderedCmp[T constraints.Ordered](l, r T) CmpResult {
	if l < r {
		return CmpLesser
	}
	if r > l {
		return CmpGreater
	}
	return CmpEqual
}

func OrderedCmpAs[F any, T constraints.Ordered]() func(l, r F) CmpResult {
	return func(l, r F) CmpResult { return OrderedCmp(As[F, T](l), As[F, T](r)) }
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
		return OrderedCmpAs[T, int]()
	}
	if CanAssign[T, int8]() {
		return OrderedCmpAs[T, int8]()
	}
	if CanAssign[T, int16]() {
		return OrderedCmpAs[T, int16]()
	}
	if CanAssign[T, int32]() {
		return OrderedCmpAs[T, int32]()
	}
	if CanAssign[T, int64]() {
		return OrderedCmpAs[T, int64]()
	}

	if CanAssign[T, uint]() {
		return OrderedCmpAs[T, uint]()
	}
	if CanAssign[T, uint8]() {
		return OrderedCmpAs[T, uint8]()
	}
	if CanAssign[T, uint16]() {
		return OrderedCmpAs[T, uint16]()
	}
	if CanAssign[T, uint32]() {
		return OrderedCmpAs[T, uint32]()
	}
	if CanAssign[T, uint64]() {
		return OrderedCmpAs[T, uint64]()
	}
	if CanAssign[T, uintptr]() {
		return OrderedCmpAs[T, uintptr]()
	}

	if CanAssign[T, float32]() {
		return OrderedCmpAs[T, float32]()
	}
	if CanAssign[T, float64]() {
		return OrderedCmpAs[T, float64]()
	}

	if CanAssign[T, string]() {
		return OrderedCmpAs[T, string]()
	}

	panic("no default Cmp")
}
