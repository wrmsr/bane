package ndarray

import (
	"fmt"
	"testing"
)

func TestInline(t *testing.T) {
	v := ViewOf(ShapeOf(3, 7), Strides{}, 0)
	fmt.Println(foo(v))
	fmt.Println(_def_inl_foo(v))
}
