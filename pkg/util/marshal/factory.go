package marshal

import (
	"reflect"
)

///

type MarshalerFactoryContext struct {
	Factory func(ctx MarshalerFactoryContext, ty reflect.Type) (Marshaler, error)
}

type MarshalerFactory interface {
	MakeMarshaler(ctx MarshalerFactoryContext, ty reflect.Type) (Marshaler, error)
}

//

type FuncMarshalerFactory struct {
	fn func(ctx MarshalerFactoryContext, ty reflect.Type) (Marshaler, error)
}

func NewFuncMarshalerFactory(fn func(ctx MarshalerFactoryContext, ty reflect.Type) (Marshaler, error)) FuncMarshalerFactory {
	return FuncMarshalerFactory{fn: fn}
}

var _ MarshalerFactory = FuncMarshalerFactory{}

func (mf FuncMarshalerFactory) MakeMarshaler(ctx MarshalerFactoryContext, ty reflect.Type) (Marshaler, error) {
	return mf.fn(ctx, ty)
}

//

type SimpleMarshalerFactory struct {
	m map[reflect.Type]Marshaler
}

func NewSimpleMarshalerFactory(m map[reflect.Type]Marshaler) SimpleMarshalerFactory {
	return SimpleMarshalerFactory{m: m}
}

var _ MarshalerFactory = SimpleMarshalerFactory{}

func (mf SimpleMarshalerFactory) MakeMarshaler(ctx MarshalerFactoryContext, ty reflect.Type) (Marshaler, error) {
	if m, ok := mf.m[ty]; ok {
		return m, nil
	}
	return nil, nil
}

///

type UnmarshalerFactoryContext struct {
	Factory func(ctx UnmarshalerFactoryContext, ty reflect.Type) (Unmarshaler, error)
}

type UnmarshalerFactory interface {
	MakeUnmarshaler(ctx UnmarshalerFactoryContext, ty reflect.Type) (Unmarshaler, error)
}

//

type FuncUnmarshalerFactory struct {
	fn func(ctx UnmarshalerFactoryContext, ty reflect.Type) (Unmarshaler, error)
}

func NewFuncUnmarshalerFactory(fn func(ctx UnmarshalerFactoryContext, ty reflect.Type) (Unmarshaler, error)) FuncUnmarshalerFactory {
	return FuncUnmarshalerFactory{fn: fn}
}

var _ UnmarshalerFactory = FuncUnmarshalerFactory{}

func (uf FuncUnmarshalerFactory) MakeUnmarshaler(ctx UnmarshalerFactoryContext, ty reflect.Type) (Unmarshaler, error) {
	return uf.fn(ctx, ty)
}

//

type SimpleUnmarshalerFactory struct {
	m map[reflect.Type]Unmarshaler
}

func NewSimpleUnmarshalerFactory(m map[reflect.Type]Unmarshaler) SimpleUnmarshalerFactory {
	return SimpleUnmarshalerFactory{m: m}
}

var _ UnmarshalerFactory = SimpleUnmarshalerFactory{}

func (uf SimpleUnmarshalerFactory) MakeUnmarshaler(ctx UnmarshalerFactoryContext, ty reflect.Type) (Unmarshaler, error) {
	if u, ok := uf.m[ty]; ok {
		return u, nil
	}
	return nil, nil
}
