package annotate

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

//

type NotEmpty struct{}

type SomeStruct struct {
	S string
	I int
}

func (s SomeStruct) SomeMethod() {}

var _ = On[SomeStruct]().
	Field("S", NotEmpty{})

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

	for _, ma := range ta.mc.m {
		for _, a := range ma.c.s {
			switch a := a.(type) {
			case NotEmpty:
				switch ma.m.K {
				case FieldMember:
					fv := rv.Elem().Field(ma.ReflectField().Index[0])
					switch fv.Type() {
					case reflect.TypeOf(""):
						if fv.Interface() == "" {
							panic("no empty")
						}
					default:
						panic(fv)
					}
				default:
					panic(ma)
				}
			default:
				panic(a)
			}
		}
	}
}

func TestFuncName(t *testing.T) {
	var fn any = foo
	fmt.Println(fn)

	rv := reflect.ValueOf(foo)
	fmt.Println(runtime.FuncForPC(rv.Pointer()).Name())

	for _, ss := range []SomeStruct{
		{S: "barf"},
		{},
	} {
		fmt.Printf("%#v\n", ss)
		doThing(&ss)
	}
}
