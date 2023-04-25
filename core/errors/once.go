package errors

import (
	syncu "github.com/wrmsr/bane/core/sync"
)

func ErrOnce(fn func() error) func() error {
	lz := syncu.Lazy[error]{}
	return func() error { return lz.Get(fn) }
}
