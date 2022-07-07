package types

import "reflect"

//

type Pair[L, R any] struct {
	L L
	R R
}

func (p Pair[L, R]) GetL() L { return p.L }
func (p Pair[L, R]) GetR() R { return p.R }

func (p Pair[L, R]) Unpack() (L, R) { return p.L, p.R }

func PairOf[L, R any](l L, r R) Pair[L, R] {
	return Pair[L, R]{l, r}
}

//

func (p Pair[L, R]) ReflectTypeArgs() []reflect.Type {
	return []reflect.Type{reflect.TypeOf(p.L), reflect.TypeOf(p.R)}
}
