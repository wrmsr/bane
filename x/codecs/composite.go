package codecs

//

type CompositeEncoder[F0, F1, T any] struct {
	e0 Encoder[F0, F1]
	e1 Encoder[F1, T]
}

var _ Encoder[int8, int32] = CompositeEncoder[int8, int16, int32]{}

func (e CompositeEncoder[F0, F1, T]) Encode(f F0) T {
	return e.e1.Encode(e.e0.Encode(f))
}

//

type CompositeDecoder[F0, F1, T any] struct {
	d0 Decoder[F0, F1]
	d1 Decoder[F1, T]
}

var _ Decoder[int8, int32] = CompositeDecoder[int8, int16, int32]{}

func (d CompositeDecoder[F0, F1, T]) Decode(t T) F0 {
	return d.d0.Decode(d.d1.Decode(t))
}

//

type CompositeCodec[F0, F1, T any] struct {
	c0 Codec[F0, F1]
	c1 Codec[F1, T]
}

var _ Codec[int8, int32] = CompositeCodec[int8, int16, int32]{}

func (e CompositeCodec[F0, F1, T]) Encode(f F0) T {
	return e.c1.Encode(e.c0.Encode(f))
}

func (d CompositeCodec[F0, F1, T]) Decode(t T) F0 {
	return d.c0.Decode(d.c1.Decode(t))
}

func Compose[F0, F1, T any](c0 Codec[F0, F1], c1 Codec[F1, T]) CompositeCodec[F0, F1, T] {
	return CompositeCodec[F0, F1, T]{c0: c0, c1: c1}
}
