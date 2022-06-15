package check

import bt "github.com/wrmsr/bane/pkg/utils/types"

func NotZero[T comparable](v T, s ...string) T {
	if v == bt.Zero[T]() {
		if len(s) > 0 {
			panic(s[0])
		}
		panic("cannot be nil")
	}
	return v
}
