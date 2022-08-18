package ndarray

import (
	"fmt"
	"testing"
)

func TestNdArray(t *testing.T) {
	idx := []Dim{1, 2, 0}

	a := New[int]([]int{3, 3, 3})
	fmt.Println(a)
	fmt.Println(a.Get(idx...))

	*a.At(idx...) = 420
	fmt.Println(a)
	fmt.Println(a.Get(idx...))
}
