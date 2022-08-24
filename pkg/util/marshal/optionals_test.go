package marshal

import (
	"reflect"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TestOptionals(t *testing.T) {
	m := NewOptionalMarshaler(rfl.TypeOf[bt.Optional[int]](), PrimitiveMarshaler{})
	u := NewOptionalUnmarshaler(rfl.TypeOf[bt.Optional[int]](), NewConvertUnmarshaler(rfl.TypeOf[int](), PrimitiveUnmarshaler{}))

	for _, v := range []bt.Optional[int]{
		bt.Just(10),
		bt.None[int](),
	} {
		mv := check.Must1(m.Marshal(MarshalContext{}, reflect.ValueOf(v)))
		v2 := check.Must1(u.Unmarshal(UnmarshalContext{}, mv)).Interface().(bt.Optional[int])
		tu.AssertDeepEqual(t, v, v2)
	}
}

func TestOptionalsFactory(t *testing.T) {
	m := check.Must1(
		optionalMarshalerFactory.Make(
			MarshalContext{
				Make: NewPrimitiveMarshalerFactory().Make,
			},
			rfl.TypeOf[bt.Optional[int]](),
		),
	)

	u := check.Must1(
		optionalUnmarshalerFactory.Make(
			UnmarshalContext{
				Make: NewConvertPrimitiveUnmarshalerFactory().Make,
			},
			rfl.TypeOf[bt.Optional[int]](),
		),
	)

	for _, v := range []bt.Optional[int]{
		bt.Just(10),
		bt.None[int](),
	} {
		mv := check.Must1(m.Marshal(MarshalContext{}, reflect.ValueOf(v)))
		v2 := check.Must1(u.Unmarshal(UnmarshalContext{}, mv)).Interface().(bt.Optional[int])
		tu.AssertDeepEqual(t, v, v2)
	}
}
