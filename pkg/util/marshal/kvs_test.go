package marshal

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	ctr "github.com/wrmsr/bane/pkg/util/container"
)

func TestKvIterable(t *testing.T) {
	m := ctr.NewOrderedMapBuilder[string, string]().
		Put("100", "one hundred").
		Put("200", "two hundred").
		Build()

	mv := check.Must1(KvIterableMarshaler{
		k: PrimitiveMarshaler{},
		v: PrimitiveMarshaler{},
	}.Marshal(
		MarshalContext{},
		reflect.ValueOf(m),
	))

	fmt.Println(mv)
}
