package codecs

type InvertCodec[F, T any] struct {
	c Codec[T, F]
}

var _ Codec[int8, int32] = InvertCodec[int8, int32]{}

func (c InvertCodec[F, T]) Encode(f F) T {
	return c.c.Decode(f)
}

func (c InvertCodec[F, T]) Decode(t T) F {
	return c.c.Encode(t)
}

func Invert[F, T any](c Codec[F, T]) InvertCodec[T, F] {
	return InvertCodec[T, F]{c: c}
}
