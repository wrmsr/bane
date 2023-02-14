package annotate

import (
	"fmt"
	"reflect"
	"testing"
)

//

type Annotator struct {}


//

func Annotate(obj any, anns ...any) any {
	fmt.Println(obj)
	rv := reflect.ValueOf(obj)
	fmt.Println(rv)
	fmt.Println(rv.Type())
	return nil
}

//

type SomeAnn struct {
	S string
}

type SomeStruct struct {
	S string
	I int
}

func (s SomeStruct) SomeMethod() {}

var _ = Annotate(SomeStruct{}, SomeAnn{})
var _ = Annotate(SomeStruct{}.S, SomeAnn{})
var _ = Annotate(SomeStruct{}.SomeMethod, SomeAnn{})

//

func TestAnnotate(t *testing.T) {
}
