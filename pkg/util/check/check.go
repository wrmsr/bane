package check

import (
	"fmt"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

func NotZero[T comparable](v T, s ...string) T {
	if v == bt.Zero[T]() {
		if len(s) > 0 {
			panic(s[0])
		}
		panic("cannot be nil")
	}
	return v
}

func NotNil(v any, s ...string) any {
	if v == nil {
		if len(s) > 0 {
			panic(s[0])
		}
		panic("cannot be nil")
	}
	return v
}

func NoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Condition(b bool) {
	if !b {
		panic("condition not met")
	}
}

func Equal[T comparable](actual, expected T) T {
	if actual != expected {
		panic(fmt.Sprintf("must be equal: %v != %v", actual, expected))
	}
	return actual
}

func NotEqual[T comparable](actual, expected T) T {
	if actual == expected {
		panic(fmt.Sprintf("must not be equal: %v == %v", actual, expected))
	}
	return actual
}

//

type HasLen interface {
	Len() int
}

func NotEmpty(ctr HasLen) {
	if ctr.Len() < 1 {
		panic("must not be empty")
	}
}

//

type Container[T any] interface {
	Contains(v T) bool
}

func Contains[T any](ctr Container[T], v T) T {
	if !ctr.Contains(v) {
		panic(fmt.Sprintf("container must contain %v", v))
	}
	return v
}

func NotContains[T any](ctr Container[T], v T) T {
	if ctr.Contains(v) {
		panic(fmt.Sprintf("container must not contain %v", v))
	}
	return v
}
