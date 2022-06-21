package slices

import (
	"sort"

	"golang.org/x/exp/constraints"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

func DeleteAt[T any](s []T, i int) []T {
	copy(s[i:], s[i+1:])
	s[len(s)-1] = bt.Zero[T]()
	return s[:len(s)-1]
}

func Join[T any](ss ...[]T) []T {
	var r []T
	for _, s := range ss {
		r = append(r, s...)
	}
	return r
}

func Find[T comparable](s []T, v T) (int, bool) {
	for i, c := range s {
		if c == v {
			return i, true
		}
	}
	return -1, false
}

func Contains[T comparable](s []T, v T) bool {
	_, ok := Find(s, v)
	return ok
}

func Reversed[T any](s []T) []T {
	r := make([]T, len(s))
	n := len(s) - 1
	for _, e := range s {
		r[n] = e
		n--
	}
	return r
}

func Sort[T constraints.Ordered](s []T) []T {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})
	return s
}

func Equal[T comparable](l, r []T) bool {
	if len(l) != len(r) {
		return false
	}
	for i := range l {
		if l[i] != r[i] {
			return false
		}
	}
	return true
}
