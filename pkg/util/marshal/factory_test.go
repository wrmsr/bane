package marshal

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TestDefaultFactory(t *testing.T) {
	mf := NewCompositeFactory(
		FirstComposite,
		NewPrimitiveMarshalerFactory(),
		NewPointerMarshalerFactory(),
		NewIndexMarshalerFactory(),
		NewMapMarshalerFactory(),
		NewBase64MarshalerFactory(),
		NewOptionalMarshalerFactory(),
	)

	uf := NewCompositeFactory(
		FirstComposite,
		NewConvertPrimitiveUnmarshalerFactory(),
		NewPointerUnmarshalerFactory(),
		NewIndexUnmarshalerFactory(),
		NewMapUnmarshalerFactory(),
		NewBase64UnmarshalerFactory(),
		NewOptionalUnmarshalerFactory(),
	)

	o := map[opt.Optional[int]][]*string{
		opt.Just(10):    {bt.PtrTo("ten"), bt.PtrTo("one zero")},
		opt.Just(20):    {bt.PtrTo("twenty"), bt.PtrTo("two zero")},
		opt.None[int](): {bt.PtrTo("none")},
	}
	fmt.Println(o)

	m := check.Must1(mf.Make(MarshalerFactoryContext{Make: mf.Make}, reflect.TypeOf(o)))
	u := check.Must1(uf.Make(UnmarshalerFactoryContext{Make: uf.Make}, reflect.TypeOf(o)))

	v := check.Must1(m.Marshal(MarshalContext{}, reflect.ValueOf(o)))
	fmt.Println(v)

	o2 := check.Must1(u.Unmarshal(UnmarshalContext{}, v))
	fmt.Println(o2)

	//tu.AssertDeepEqual(t, o, o2)
}
