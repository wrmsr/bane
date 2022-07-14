package types

import "golang.org/x/exp/constraints"

//

type CmpResult int8

const (
	CmpLess    = -1
	CmpEqual   = 0
	CmpGreater = 1
)

//

type Comparer[T any] interface {
	Compare(o T) CmpResult
}

//

type CmpImpl[T any] func(l, r T) CmpResult

func OrderedCmp[T constraints.Ordered](l, r T) CmpResult {
	if l < r {
		return CmpLess
	}
	if r > l {
		return CmpGreater
	}
	return CmpEqual
}

func BoolCmp(l, r bool) CmpResult {
	return OrderedCmp(BoolNum[int](l), BoolNum[int](r))
}
