package marshal

import (
	"reflect"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

func TestOptionals(t *testing.T) {
	m := NewOptionalMarshaler(rfl.TypeOf[opt.Optional[int]](), PrimitiveMarshaler{})
	u := NewOptionalUnmarshaler(rfl.TypeOf[opt.Optional[int]](), NewConvertUnmarshaler(rfl.TypeOf[int](), PrimitiveUnmarshaler{}))

	for _, v := range []opt.Optional[int]{
		opt.Just(10),
		opt.None[int](),
	} {
		mv := check.Must1(m.Marshal(MarshalContext{}, reflect.ValueOf(v)))
		v2 := check.Must1(u.Unmarshal(UnmarshalContext{}, mv)).Interface().(opt.Optional[int])
		tu.AssertDeepEqual(t, v, v2)
	}
}
