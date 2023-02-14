package annotate

import (
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
