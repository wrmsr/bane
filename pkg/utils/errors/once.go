package errors

import (
	bat "github.com/wrmsr/bane/pkg/utils/atomic"
)

func ErrOnce(fn func() error) func() error {
	lz := bat.Lazy[error]{Fn: fn}
	return lz.Get
}
