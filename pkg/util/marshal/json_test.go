package marshal

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	ju "github.com/wrmsr/bane/pkg/util/json"
)

func TestJson(t *testing.T) {
	v := Array{v: []Value{
		Int{v: 420},
		Float{v: 4.2},
		String{v: "four twenty"},
	}}
	fmt.Println(v)

	j := check.Must1(ju.MarshalString(v))
	fmt.Println(j)

	v2 := check.Must1(JsonUnmarshal([]byte(j)))
	fmt.Println(v2)

	tu.AssertDeepEqual(t, v, v2)
}