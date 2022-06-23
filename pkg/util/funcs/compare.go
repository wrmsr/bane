package funcs

import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](v0, v1 T) T {
	if v0 > v1 {
		return v1
	}
	return v0
}

func Max[T constraints.Ordered](v0, v1 T) T {
	if v0 < v1 {
		return v1
	}
	return v0
}
