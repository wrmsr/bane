package marshal

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	stu "github.com/wrmsr/bane/pkg/util/structs"
)

type PolyI interface{ isPolyI() }

type PolyA struct{ I int }
type PolyB struct{ S string }
type PolyC struct{ L, R PolyI }

func (p PolyA) isPolyI() {}
func (p PolyB) isPolyI() {}
func (p PolyC) isPolyI() {}

func TestPolymorphism(t *testing.T) {
	i := PolyC{
		L: PolyA{
			I: 420,
		},
		R: PolyC{
			L: PolyA{
				I: 421,
			},
			R: PolyB{
				S: "four twenty three",
			},
		},
	}

	p := NewPolymorphism(rfl.TypeOf[PolyI](),
		PolymorphismEntry{Tag: "a", Impl: rfl.TypeOf[PolyA]()},
		PolymorphismEntry{Tag: "b", Impl: rfl.TypeOf[PolyB]()},
		PolymorphismEntry{Tag: "c", Impl: rfl.TypeOf[PolyC]()},
	)

	sic := stu.NewStructInfoCache()

	mf := NewTypeCacheMarshalerFactory(NewCompositeMarshalerFactory(
		FirstComposite,
		NewPolymorphicMarshalerFactory(p),
		NewStructMarshalerFactory(sic),
		NewPrimitiveMarshalerFactory(),
	))

	uf := NewTypeCacheUnmarshalerFactory(NewCompositeUnmarshalerFactory(
		FirstComposite,
		NewPolymorphicUnmarshalerFactory(p),
		NewStructUnmarshalerFactory(sic),
		NewConvertPrimitiveUnmarshalerFactory(),
	))

	m := check.Must1(mf.Make(MarshalContext{Make: mf.Make}, reflect.TypeOf(i)))
	u := check.Must1(uf.Make(UnmarshalContext{Make: uf.Make}, reflect.TypeOf(i)))

	v := check.Must1(m.Marshal(MarshalContext{}, reflect.ValueOf(i)))
	fmt.Println(v)

	o2 := check.Must1(u.Unmarshal(UnmarshalContext{}, v))
	fmt.Println(o2)
}
