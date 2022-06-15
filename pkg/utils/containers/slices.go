package containers

import bt "github.com/wrmsr/bane/pkg/utils/types"

func DeleteAt[T any](s []T, i int) []T {
	copy(s[i:], s[i+1:])
	s[len(s)-1] = bt.Zero[T]()
	return s[:len(s)-1]
}
