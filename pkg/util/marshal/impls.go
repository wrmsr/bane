package marshal

import "reflect"

//

type PrimitiveMarshaller struct{}

var _ Marshaller = PrimitiveMarshaller{}

func (p PrimitiveMarshaller) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	//TODO implement me
	panic("implement me")
}

//

type PrimitiveUnmarshaller struct{}

var _ Unmarshaller = PrimitiveUnmarshaller{}

func (p PrimitiveUnmarshaller) Unmarshal(ctx MarshalContext, mv Value) (reflect.Value, error) {
	//TODO implement me
	panic("implement me")
}
