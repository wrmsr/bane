package tg

import (
	"github.com/wrmsr/bane/pkg/util/check"
	"github.com/wrmsr/bane/pkg/util/slices"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

type Buffer struct {
	shape Shape
	s     []float32
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

func (b *Buffer) MovementOp(op Op, arg any) *Buffer {
	switch op {
	case ExpandOp:
		return b.Expand(arg.(Shape))
	case ReshapeOp:
		return b.Reshape(arg.(Shape))
	}
	panic("nyi")
}
