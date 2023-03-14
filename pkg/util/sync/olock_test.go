package sync

import (
	"sync/atomic"
	"testing"
)

type Olock struct {
	o  atomic.Value
	d  int
	_c Lazy[chan struct{}]
}

func (l *Olock) c() chan struct{} {
	return l._c.Get(func() chan struct{} {
		c := make(chan struct{}, 1)
		c <- struct{}{}
		return c
	})
}

func (l *Olock) Acquire(o any) int {

}

func (l *Olock) Release(o any) int {

}

func TestOlock(t *testing.T) {

}
