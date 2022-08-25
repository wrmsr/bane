package tests

import (
	"fmt"
	"testing"

	nd "github.com/wrmsr/bane/exp/util/ndarray"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

func TestNdArray(t *testing.T) {
	a := nd.Of[int](
		nd.ShapeOf(3, 3, 3),
		nd.Strides{},
		0,
		bt.RangeTo(27).Slice(),
	)
	fmt.Println(a)
	fmt.Println("====")

	n := 0
	for i := nd.Dim(0); i < 3; i++ {
		for j := nd.Dim(0); j < 3; j++ {
			for k := nd.Dim(0); k < 3; k++ {
				tu.AssertEqual(t, a.Get(i, j, k), n)
				n++
			}
		}
	}

	//idx := []Dim{1, 2, 0}
	//fmt.Println(a.Get(idx...))
	//fmt.Println("====")

	//*a.At(idx...) = 420
	//fmt.Println(a)
	//fmt.Println(a.Get(idx...))
	//fmt.Println("====")

	for _, sl := range [][]any{
		{1},
		{nil, 1},
		{nil, nil, 1},
		{[]any{1, nil}},
		{[]any{1, nil}, 1},
		{[]any{1, nil}, nil, 1},
		{[]any{nil, nil, 2}},
		{1, nil, []any{nil, nil, 2}},
	} {
		fmt.Println(sl)
		fmt.Println("====")

		sa := a.Slice(sl...)
		fmt.Println(sa.View())
		fmt.Println(sa)
		fmt.Println("====")

		sqa := sa.Squeeze()
		fmt.Println(sqa.View())
		fmt.Println(sqa)
		fmt.Println("====")
	}
}

func TestScalar(t *testing.T) {
	sc := nd.Of[int](nd.ShapeOf(1), nd.Strides{}, 0, nil)
	*sc.At(0) = 420
	fmt.Println(sc)
}

func BenchmarkSliceView(b *testing.B) {
	v := nd.ViewOf(
		nd.ShapeOf(3, 3, 3),
		nd.Strides{},
		0,
	)

	o := []any{1, nil, []any{1, nil, 2}}
	//o := []any{1}

	fmt.Println(b.N)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		sv := v.Slice(o...)
		if sv.Len() < 0 {
			panic("oops")
		}
	}
}
