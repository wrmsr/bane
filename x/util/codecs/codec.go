package codecs

type Encoder[F, T any] interface {
	Encode(F) T
}

type Decoder[F, T any] interface {
	Decode(T) F
}

type Codec[F, T any] interface {
	Encoder[F, T]
	Decoder[F, T]
}
