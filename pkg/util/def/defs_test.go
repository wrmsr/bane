package def

import "testing"

//

var _ = Struct("Foo",
	Field("bar", Type[int]()),
	Field("baz", Type[int]()),
)

//

func TestMockup(t *testing.T) {

}
