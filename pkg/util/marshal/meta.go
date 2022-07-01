package marshal

import (
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type MarshalPredicate = func(ctx MarshalContext, rv reflect.Value) bool

func TrueMarshalPredicate(ctx MarshalContext, rv reflect.Value) bool {
	return true
}

func PredicatedMarshaler(fn MarshalPredicate, m Marshaler) bt.Pair[MarshalPredicate, Marshaler] {
	return bt.PairOf(fn, m)
}

type SwitchingMarshaler struct {
	s []bt.Pair[MarshalPredicate, Marshaler]
}

func NewSwitchingMarshaler(s ...bt.Pair[MarshalPredicate, Marshaler]) SwitchingMarshaler {
	return SwitchingMarshaler{s: s}
}

var _ Marshaler = SwitchingMarshaler{}

func (m SwitchingMarshaler) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	for _, e := range m.s {
		if e.L(ctx, rv) {
			return e.R.Marshal(ctx, rv)
		}
	}
	return nil, _unhandledType
}

//

type UnmarshalPredicate = func(ctx UnmarshalContext, mv Value) bool

func TrueUnmarshalPredicate(ctx UnmarshalContext, mv Value) bool {
	return true
}

func PredicatedUnmarshaler(fn UnmarshalPredicate, u Unmarshaler) bt.Pair[UnmarshalPredicate, Unmarshaler] {
	return bt.PairOf(fn, u)
}

type SwitchingUnmarshaler struct {
	s []bt.Pair[UnmarshalPredicate, Unmarshaler]
}

func NewSwitchingUnmarshaler(s ...bt.Pair[UnmarshalPredicate, Unmarshaler]) SwitchingUnmarshaler {
	return SwitchingUnmarshaler{s: s}
}

var _ Unmarshaler = SwitchingUnmarshaler{}

func (m SwitchingUnmarshaler) Unmarshal(ctx UnmarshalContext, mv Value) (reflect.Value, error) {
	for _, e := range m.s {
		if e.L(ctx, mv) {
			return e.R.Unmarshal(ctx, mv)
		}
	}
	return rfl.Invalid(), _unhandledType
}
