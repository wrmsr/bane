package strings

import (
	"fmt"
	"testing"
)

func TestToCamel(t *testing.T) {
	fmt.Println(ToCamel("foo_bar_baz"))
}
