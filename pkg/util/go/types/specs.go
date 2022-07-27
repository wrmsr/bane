package types

import (
	"go/types"
	"strings"

	rtu "github.com/wrmsr/bane/pkg/util/runtime"
)

type Spec struct {
	Name string
	Args []Spec

	Origin types.Type
}

func (s Spec) String() string {
	return s.Origin.String()
}

func (s Spec) Parse() rtu.ParsedName {
	return rtu.ParseName(s.Name)
}

func typeNameString(tn *types.TypeName) string {
	var sb strings.Builder
	var s string
	s = tn.Pkg().Path()
	if s != "" {
		sb.WriteString(s)
		sb.WriteByte('.')
	}
	sb.WriteString(tn.Name())
	return sb.String()
}

func SpecOf(o any) (tr Spec) {
	switch o := o.(type) {

	case *types.Basic:
		tr.Origin = o
		tr.Name = o.String()

	case *types.Named:
		tr.Origin = o
		tr.Name = typeNameString(o.Obj())
		if ta := o.TypeArgs(); ta != nil {
			for i := 0; i < ta.Len(); i++ {
				tr.Args = append(tr.Args, SpecOf(ta.At(i)))
			}
		}

	default:
		panic(o)
	}
	return
}
