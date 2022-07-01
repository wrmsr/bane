package marshal

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

func TestSlice(t *testing.T) {
	s := []int{1, 10, 100}

	mv := check.Must1(NewIndexMarshaler(PrimitiveMarshaler{}).Marshal(MarshalContext{}, reflect.ValueOf(s)))
	fmt.Println(mv)

	s2 := check.Must1(NewSliceUnmarshaler(rfl.TypeOf[[]int](), NewConvertUnmarshaler(rfl.TypeOf[int](), PrimitiveUnmarshaler{})).Unmarshal(UnmarshalContext{}, mv))
	fmt.Println(s2)

	tu.AssertDeepEqual(t, s, s2)
}
