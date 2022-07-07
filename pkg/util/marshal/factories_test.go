package marshal

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	stu "github.com/wrmsr/bane/pkg/util/structs"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TestDefaultFactory(t *testing.T) {
	sic := stu.NewStructInfoCache()

	mf := NewTypeCacheMarshalerFactory(NewCompositeFactory(
		FirstComposite,
		NewPrimitiveMarshalerFactory(),
		NewPointerMarshalerFactory(),
		NewIndexMarshalerFactory(),
		NewMapMarshalerFactory(),
		NewBase64MarshalerFactory(),
		NewTimeMarshalerFactory(DefaultTimeMarshalLayout),
		NewOptionalMarshalerFactory(),
		NewStructMarshalerFactory(sic),
	))

	uf := NewTypeCacheUnmarshalerFactory(NewCompositeFactory(
		FirstComposite,
		NewConvertPrimitiveUnmarshalerFactory(),
		NewPointerUnmarshalerFactory(),
		NewIndexUnmarshalerFactory(),
		NewMapUnmarshalerFactory(),
		NewBase64UnmarshalerFactory(),
		NewTimeUnmarshalerFactory(DefaultTimeUnmarshalLayouts()),
		NewOptionalUnmarshalerFactory(),
		NewStructUnmarshalerFactory(sic),
	))

	o := map[opt.Optional[int]][]*string{
		opt.Just(10):    {bt.PtrTo("ten"), bt.PtrTo("one zero")},
		opt.Just(20):    {bt.PtrTo("twenty"), bt.PtrTo("two zero")},
		opt.None[int](): {bt.PtrTo("none")},
	}
	fmt.Println(o)

	m := check.Must1(mf.Make(MarshalContext{Make: mf.Make}, reflect.TypeOf(o)))
	u := check.Must1(uf.Make(UnmarshalContext{Make: uf.Make}, reflect.TypeOf(o)))

	v := check.Must1(m.Marshal(MarshalContext{}, reflect.ValueOf(o)))
	fmt.Println(v)

	o2 := check.Must1(u.Unmarshal(UnmarshalContext{}, v))
	fmt.Println(o2)
}
