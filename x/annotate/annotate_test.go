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

func doThing(v any) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer {
		panic(rv)
	}
	ta := Global().Get(rv.Elem().Type())
	if ta == nil {
		return
	}
	fmt.Println(ta)
}

func TestFuncName(t *testing.T) {
	var fn any = foo
	fmt.Println(fn)

	rv := reflect.ValueOf(foo)
	fmt.Println(runtime.FuncForPC(rv.Pointer()).Name())

	ss := SomeStruct{}
	doThing(&ss)
}
