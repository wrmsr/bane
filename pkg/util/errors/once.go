package errors

import (
	syncu "github.com/wrmsr/bane/pkg/util/sync"
)

func ErrOnce(fn func() error) func() error {
	lz := syncu.Lazy[error]{}
	return func() error { return lz.Get(fn) }
}
