package marshal

import (
	"reflect"

	ctr "github.com/wrmsr/bane/pkg/util/container"
)

///

type UnhandledTypeError struct{}

func (e UnhandledTypeError) Error() string {
	return "unhandled type"
}

var _unhandledType = UnhandledTypeError{}

func UnhandledType() UnhandledTypeError { return _unhandledType }

///

type MarshalOpt interface {
	isMarshalOpt()
}

type BaseMarshalOpt struct{}

func (o BaseMarshalOpt) isMarshalOpt() {}

//

type MarshalContext struct {
	Opts ctr.TypeMap[MarshalOpt]
}

type Marshaler interface {
	Marshal(ctx MarshalContext, rv reflect.Value) (Value, error)
}

//

type FuncMarshaler struct {
	fn func(ctx MarshalContext, rv reflect.Value) (Value, error)
}

func NewFuncMarshaler(fn func(ctx MarshalContext, rv reflect.Value) (Value, error)) FuncMarshaler {
	return FuncMarshaler{fn: fn}
}

var _ Marshaler = FuncMarshaler{}

func (m FuncMarshaler) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	return m.fn(ctx, rv)
}

///

type UnmarshalOpt interface {
	isUnmarshalOpt()
}

type BaseUnmarshalOpt struct{}

func (o BaseUnmarshalOpt) isUnmarshalOpt() {}

//

type UnmarshalContext struct {
	Opts ctr.TypeMap[UnmarshalOpt]
}

type Unmarshaler interface {
	Unmarshal(ctx UnmarshalContext, mv Value) (reflect.Value, error)
}

//

type FuncUnmarshaler struct {
	fn func(ctx UnmarshalContext, mv Value) (reflect.Value, error)
}

func NewFuncUnmarshaler(fn func(ctx UnmarshalContext, mv Value) (reflect.Value, error)) FuncUnmarshaler {
	return FuncUnmarshaler{fn: fn}
}

var _ Unmarshaler = FuncUnmarshaler{}

func (u FuncUnmarshaler) Unmarshal(ctx UnmarshalContext, mv Value) (reflect.Value, error) {
	return u.fn(ctx, mv)
}
