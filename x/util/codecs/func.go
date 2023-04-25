package codecs

type FuncCodec[F, T any] struct {
	e func(F) T
	d func(T) F
}

var _ Codec[int8, int32] = FuncCodec[int8, int32]{}

func (c FuncCodec[F, T]) Encode(f F) T {
	return c.e(f)
}

func (c FuncCodec[F, T]) Decode(t T) F {
	return c.d(t)
}

func Of[F, T any](e func(F) T, d func(T) F) FuncCodec[F, T] {
	return FuncCodec[F, T]{e: e, d: d}
}
