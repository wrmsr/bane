package tg

import (
	"github.com/wrmsr/bane/pkg/util/maps"
)

//

type Lazy interface {
	isLazy()
}

//

type LazyOp struct {
	op  Op
	src []Lazy
	arg any
}

var _ Lazy = &LazyOp{}

func (o *LazyOp) isLazy() {}

func (o *LazyOp) ForEachBuffer(fn func(*LazyBuffer) bool) bool {
	for _, s := range o.src {
		switch s := s.(type) {
		case *LazyOp:
			if !s.ForEachBuffer(fn) {
				return false
			}
		case *LazyBuffer:
			if !fn(s) {
				return false
			}
		}
	}
	return true
}

func (o *LazyOp) ForEachOp(fn func(*LazyOp) bool) bool {
	for _, s := range o.src {
		switch s := s.(type) {
		case *LazyOp:
			if !fn(s) {
				return false
			}
		}
	}
	return true
}

//

type LazyBuffer struct {
	st *ShapeTracker
	ot OpType
	op *LazyOp

	realized *Buffer
	children maps.Set[*LazyBuffer]
}

func NewLazyBuffer(st *ShapeTracker, ot OpType, op *LazyOp) *LazyBuffer {
	b := &LazyBuffer{
		st: st,
		ot: ot,
		op: op,

		children: maps.MakeSet[*LazyBuffer](),
	}

	op.ForEachBuffer(func(s *LazyBuffer) bool {
		s.children.Add(b)
		return true
	})

	return b
}

var _ Lazy = &LazyBuffer{}

func (b *LazyBuffer) isLazy() {}
