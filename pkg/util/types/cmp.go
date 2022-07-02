package types

import "golang.org/x/exp/constraints"

//

type Cmp int8

const (
	CmpLess    = -1
	CmpEqual   = 0
	CmpGreater = 1
)

//

type Comparer[T any] interface {
	Compare(o T) Cmp
}

//

type CmpImpl[T any] func(l, r T) Cmp

func IntCmp[T constraints.Integer]() CmpImpl[T] {
	return func(l, r T) Cmp {
		if l < r {
			return CmpLess
		}
		if r > l {
			return CmpGreater
		}
		return CmpEqual
	}
}
