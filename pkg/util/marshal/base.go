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
	Make func(ctx MarshalContext, ty reflect.Type) (Marshaler, error)
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

//

type MarshalerFactory = Factory[Marshaler, MarshalContext, reflect.Type]

func NewTypeMapMarshalerFactory(m map[reflect.Type]Marshaler) MarshalerFactory {
	return NewTypeMapFactory[Marshaler, MarshalContext](m)
}

func NewTypeCacheMarshalerFactory(f MarshalerFactory) MarshalerFactory {
	return NewTypeCacheFactory[Marshaler, MarshalContext](f)
}

///

type UnmarshalOpt interface {
	isUnmarshalOpt()
}

type BaseUnmarshalOpt struct{}

func (o BaseUnmarshalOpt) isUnmarshalOpt() {}

//

type UnmarshalContext struct {
	Make func(ctx UnmarshalContext, ty reflect.Type) (Unmarshaler, error)
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

//

type UnmarshalerFactory = Factory[Unmarshaler, UnmarshalContext, reflect.Type]

func NewTypeMapUnmarshalerFactory(m map[reflect.Type]Unmarshaler) UnmarshalerFactory {
	return NewTypeMapFactory[Unmarshaler, UnmarshalContext](m)
}

func NewTypeCacheUnmarshalerFactory(f UnmarshalerFactory) UnmarshalerFactory {
	return NewTypeCacheFactory[Unmarshaler, UnmarshalContext](f)
}
