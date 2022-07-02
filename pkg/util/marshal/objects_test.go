package marshal

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	stu "github.com/wrmsr/bane/pkg/util/structs"
)

func TestObjects(t *testing.T) {
	asi := stu.NewStructInfoCache().Info((*A)(nil))
	am := NewObjectMarshaler(
		NewObjectMarshalerField("X", NewStructFieldGetter(asi.Field("X")), PrimitiveMarshaler{}),
		NewObjectMarshalerField("Y", NewStructFieldGetter(asi.Field("Y")), PrimitiveMarshaler{}),
	)

	csi := stu.NewStructInfoCache().Info((*C)(nil))
	cm := NewObjectMarshaler(
		NewObjectMarshalerField("A", NewStructFieldGetter(csi.Field("A")), am),
		NewObjectMarshalerField("Z", NewStructFieldGetter(csi.Field("Z")), PrimitiveMarshaler{}),
	)

	c := testC
	mv := check.Must1(cm.Marshal(MarshalContext{}, reflect.ValueOf(c)))
	tu.AssertEqual(t, mv.String(), `{"A": {"X": 100, "Y": "two hundred"}, "Z": 420}`)

	au := NewObjectUnmarshaler(
		NewStructFactory(asi),
		NewObjectUnmarshalerField("X", NewStructFieldSetter(asi.Field("X")), NewConvertUnmarshaler(rfl.TypeOf[int](), PrimitiveUnmarshaler{})),
		NewObjectUnmarshalerField("Y", NewStructFieldSetter(asi.Field("Y")), PrimitiveUnmarshaler{}),
	)

	cu := NewObjectUnmarshaler(
		NewStructFactory(csi),
		NewObjectUnmarshalerField("A", NewStructFieldSetter(csi.Field("A")), au),
		NewObjectUnmarshalerField("Z", NewStructFieldSetter(csi.Field("Z")), NewConvertUnmarshaler(rfl.TypeOf[int32](), PrimitiveUnmarshaler{})),
	)

	c2 := check.Must1(cu.Unmarshal(UnmarshalContext{}, mv)).Interface().(C)
	fmt.Println(c2)

	c3 := c
	c3.B.Z = 0
	tu.AssertEqual(t, c2, c3)
}
