package sync

import "sync"

type OMutex struct {
	mtx sync.Mutex

	o any
	d int
	w int
	c chan struct{}
}

func (l *OMutex) Lock(o any) int {
	l.mtx.Lock()
	if l.c == nil {
		l.c = make(chan struct{}, 1)
	}
	if l.o == nil {
		if l.d != 0 {
			l.mtx.Unlock()
			panic("oops")
		}
		l.o = o
	} else if l.o != o && l.o != nil {
		if l.d < 1 {
			l.mtx.Unlock()
			panic("oops")
		}
		l.w++
		c := l.c
		l.mtx.Unlock()
		_ = <-c
		l.mtx.Lock()
		l.w--
		if l.o != nil || l.d != 0 {
			l.mtx.Unlock()
			panic("oops")
		}
		l.o = o
	}
	l.d++
	d := l.d
	l.mtx.Unlock()
	return d
}

func (l *OMutex) TryLock(o any) int {
	l.mtx.Lock()
	var d int
	if l.o == o {
		d = l.d
	}
	l.mtx.Unlock()
	return d
}

func (l *OMutex) Unlock(o any) int {
	l.mtx.Lock()
	if l.o != o || l.d < 1 || l.c == nil {
		l.mtx.Unlock()
		panic("oops")
	}
	l.d--
	d := l.d
	if d < 1 {
		l.o = nil
		if l.w > 0 {
			l.c <- struct{}{}
		}
	}
	l.mtx.Unlock()
	return d
}
