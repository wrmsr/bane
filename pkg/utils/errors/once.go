package errors

import (
	au "github.com/wrmsr/bane/pkg/utils/atomic"
)

func ErrOnce(fn func() error) func() error {
	lz := au.Lazy[error]{Fn: fn}
	return lz.Get
}
