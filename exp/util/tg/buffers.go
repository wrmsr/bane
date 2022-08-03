package tg

import "github.com/wrmsr/bane/pkg/util/check"

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
	case MulOp:
		check.Condition(b.shape.Equals(y.shape))
		z := NewBuffer(b.shape)
		for i, x := range b.s {
			z.s[i] = x * y.s[i]
		}
		return z
	}
	panic("nyi")
}
