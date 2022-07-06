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

func (mf SimpleMarshalerFactory) MakeMarshaler(ctx MarshalerFactoryContext, ty reflect.Type) (Marshaler, error) {
	if _, ok := mf.tys[ty]; !ok {
		return nil, nil
	}
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

func (uf SimpleUnmarshalerFactory) MakeUnmarshaler(ctx UnmarshalerFactoryContext, ty reflect.Type) (Unmarshaler, error) {
	if _, ok := mf.tys[ty]; !ok {
		return nil, nil
	}
	panic("implement me")
}
