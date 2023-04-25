package def

import "testing"

//

//var _ = Struct[Foo](
//	Field("bar", Type[int]()),
//	Field("baz", Type[int]()),
//)

//

type MyEnum int8

const (
	InvalidMyEnum MyEnum = iota
	FooMyEnum
	BarMyEnum
)

//

func TestMockup(t *testing.T) {

}
