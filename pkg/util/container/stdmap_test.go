package container

import (
	"fmt"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

func TestMap(t *testing.T) {
	m := NewMutStdMap[int, string](nil)
	m.Put(10, "ten")
	m.Put(20, "twenty")
	m.Put(30, "thirty")
	tu.AssertEqual(t, m.Contains(10), true)
	tu.AssertEqual(t, m.Get(10), "ten")
	tu.AssertEqual(t, m.Contains(11), false)

	m.Put(11, "eleven")
	tu.AssertEqual(t, m.Contains(11), true)

	m.Delete(20)
	tu.AssertEqual(t, m.Contains(20), false)
}

//func TestMapReflect(t *testing.T) {
//	m := NewStdMap[int, string](nil)
//	ta := rfl.TypeArgs(reflect.TypeOf(m))
//	tu.AssertDeepEqual(t, ta, []reflect.Type{rfl.TypeOf[int](), rfl.TypeOf[string]()})
//}

func TestMapReflect2(t *testing.T) {
	ty := rfl.TypeOf[Map[int, string]]()
	ity := check.Ok1(ty.MethodByName("Iterate")).Type.Out(0)
	kvty := check.Ok1(ity.MethodByName("Next")).Type.Out(0)
	kty := check.Ok1(kvty.MethodByName("GetK")).Type.Out(0)
	vty := check.Ok1(kvty.MethodByName("GetV")).Type.Out(0)
	fmt.Println(kty, vty)
}
