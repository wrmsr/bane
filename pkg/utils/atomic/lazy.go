package atomic

import (
	"sync"
	"sync/atomic"
)

type Lazy[T any] struct {
	fn func()
	o  sync.Once
	v  atomic.Value
}

func NewLazy()
