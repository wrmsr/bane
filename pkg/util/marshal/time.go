package marshal

import (
	"fmt"
	"reflect"
	"time"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

//

type TimeMarshaler struct {
	layout string
}

func NewTimeMarshaler(layout string) TimeMarshaler {
	return TimeMarshaler{layout: layout}
}

var _ Marshaler = TimeMarshaler{}

func (m TimeMarshaler) Marshal(ctx MarshalContext, rv reflect.Value) (Value, error) {
	if !rv.Type().AssignableTo(rfl.Time()) {
		return nil, _unhandledType
	}

	t := rv.Interface().(time.Time)
	return String{v: t.Format(m.layout)}, nil
}

//

type TimeUnmarshaler struct {
	layouts []string
}

func NewTimeUnmarshaler(layouts []string) TimeUnmarshaler {
	return TimeUnmarshaler{layouts: layouts}
}

var _ Unmarshaler = TimeUnmarshaler{}

func (u TimeUnmarshaler) Unmarshal(ctx UnmarshalContext, mv Value) (reflect.Value, error) {
	switch mv := mv.(type) {

	case Null:
		return rfl.ZeroTime(), nil

	case String:
		for _, l := range u.layouts {
			if t, err := time.Parse(l, mv.v); err == nil {
				return reflect.ValueOf(t), nil
			}
		}
		return rfl.Invalid(), fmt.Errorf("cannot parse time: %s", mv.v)

	}
	return rfl.Invalid(), _unhandledType
}