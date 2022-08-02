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
			r = v
		}
	}
	return r
}

func Max[T constraints.Ordered](vs ...T) T {
	var r T
	for i, v := range vs {
		if i == 0 || v > r {
			r = v
		}
	}
	return r
}

func Clamp[T constraints.Ordered](v, lo, hi T) T {
	if v < lo {
		v = lo
	}
	if v > hi {
		v = hi
	}
	return v
}

func Sum[T Rational](vs ...T) T {
	var r T
	for _, v := range vs {
		r += v
	}
	return r
}

func Prod[T Rational](vs ...T) T {
	var r T = 1
	for _, v := range vs {
		r *= v
	}
	return r
}
