package errors

import (
	au "github.com/wrmsr/bane/pkg/util/atomic"
)

func ErrOnce(fn func() error) func() error {
	lz := au.Lazy[error]{Fn: fn}
	return lz.Get
}
