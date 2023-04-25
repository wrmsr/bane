package annotate

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

//

type SomeAnn struct {
	S string
}

type SomeStruct struct {
	S string
	I int
}

func (s SomeStruct) SomeMethod() {}

var _ = On[SomeStruct](SomeAnn{}).
	Field("S", SomeAnn{}).
	Method("SomeMethod", SomeAnn{})

//

func TestAnnotate(t *testing.T) {
}

//

func foo() {}

func TestFuncName(t *testing.T) {
	var fn any = foo
	fmt.Println(fn)
	rv := reflect.ValueOf(foo)
	fmt.Println(runtime.FuncForPC(rv.Pointer()).Name())
}
