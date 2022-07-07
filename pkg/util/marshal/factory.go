package marshal

import (
	"reflect"
)

///

type Factory[R, C, A any] interface {
	Make(ctx C, a A) (R, error)
}

//

type FuncFactory[R, C, A any] struct {
	fn func(ctx C, a A) (R, error)
}

func NewFuncFactory[R, C, A any](fn func(ctx C, a A) (R, error)) FuncFactory[R, C, A] {
	return FuncFactory[R, C, A]{fn: fn}
}

var _ Factory[int, uint, string] = FuncFactory[int, uint, string]{}

func (f FuncFactory[R, C, A]) Make(ctx C, a A) (R, error) {
	return f.fn(ctx, a)
}

//

type TypeMapFactory[R, C any] struct {
	m map[reflect.Type]R
}

func NewTypeMapFactory[R, C any](m map[reflect.Type]R) TypeMapFactory[R, C] {
	return TypeMapFactory[R, C]{m: m}
}

var _ Factory[int, uint, reflect.Type] = TypeMapFactory[int, uint]{}

func (f TypeMapFactory[R, C]) Make(ctx C, a reflect.Type) (R, error) {
	if m, ok := f.m[a]; ok {
		return m, nil
	}
	var z R
	return z, nil
}

///

type MarshalerFactoryContext struct {
	Make func(ctx MarshalerFactoryContext, ty reflect.Type) (Marshaler, error)
}

type MarshalerFactory = Factory[Marshaler, MarshalerFactoryContext, reflect.Type]

func NewFuncMarshalerFactory(fn func(ctx MarshalerFactoryContext, ty reflect.Type) (Marshaler, error)) MarshalerFactory {
	return NewFuncFactory(fn)
}

func NewTypeMapMarshalerFactory(m map[reflect.Type]Marshaler) MarshalerFactory {
	return NewTypeMapFactory[Marshaler, MarshalerFactoryContext](m)
}

///

type UnmarshalerFactoryContext struct {
	Make func(ctx UnmarshalerFactoryContext, ty reflect.Type) (Unmarshaler, error)
}

type UnmarshalerFactory = Factory[Unmarshaler, UnmarshalerFactoryContext, reflect.Type]

func NewFuncUnmarshalerFactory(fn func(ctx UnmarshalerFactoryContext, ty reflect.Type) (Unmarshaler, error)) UnmarshalerFactory {
	return NewFuncFactory(fn)
}

func NewTypeMapUnmarshalerFactory(m map[reflect.Type]Unmarshaler) UnmarshalerFactory {
	return NewTypeMapFactory[Unmarshaler, UnmarshalerFactoryContext](m)
}
