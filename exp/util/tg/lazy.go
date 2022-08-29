package tg

import "C"
import (
	"fmt"
	"strings"

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
	fmt.Printf(
		"[%s] : %s -> %v\n",
		strings.Join(slices.Map(Op.String, op), ", "),
		strings.Join(slices.Map(func(b *Buffer) string { return fmt.Sprintf("%v", b.Shape()) }, inp), ", "),
		ret.Shape(),
	)
	fmt.Println(ret.Nd())
}

func (b *LazyBuffer) Realize() *Buffer {
	if b.realized != nil {
		return b.realized
	}

	ro := realize[b.ot](b)
	b.realized = ro.data

	var ops []Op
	b.op.ForEachOp(func(op *LazyOp) bool {
		ops = append(ops, op.op)
		return true
	})

	logOp(ro.ot, ops, ro.data, ro.srcs)

	return b.realized
}

//

func ElementwiseOp(op Op, srcs ...*LazyBuffer) *LazyBuffer {
	opSrcs := slices.Map(func(b *LazyBuffer) Lazy { return b }, srcs)
	if len(opSrcs) == 1 { // || MERGE_ELEMENTWISE_OPS
		for i, x := range srcs {
			if x.ot == BinaryOpType && len(x.children) == 0 && x.realized != nil {
				opSrcs[i] = x.op
			}
		}
	}

	return NewLazyBuffer(
		srcs[0].st,
		BinaryOpType,
		&LazyOp{
			op:   op,
			srcs: opSrcs,
		},
	)
}

func (b *LazyBuffer) UnaryOp(op Op) *LazyBuffer {
	return ElementwiseOp(op, b)
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

	// FIXME: ?!?! - does tg relu mutate in place? tg gets transposed strides (st[0]<st[1]), we don't - being transposed
	//   means it's not contiguous so this isn't taken, in ours it *is* taken and the transpose erroneously disappears
	//if b.realized == nil && ret.st.Contiguous() {
	//	root := ret.op.GetBuffers()[0]
	//	if ret.st.Shape().Equals(root.Shape()) {
	//		return root
	//	}
	//}

	return ret
}

func (b *LazyBuffer) ProcessingOp(op Op, w *LazyBuffer, arg any) *LazyBuffer {
	switch op {
	case ConvOp:
		ca := arg.(ConvArgs)
		return NewLazyBuffer(
			NewShapeTracker(ca.outShape),
			ProcessingOpType,
			&LazyOp{
				op:   op,
				srcs: []Lazy{b, w},
				arg:  ca,
			},
		)
	}
	panic(op)
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
