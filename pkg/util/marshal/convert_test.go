package marshal

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
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

	//fmt.Println(check.Must1(ju.MarshalString(MyInt(420))))

	_ = check.Must1(Marshal([]MyInt{1}))
}
