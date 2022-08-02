package container

import (
	"sync"
	"sync/atomic"

	its "github.com/wrmsr/bane/pkg/util/iterators"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

type CowList[T any] struct {
	mtx sync.Mutex
	r   atomic.Value
}

func NewCowList[T any](it bt.Iterable[T]) *CowList[T] {
	l := &CowList[T]{}
	if it != nil {
		l.r.Store(its.Seq(it))
	}
	return l
}

func NewCowListOf[T any](vs ...T) *CowList[T] {
	return NewCowList(its.Of(vs...))
}

func (l *CowList[T]) get() []T {
	if r := l.r.Load(); r != nil {
		return r.([]T)
	}
	return nil
}

var _ SyncMutList[int] = &CowList[int]{}

func (l *CowList[T]) isMut()  {}
func (l *CowList[T]) isSync() {}

func (l *CowList[T]) Len() int {
	return len(l.get())
}

func (l *CowList[T]) Get(i int) T {
	return l.get()[i]
}

func (l *CowList[T]) Iterate() bt.Iterator[T] {
	return its.OfSlice(l.get()).Iterate()
}

func (l *CowList[T]) ForEach(fn func(v T) bool) bool {
	for _, v := range l.get() {
		if !fn(v) {
			return false
		}
	}
	return true
}

func (l *CowList[T]) mut(fn func([]T) []T) {
	l.mtx.Lock()
	defer l.mtx.Unlock()

	l.r.Store(fn(slices.Clone(l.get())))
}

func (l *CowList[T]) Put(i int, v T) {
	l.mut(func(s []T) []T {
		s[i] = v
		return s
	})
}

func (l *CowList[T]) Append(v T) {
	l.mut(func(s []T) []T {
		return append(s, v)
	})
}

func (l *CowList[T]) Delete(i int) {
	l.mut(func(s []T) []T {
		return append(s[:i], s[i+1:]...)
	})
}
