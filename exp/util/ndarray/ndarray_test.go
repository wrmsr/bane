package ndarray

import (
	"fmt"
	"testing"

	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TestNdArray(t *testing.T) {
	a := Of[int](
		[]int{3, 3, 3},
		nil,
		0,
		bt.RangeTo(27).Slice(),
	)
	fmt.Println(a)
	fmt.Println("====")

	//idx := []Dim{1, 2, 0}
	//fmt.Println(a.Get(idx...))
	//fmt.Println("====")

	//*a.At(idx...) = 420
	//fmt.Println(a)
	//fmt.Println(a.Get(idx...))
	//fmt.Println("====")

	for _, sl := range [][]any{
		{1, nil, nil},
		{nil, 1, nil},
		{nil, nil, 1},
	} {
		fmt.Println(sl)
		fmt.Println("====")

		sa := a.Slice(sl...)
		fmt.Println(sa)
		fmt.Println("====")

		fmt.Println(sa.Squeeze())
		fmt.Println("====")
	}
}
