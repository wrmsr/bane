package annotate

import (
	"fmt"
	"reflect"
	"testing"
)

//

type Annotator interface {
	Field(n string, anns ...any) Annotator
	Method(n string, anns ...any) Annotator
}

type annotator struct {
	ty reflect.Type

	s []any

	fm map[string][]any
	mm map[string][]any
}

var _ Annotator = &annotator{}

func Annotate[T any](anns ...any) Annotator {
	var z T
	ty := reflect.TypeOf(z)
	if ty.Kind() != reflect.Struct {
		panic(fmt.Errorf("must be struct: %v", ty))
	}
	return &annotator{
		ty: ty,

		s: anns,
	}
}

func (a *annotator) Field(n string, anns ...any) Annotator {
	if a.fm == nil {
		a.fm = make(map[string][]any)
	}
	a.fm[n] = append(a.fm[n], anns...)
	return a
}

func (a *annotator) Method(n string, anns ...any) Annotator {
	if a.mm == nil {
		a.mm = make(map[string][]any)
	}
	a.mm[n] = append(a.mm[n], anns...)
	return a
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

var _ = Annotate[SomeStruct](SomeAnn{}).
	Field("S", SomeAnn{}).
	Method("SomeMethod", SomeAnn{})

//

func TestAnnotate(t *testing.T) {
}
