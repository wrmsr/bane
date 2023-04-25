package types

//

type AnyBox interface {
	Get() any
}

type MutAnyBox interface {
	AnyBox

	Set(v any)
}

//

type Box[T any] interface {
	Get() T
}

type MutBox[T any] interface {
	Box[T]

	Set(v T)
}

//

type BoxImpl[T any] struct {
	v T
}

var _ Box[int] = BoxImpl[int]{}

func (b BoxImpl[T]) Get() T {
	return b.v
}

//

type MutBoxImpl[T any] struct {
	b BoxImpl[T]
}

var _ MutBox[int] = &MutBoxImpl[int]{}

func (b *MutBoxImpl[T]) Get() T { return b.b.Get() }

func (b *MutBoxImpl[T]) Set(v T) {
	b.b.v = v
}
