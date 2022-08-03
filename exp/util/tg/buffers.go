package tg

type Buffer struct {
	s []float32
}

func (b *Buffer) UnaryOp(op Op) *Buffer {
	panic("nyi")
}

func (b *Buffer) BinaryOp(op Op, y *Buffer) *Buffer {
	panic("nyi")
}
