package tg

import (
	"math"

	"github.com/wrmsr/bane/core/check"
	nd "github.com/wrmsr/bane/core/ndarray"
	"github.com/wrmsr/bane/core/slices"
	bt "github.com/wrmsr/bane/core/types"
)

type Buffer struct {
	s []float32

	shape Shape

	strides bt.Optional[Strides]
}

func NewBuffer(shape Shape) *Buffer {
	return &Buffer{
		s: make([]float32, int(shape.Dim())),

		shape: shape,
	}
}

func BufferOf(shape Shape, s []float32) *Buffer {
	check.Condition(int(shape.Dim()) == len(s))
	return &Buffer{
		s: s,

		shape: shape,
	}
}

func BufferOfNd(a nd.NdArray[float32]) *Buffer {
	check.Condition(a.View().Offset() == 0)
	return &Buffer{
		s:       a.Data(),
		shape:   a.View().Shape().Slice(),
		strides: bt.Just(Strides(a.View().Strides().Slice())),
	}
}

func (b *Buffer) Shape() Shape { return b.shape }

func (b *Buffer) Strides() Strides {
	return bt.SetIfAbsent(&b.strides, func() Strides {
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

	case ExpOp:
		z := NewBuffer(b.shape)
		for i, x := range b.s {
			z.s[i] = float32(math.Exp(float64(x)))
		}
		return z

	case LogOp:
		z := NewBuffer(b.shape)
		for i, x := range b.s {
			z.s[i] = float32(math.Log(float64(x)))
		}
		return z

	case NegOp:
		z := NewBuffer(b.shape)
		for i, x := range b.s {
			z.s[i] = -x
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

	case CmpEqOp:
		check.Condition(b.shape.Equals(y.shape))
		z := NewBuffer(b.shape)
		for i, x := range b.s {
			if x == y.s[i] {
				z.s[i] = 1.
			}
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
	if b.shape.Dim() == 1 {
		return MakeConstBuffer(b.s[0], newShape)
	}
	if len(b.shape) == 2 && len(newShape) == 2 && b.shape[0] == newShape[0] && b.shape[1] == 1 {
		r := NewBuffer(newShape)
		for i := Dim(0); i < b.shape[0]; i++ {
			v := b.Get(i, 0)
			for j := Dim(0); j < newShape[1]; j++ {
				r.set(v, i, j)
			}
		}
		return r
	}
	panic("nyi")
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
		t := b.Nd().Transpose(nd.DimsOf(order...))
		return BufferOfNd(t)
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

func (b *Buffer) Flip(axes []Dim) *Buffer {
	q := make([]bool, len(b.shape))
	for _, a := range axes {
		q[a] = true
	}

	r := BufferOf(b.shape, make([]float32, b.shape.Dim()))

	bsl := make([]Dim, len(b.shape))
	rsl := make([]Dim, len(b.shape))

	var rec func(int)
	rec = func(a int) {
		if a < len(b.shape) {
			l := b.shape[a]
			for i := Dim(0); i < l; i++ {
				if q[a] {
					bsl[a] = l - i - 1
				} else {
					bsl[a] = i
				}
				rsl[a] = i
				rec(a + 1)
			}
		} else {
			r.set(b.Get(bsl...), rsl...)
		}
	}
	rec(0)
	return r
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

// dot(a, b)[i,j,k,m] = sum(a[i,j,:] * b[k,:,m])
// tensordot, really
func (b *Buffer) Dot(y *Buffer, baxes, yaxes []int) *Buffer {
	if !slices.Equal(baxes, []int{1, 0}) || !slices.Equal(yaxes, []int{0, 1}) {
		panic("nyi")
	}

	if len(b.shape) != 3 || len(y.shape) != 3 {
		panic("nyi")
	}

	if b.shape[0] != y.shape[1] ||
		b.shape[1] != y.shape[0] {
		panic("nyi")
	}

	ret := NewBuffer(Shape{b.shape[2], y.shape[2]})
	for i := Dim(0); i < b.shape[2]; i++ {
		for j := Dim(0); j < y.shape[2]; j++ {
			for k := Dim(0); k < b.shape[0]; k++ {
				for n := Dim(0); n < y.shape[0]; n++ {
					ret.set(ret.Get(i, j)+b.Get(k, n, i)*y.Get(n, k, j), i, j)
				}
			}
		}
	}

	return ret
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

	case FlipOp:
		return b.Flip(arg.([]Dim))

	}
	panic("nyi")
}

func (b *Buffer) Nd() nd.NdArray[float32] {
	return nd.Maker[float32]{
		Shape:   nd.ShapeOf(b.shape...),
		Strides: nd.StridesOf(b.Strides()...),
		Data:    b.s,
	}.Make()
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
		gx := x.Reshape(Shape{ca.bs, ca.groups, ca.cin, x.shape[2], x.shape[3]})
		gst := gx.Strides()
		tx := Buffer{
			s: b.s,
			shape: Shape{
				ca.bs,
				ca.groups,
				ca.cin,
				ca.oy,
				ca.ox,
				ca.h,
				ca.w,
			},
			strides: bt.Just[Strides](Strides{
				gst[0],
				gst[1],
				gst[2],
				gst[3] * ca.sy,
				gst[4] * ca.sx,
				gst[3] * ca.dy,
				gst[4] * ca.dx,
			}),
		}
		tw := w.Reshape(Shape{ca.groups, ca.rcout, ca.cin, ca.h, ca.w})

		tmp := NewBuffer(Shape{ca.bs, ca.groups, ca.oy, ca.ox, ca.rcout})

		ndtx := tx.Nd()
		ndtw := tw.Nd()
		ndtmp := tmp.Nd()

		for g := Dim(0); g < ca.groups; g++ {
			txg := ndtx.Slice(nil, g)
			twg := ndtw.Slice(g)
			q := NdTensorDot(txg, twg, nd.DimsOf(1, 4, 5), nd.DimsOf(1, 2, 3))
			ndtmp.Slice(nil, g).Assign(q)
		}

		tmp2 := ndtmp.MoveAxis(4, 2)
		tmp3 := tmp2.Reshape(nd.ShapeOf(ca.bs, ca.groups*ca.rcout, ca.oy, ca.ox))

		return BufferOfNd(tmp3)
	}
	panic("nyi")
}

func (b *Buffer) reduce1(axis int, fn func(float32, bt.Optional[float32]) float32) *Buffer {
	bnd := b.Nd()

	mnsh := bnd.View().Shape().Mutate()
	mnsh.Set(axis, 1)
	nsh := nd.Shape{mnsh.Decay()}

	o := nsh.Order()
	sl := nd.NewMutDims(o)
	c := nd.New[float32](nsh)

	var rec func(int)
	rec = func(i int) {
		if i == axis {
			rec(i + 1)
		} else if i < o {
			n := bnd.View().Shape().Get(i)
			for j := nd.Dim(0); j < n; j++ {
				sl.Set(i, j)
				rec(i + 1)
			}
		} else {
			n := bnd.View().Shape().Get(axis)
			var r bt.Optional[float32]
			for j := nd.Dim(0); j < n; j++ {
				sl.Set(axis, j)
				v := bnd.Get(sl.Decay())
				r = bt.Just(fn(v, r))
			}
			sl.Set(axis, 0)
			*c.At(sl.Decay()) = r.Value()
		}
	}
	rec(0)

	return BufferOfNd(c)
}

func (b *Buffer) reduce(axes []int, fn func(float32, bt.Optional[float32]) float32) *Buffer {
	r := b
	for _, axis := range axes {
		r = r.reduce1(axis, fn)
	}
	return r
}

func (b *Buffer) ReduceOp(op Op, nsh Shape) *Buffer {
	check.Equal(len(b.shape), len(nsh))

	var axis []int
	for i := 0; i < len(nsh); i++ {
		if b.shape[i] != nsh[i] {
			axis = append(axis, i)
		}
	}

	if len(axis) < 1 {
		return b
	}

	switch op {

	case MaxOp:
		return b.reduce(axis, func(f float32, b bt.Optional[float32]) float32 {
			if b.Present() && !(f > b.Value()) {
				return b.Value()
			}
			return f
		})

	case SumOp:
		return b.reduce(axis, func(f float32, b bt.Optional[float32]) float32 {
			return f + b.Or(0)
		})

	}
	panic("nyi")
}
