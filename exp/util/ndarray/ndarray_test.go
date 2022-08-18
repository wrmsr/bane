package ndarray

import (
	"fmt"
	"testing"
)

func TestNdArray(t *testing.T) {
	a := New[int]([]int{3, 3, 3})
	fmt.Println(a)
}
