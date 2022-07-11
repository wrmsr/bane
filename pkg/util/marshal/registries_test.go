package marshal

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/wrmsr/bane/pkg/util/check"
	tu "github.com/wrmsr/bane/pkg/util/dev/testing"
	rfl "github.com/wrmsr/bane/pkg/util/reflect"
)

type RegThing struct {
	I int
	S string
}

func TestRegistry(t *testing.T) {
	reg := NewRegistry(nil)

	reg.Register(rfl.TypeOf[RegThing](),
		SetMarshaler{NewFuncMarshaler(func(ctx MarshalContext, rv reflect.Value) (Value, error) {
			return MakeString("reg_thing"), nil
		})})

	o := reg.m[rfl.TypeOf[RegThing]()].m[rfl.TypeOf[SetMarshaler]()][0]
	tu.AssertEqual(t, o != nil, true)

	rt := RegThing{I: 420, S: "four twenty"}
	mv := check.Must1(o.(SetMarshaler).M.Marshal(MarshalContext{}, reflect.ValueOf(rt)))
	fmt.Println(mv)
}
