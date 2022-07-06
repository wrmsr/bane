package marshal

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	ctr "github.com/wrmsr/bane/pkg/util/container"
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

func TestOptionalsFactory(t *testing.T) {
	fmt.Println(rfl.TypeOf[opt.Optional[ctr.Map[int, opt.Optional[string]]]]().Name())
	fmt.Println(rfl.TypeOf[opt.Optional[ctr.Map[int, opt.Optional[string]]]]().Name())

	m := check.Must1(optionalMarshalerFactory.MakeMarshaler(MarshalerFactoryContext{}, rfl.TypeOf[opt.Optional[int]]()))
	fmt.Println(m)
}
