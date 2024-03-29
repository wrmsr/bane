package check

import (
	"golang.org/x/exp/constraints"

	bt "github.com/wrmsr/bane/core/types"
)

//

func NotZero[T comparable](v T, details ...any) T {
	if v == bt.Zero[T]() {
		panic(checkError("cannot be zero", []any{v}, details...))
	}
	return v
}

func NotNil(v any, details ...any) any {
	if v == nil {
		panic(checkError("cannot be nil", []any{v}, details...))
	}
	return v
}

func Nil(v any, details ...any) any {
	if v != nil {
		panic(checkError("must be nil", []any{v}, details...))
	}
	return v
}

func Condition(b bool, details ...any) {
	if !b {
		panic(checkError("condition not met", nil, details...))
	}
}

func Equal[T comparable](actual, expected T, details ...any) T {
	if actual != expected {
		panic(checkError("must be equal", []any{actual, expected}, details...))
	}
	return actual
}

func NotEqual[T comparable](actual, expected T, details ...any) T {
	if actual == expected {
		panic(checkError("must not be equal", []any{actual, expected}, details...))
	}
	return actual
}

func Single[T any](s []T, details ...any) T {
	if len(s) != 1 {
		panic(checkError("must be single", []any{s}, details...))
	}
	return s[0]
}

func Pair[T any](s []T, details ...any) (T, T) {
	if len(s) != 2 {
		panic(checkError("must be pair", []any{s}, details...))
	}
	return s[0], s[1]
}

func NotEmptySlice[T any](s []T, details ...any) []T {
	if len(s) < 1 {
		panic(checkError("must not be empty slice", []any{s}, details...))
	}
	return s
}

func EmptySlice[T any](s []T, details ...any) []T {
	if len(s) > 0 {
		panic(checkError("must be empty slice", []any{s}, details...))
	}
	return s
}

func NotEmptyStr(s string, details ...any) string {
	if s == "" {
		panic(checkError("must not be empty string", []any{s}, details...))
	}
	return s
}

func Between[T constraints.Ordered](v, lo, hi T, details ...any) T {
	if v < lo || v > hi {
		panic(checkError("must be between", []any{v, lo, hi}, details...))
	}
	return v
}

//

type HasLen interface {
	Len() int
}

func NotEmpty[T HasLen](ctr T, details ...any) T {
	if ctr.Len() < 1 {
		panic(checkError("must not be empty", []any{ctr}, details...))
	}
	return ctr
}

//

type Container[T any] interface {
	Contains(v T) bool
}

func Contains[T any](ctr Container[T], v T, details ...any) T {
	if !ctr.Contains(v) {
		panic(checkError("container must contain", []any{ctr, v}, details...))
	}
	return v
}

func NotContains[T any](ctr Container[T], v T, details ...any) T {
	if ctr.Contains(v) {
		panic(checkError("container must not contain", []any{ctr, v}, details...))
	}
	return v
}
