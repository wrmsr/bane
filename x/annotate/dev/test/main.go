//go:build !nodev

package main

import (
	"fmt"
	"reflect"

	"github.com/wrmsr/bane/x/annotate"
)

//

type Validate struct {
	Fn any
}

type NonZero struct{}

//

func CheckNonEmpty(s string) {
	if s == "" {
		panic("must not be empty")
	}
}

//

type SomeStruct struct {
	S string
	I int
}

func (s SomeStruct) SomeMethod() {}

var _ = annotate.On[SomeStruct]().
	Field("S", Validate{CheckNonEmpty}).
	Field("I", NonZero{})

//

func ValidateThing(v any) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer {
		panic(rv)
	}
	ta := annotate.Global().Get(rv.Elem().Type())
	if ta == nil {
		return
	}
	fmt.Printf("%#v\n", ta)
}

func main() {
	ss := SomeStruct{}
	ValidateThing(&ss)
}
