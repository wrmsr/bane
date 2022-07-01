package marshal

import (
	"reflect"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	stu "github.com/wrmsr/bane/pkg/util/structs"
)

func TestObjects(t *testing.T) {
	asi := stu.NewStructInfoCache().Info((*A)(nil))
	am := NewObjectMarshaler(
		NewObjectMarshalerField("X", NewObjectFieldGetter(asi.Field("X")), PrimitiveMarshaler{}),
		NewObjectMarshalerField("Y", NewObjectFieldGetter(asi.Field("Y")), PrimitiveMarshaler{}),
	)

	csi := stu.NewStructInfoCache().Info((*C)(nil))
	cm := NewObjectMarshaler(
		NewObjectMarshalerField("A", NewObjectFieldGetter(csi.Field("A")), am),
		NewObjectMarshalerField("Z", NewObjectFieldGetter(csi.Field("Z")), PrimitiveMarshaler{}),
	)

	c := testC
	mv := check.Must1(cm.Marshal(MarshalContext{}, reflect.ValueOf(c)))
	tu.AssertEqual(t, mv.String(), `{"A": {"X": 100, "Y": "two hundred"}, "Z": 420}`)
}
