package marshal

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

type MyInt int

func TestConvertUnmarshal(t *testing.T) {
	t0 := rfl.TypeOf[MyInt]()
	t1 := rfl.TypeOf[[]MyInt]()
	t2 := rfl.TypeOf[int]()
	t3 := rfl.TypeOf[[]int]()

	fmt.Println(t0)
	fmt.Println(t1)
	fmt.Println(t2)
	fmt.Println(t3)

	fmt.Println(t0.ConvertibleTo(t2))
	fmt.Println(t2.ConvertibleTo(t0))
	fmt.Println(t1.ConvertibleTo(t3))

	v := []MyInt{1}

	mv := check.Must1(Marshal(v))
	fmt.Println(mv)

	var v2 []MyInt
	check.Must(Unmarshal(mv, &v2))
	fmt.Println(v2)

	tu.AssertDeepEqual(t, v, v2)
}
