package tg

import "github.com/wrmsr/bane/pkg/util/check"

type Buffer struct {
	shape Shape
	s     []float32
}

func NewBuffer(shape Shape, s []float32) *Buffer {
	check.Condition(int(shape.Dim()) == len(s))
	return &Buffer{shape: shape, s: s}
}

func (b *Buffer) UnaryOp(op Op) *Buffer {
	panic("nyi")
}

func (b *Buffer) BinaryOp(op Op, y *Buffer) *Buffer {
	panic("nyi")
}
