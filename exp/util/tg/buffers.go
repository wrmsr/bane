package tg

import (
	"github.com/wrmsr/bane/pkg/util/check"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

type Buffer struct {
	shape Shape
	s     []float32

	st opt.Optional[Strides]
}

func NewBuffer(shape Shape) *Buffer {
	return &Buffer{
		shape: shape,
		s:     make([]float32, int(shape.Dim())),
	}
}

func BufferOf(shape Shape, s []float32) *Buffer {
	check.Condition(int(shape.Dim()) == len(s))
	return &Buffer{
		shape: shape,
		s:     s,
	}
}

func (b *Buffer) Shape() Shape { return b.shape }

func (b *Buffer) Strides() Strides {
	return opt.SetIfAbsent(&b.st, func() Strides {
		return StridesForShape(b.shape)
	})
}

func (b *Buffer) Get(idxs ...Dim) float32 {
	return b.s[b.Strides().Offset(idxs...)]
}

func (b *Buffer) set(v float32, idxs ...Dim) {
	b.s[b.Strides().Offset(idxs...)] = v
}

func (b *Buffer) UnaryOp(op Op) *Buffer {
	switch op {
	case ReluOp:
		z := NewBuffer(b.shape)
		for i, x := range b.s {
			if !(x < 0) {
				z.s[i] = x
			}
		}
		return z
	case SignOp:
		z := NewBuffer(b.shape)
		for i, x := range b.s {
			if x > 0 {
				z.s[i] = 1
			} else if x < 0 {
				z.s[i] = -1
			}
		}
		return z

	}
	panic("nyi")
}

func (b *Buffer) BinaryOp(op Op, y *Buffer) *Buffer {
	switch op {
	case AddOp:
		check.Condition(b.shape.Equals(y.shape))
		z := NewBuffer(b.shape)
		for i, x := range b.s {
			z.s[i] = x + y.s[i]
		}
		return z
	case SubOp:
		check.Condition(b.shape.Equals(y.shape))
		z := NewBuffer(b.shape)
		for i, x := range b.s {
			z.s[i] = x - y.s[i]
		}
		return z
	case MulOp:
		check.Condition(b.shape.Equals(y.shape))
		z := NewBuffer(b.shape)
		for i, x := range b.s {
			z.s[i] = x * y.s[i]
		}
		return z
	case DivOp:
		check.Condition(b.shape.Equals(y.shape))
		z := NewBuffer(b.shape)
		for i, x := range b.s {
			z.s[i] = x / y.s[i]
		}
		return z
	}
	panic("nyi")
}

func MakeConstBuffer(c float32, shape Shape) *Buffer {
	if len(shape) < 1 {
		shape = scalarShape
	}
	s := make([]float32, shape.Dim())
	for i := range s {
		s[i] = c
	}
	return BufferOf(shape, s)
}

func (b *Buffer) Expand(newShape Shape) *Buffer {
	if b.shape.Dim() != 1 {
		panic("nyi")
	}
	return MakeConstBuffer(b.s[0], newShape)
}

func (b *Buffer) Reshape(newShape Shape) *Buffer {
	check.Condition(bt.Prod[Dim](b.shape...) == bt.Prod[Dim](newShape...))
	if slices.Equal(b.shape, newShape) {
		return b
	}
	return BufferOf(newShape, b.s)
}

func (b *Buffer) Transpose(order []Dim) *Buffer {
	if len(b.shape) != 2 || len(order) != 2 || order[0] != 1 || order[1] != 0 {
		panic("nyi")
	}
	r := BufferOf(Shape{b.shape[1], b.shape[0]}, make([]float32, b.shape.Dim()))
	for i := Dim(0); i < b.shape[0]; i++ {
		for j := Dim(0); j < b.shape[1]; j++ {
			r.set(b.Get(i, j), j, i)
		}
	}
	return r
}

func (b *Buffer) Permute(order []Dim) *Buffer {
	return b.Transpose(order)
}

func (b *Buffer) Pad(padding []SliceBound) *Buffer {
	check.Condition(len(padding) == len(b.shape))
	if slices.All(padding, func(b SliceBound) bool {
		return b.Start == 0 && b.Stop == 0
	}) {
		return b
	}

	check.Condition(len(padding) == 2)

	newShape := slices.Clone(b.shape)
	for i, p := range padding {
		newShape[i] += p.Start + p.Stop
	}

	// return np.pad(x, padding)
	panic("nyi")
}

func (b *Buffer) Slice(bounds ...SliceBound) *Buffer {
	check.Condition(len(bounds) == len(b.shape))
	var nop = true
	for i, bo := range bounds {
		if bo.Start != 0 || bo.Stop != b.shape[i] {
			nop = false
			break
		}
	}
	if nop {
		return b
	}

	panic("nyi")
}

type SliceBound struct {
	Start, Stop Dim
}

func (b *Buffer) MovementOp(op Op, arg any) *Buffer {
	switch op {
	case ExpandOp:
		return b.Expand(arg.(Shape))
	case ReshapeOp:
		return b.Reshape(arg.(Shape))
	case PermuteOp:
		return b.Permute(arg.([]Dim))
	case SliceOp:
		bs := arg.([]SliceBound)
		padding := make([]SliceBound, len(bs))
		for i, p := range bs {
			padding[i] = SliceBound{bt.Max[Dim](0, -p.Start), bt.Max[Dim](0, p.Stop-b.shape[i])}
		}
		xp := b.Pad(padding)
		ps := make([]SliceBound, len(bs))
		for i, p := range bs {
			ps[i] = SliceBound{
				p.Start + padding[i].Start,
				p.Stop + padding[i].Start,
			}
		}
		return xp.Slice(ps...)
	}
	panic("nyi")
}

func (b *Buffer) ProcessingOp(op Op, w *Buffer, arg any) *Buffer {
	switch op {
	case ConvOp:
		ca := arg.(ConvArgs)
		x := b.MovementOp(
			SliceOp,
			[]SliceBound{
				{0, b.shape[0]},
				{0, b.shape[1]},
				{-ca.py, b.shape[2] + ca.py_},
				{-ca.px, b.shape[3] + ca.px_},
			},
		)
		_ = x
		/*
		   gx = x.ravel().Reshape(ca.bs, ca.groups, ca.cin, x.shape[2], x.shape[3])
		   tx = np.lix.stride_tricks.as_strided(
		       gx,
		       shape=(
		           ca.bs,
		           ca.groups,
		           ca.cin,
		           ca.oy,
		           ca.ox,
		           ca.H,
		           ca.W,
		       ),
		       strides=(
		           *gx.strides[0:3],
		           gx.strides[3] * ca.sy,
		           gx.strides[4] * ca.sx,
		           gx.strides[3] * ca.dy,
		           gx.strides[4] * ca.dx,
		       ),
		       writeable=False,
		   )
		   tw = w.reshape(ca.groups, ca.rcout, ca.cin, ca.H, ca.W)
		   tmp = np.empty((ca.bs, ca.groups, ca.oy, ca.ox, ca.rcout), dtype=x.dtype)
		   for g in range(ca.groups):
		       # ijYXyx,kjyx -> iYXk ->ikYX
		       tmp[:, g] = np.tensordot(tx[:, g], tw[g], ((1, 4, 5), (1, 2, 3)))
		   return np.moveaxis(tmp, 4, 2).reshape(ca.bs, ca.groups * ca.rcout, ca.oy, ca.ox).view(CPUBuffer)
		*/
		panic("nyi")
	}
	panic("nyi")
}
