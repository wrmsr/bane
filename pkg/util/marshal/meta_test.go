package marshal

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
)

func TestSwitchingMarshaler(t *testing.T) {
	m := NewSwitchingMarshaler(
		PredicatedMarshaler(TrueMarshalPredicate, PrimitiveMarshaler{}),
	)

	mv := check.Must1(m.Marshal(MarshalContext{}, reflect.ValueOf(420)))
	fmt.Println(mv)

	u := NewSwitchingUnmarshaler(
		PredicatedUnmarshaler(TrueUnmarshalPredicate, PrimitiveUnmarshaler{}),
	)

	rv := check.Must1(u.Unmarshal(UnmarshalContext{}, mv))
	fmt.Println(rv)
}
