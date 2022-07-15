package types

import "golang.org/x/exp/constraints"

func Abs[T Rational](v T) T {
	if v < 0 {
		return -v
	}
	return v
}

func Min[T constraints.Ordered](vs ...T) T {
	var r T
	for i, v := range vs {
		if i == 0 || v < r {
			v = r
		}
	}
	return r
}

func Max[T constraints.Ordered](vs ...T) T {
	var r T
	for i, v := range vs {
		if i == 0 || v > r {
			v = r
		}
	}
	return r
}
