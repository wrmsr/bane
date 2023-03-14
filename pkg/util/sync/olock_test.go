package sync

import (
	"fmt"
	"sync/atomic"
	"testing"
)

type Olock struct {
	o  atomic.Value
	d  int
	_c Lazy[chan struct{}]
}

type __unowned struct{}

var _unowned = __unowned{}

func (l *Olock) c() chan struct{} {
	return l._c.Get(func() chan struct{} {
		c := make(chan struct{}, 1)
		c <- struct{}{}
		return c
	})
}

func (l *Olock) Acquire(o any) int {
	if co := l.o.Load(); co != o {
		c := l.c()
		_ = <-c
		x := l.o.Load()
		if x != nil && x != _unowned {
			panic("oops")
		}
		l.o.Store(o)
	}
	l.d++
	return l.d
}

func (l *Olock) Release(o any) int {
	if co := l.o.Load(); co != o {
		panic("oops")
	}
	l.d--
	if l.d < 0 {
		panic("oops")
	}
	if l.d < 1 {
		l.o.Store(_unowned)
		l.c() <- struct{}{}
	}
	return l.d
}

func TestOlock(t *testing.T) {
	l := Olock{}
	for i := 0; i < 3; i++ {
		fmt.Println(l.Acquire(420))
	}
	for i := 0; i < 3; i++ {
		fmt.Println(l.Release(420))
	}
}
