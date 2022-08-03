package tg

import (
	its "github.com/wrmsr/bane/pkg/util/iterators"
	"github.com/wrmsr/bane/pkg/util/maps"
	"github.com/wrmsr/bane/pkg/util/slices"
)

//

type Lazy interface {
	isLazy()
}

//

type LazyOp struct {
	op   Op
	srcs []Lazy
	arg  any
}

var _ Lazy = &LazyOp{}

func (o *LazyOp) isLazy() {}

func (o *LazyOp) ForEachBuffer(fn func(*LazyBuffer) bool) bool {
	for _, s := range o.srcs {
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

func (o *LazyOp) GetBuffers() []*LazyBuffer {
	return its.SeqForEach(its.TraversableOf(o.ForEachBuffer))
}

func (o *LazyOp) ForEachOp(fn func(*LazyOp) bool) bool {
	if !fn(o) {
		return false
	}
	for _, s := range o.srcs {
		switch s := s.(type) {
		case *LazyOp:
			if !fn(s) {
				return false
			}
		}
	}
	return true
}

func (o *LazyOp) GetOps() []*LazyOp {
	return its.SeqForEach(its.TraversableOf(o.ForEachOp))
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

func (b *LazyBuffer) Shape() Shape { return b.st.Shape() }

func logOp(opType OpType, op []Op, ret *Buffer, inp []*Buffer) {

}

func (b *LazyBuffer) Realize() *Buffer {
	if b.realized != nil {
		return b.realized
	}

	ro := realize[b.ot](b)
	b.realized = ro.data

	// logOp(ro.ot, [x.op for x in getLazyOps(self.op)], ro.data, ro.srcs)

	return b.realized
}

//

func ElementwiseOp(op Op, srcs ...*LazyBuffer) *LazyBuffer {
	/*
	   if (MERGE_UNARY_OPS and len(srcs) == 1) or MERGE_ELEMENTWISE_OPS:
	       # remove the buffers from any (childless) BinaryOps that feed into this
	       srcs = tuple(
	           x.op
	           if x.optype == BinaryOps
	              and len(x.children) == 0
	              and x.realized is None else x
	           for x in srcs
	       )
	*/
	return NewLazyBuffer(
		srcs[0].st,
		BinaryOpType,
		&LazyOp{
			op:   op,
			srcs: slices.Map(func(b *LazyBuffer) Lazy { return b }, srcs),
		},
	)
}

func (b *LazyBuffer) BinaryOp(op Op, y *LazyBuffer) *LazyBuffer {
	return ElementwiseOp(op, b, y)
}

func (b *LazyBuffer) ReduceOp(op Op, newShape Shape) *LazyBuffer {
	return NewLazyBuffer(
		NewShapeTracker(newShape),
		ReduceOpType,
		&LazyOp{
			op:   op,
			srcs: []Lazy{b},
			arg:  newShape,
		},
	)
}

func (b *LazyBuffer) MovementOp(op Op, arg any) *LazyBuffer {
	st := b.st.Clone()
	st.MovementOp(op, arg)

	var src Lazy = b
	if op.Type() == MovementOpType && b.realized != nil {
		src = b.op
	}

	ret := NewLazyBuffer(
		st,
		MovementOpType,
		&LazyOp{
			op:   op,
			srcs: []Lazy{src},
			arg:  arg,
		},
	)

	if b.realized != nil && ret.st.Contiguous() {
		root := ret.op.GetBuffers()[0]
		if ret.st.Shape().Equals(root.Shape()) {
			return root
		}
	}

	return ret
}

//

func MakeLoadBuffer(data *Buffer, shape Shape) *LazyBuffer {
	return NewLazyBuffer(
		NewShapeTracker(shape),
		LoadOpType,
		&LazyOp{
			op:   FromCpuOp,
			srcs: nil,
			arg:  data,
		},
	)
}

func MakeLoadConstBuffer(c float32, shape Shape) *LazyBuffer {
	return MakeLoadBuffer(MakeConstBuffer(c, shape), shape)
}
