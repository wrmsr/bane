package marshal

import (
	"reflect"

	"github.com/wrmsr/bane/pkg/util/maps"
)

//

type MarshalerFactoryContext struct {
}

type MarshalerFactory interface {
	MakeMarshaler(ctx MarshalerFactoryContext, ty reflect.Type) (Marshaler, error)
}

//

type SimpleMarshalerFactory struct {
	tys map[reflect.Type]struct{}
	m   Marshaler
}

func NewSimpleMarshalerFactory(tys []reflect.Type, m Marshaler) SimpleMarshalerFactory {
	return SimpleMarshalerFactory{tys: maps.NewTypeSet(tys), m: m}
}

var _ MarshalerFactory = SimpleMarshalerFactory{}

func (s SimpleMarshalerFactory) MakeMarshaler(ctx MarshalerFactoryContext, ty reflect.Type) (Marshaler, error) {
	panic("implement me")
}

//

type UnmarshalerFactoryContext struct {
}

type UnmarshalerFactory interface {
	MakeUnmarshaler(ctx UnmarshalerFactoryContext, ty reflect.Type) (Unmarshaler, error)
}

//

type SimpleUnmarshalerFactory struct {
	tys map[reflect.Type]struct{}
	u   Unmarshaler
}

var _ UnmarshalerFactory = SimpleUnmarshalerFactory{}

func (s SimpleUnmarshalerFactory) MakeUnmarshaler(ctx UnmarshalerFactoryContext, ty reflect.Type) (Unmarshaler, error) {
	panic("implement me")
}
