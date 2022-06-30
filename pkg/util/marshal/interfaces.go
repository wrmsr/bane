package marshal

import (
	"reflect"

	ctr "github.com/wrmsr/bane/pkg/util/container"
)

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

type Marshaller interface {
	Marshal(ctx MarshalContext, rv reflect.Value) (Value, error)
}

//

type FuncMarshaller struct {
	fn func(ctx MarshalContext, rv reflect.Value) (Value, error)
}

func NewFuncMarshaller(fn func(ctx MarshalContext, rv reflect.Value) (Value, error)) FuncMarshaller {
	return FuncMarshaller{fn: fn}
}

var _ Marshaller = FuncMarshaller{}

func (m FuncMarshaller) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	return m.fn(ctx, rv)
}

//

type MarshallerResolver interface {
	ResolveMarshaller(ty reflect.Type) (Marshaller, error)
}

//

type FuncMarshallerResolver struct {
	fn func(ty reflect.Type) (Marshaller, error)
}

func NewFuncMarshallerResolver(fn func(ty reflect.Type) (Marshaller, error)) FuncMarshallerResolver {
	return FuncMarshallerResolver{fn: fn}
}

var _ MarshallerResolver = FuncMarshallerResolver{}

func (mr FuncMarshallerResolver) ResolveMarshaller(ty reflect.Type) (Marshaller, error) {
	return mr.fn(ty)
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

type Unmarshaller interface {
	Unmarshal(ctx MarshalContext, mv Value) (reflect.Value, error)
}

//

type FuncUnmarshaller struct {
	fn func(ctx MarshalContext, mv Value) (reflect.Value, error)
}

func NewFuncUnmarshaller(fn func(ctx MarshalContext, mv Value) (reflect.Value, error)) FuncUnmarshaller {
	return FuncUnmarshaller{fn: fn}
}

var _ Unmarshaller = FuncUnmarshaller{}

func (u FuncUnmarshaller) Unmarshal(ctx MarshalContext, mv Value) (reflect.Value, error) {
	return u.fn(ctx, mv)
}

//

type UnmarshallerResolver interface {
	ResolveUnmarshaller(ty reflect.Type) (Unmarshaller, error)
}

//

type FuncUnmarshallerResolver struct {
	fn func(ty reflect.Type) (Unmarshaller, error)
}

func NewFuncUnmarshallerResolver(fn func(ty reflect.Type) (Unmarshaller, error)) FuncUnmarshallerResolver {
	return FuncUnmarshallerResolver{fn: fn}
}

var _ UnmarshallerResolver = FuncUnmarshallerResolver{}

func (ur FuncUnmarshallerResolver) ResolveUnmarshaller(ty reflect.Type) (Unmarshaller, error) {
	return ur.fn(ty)
}
