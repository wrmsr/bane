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

func SpecOf(o any) (s Spec) {
	switch o := o.(type) {

	case Spec:
		s = o

	case *types.Basic:
		s.Origin = o
		s.Name = o.String()

	case *types.Named:
		s.Origin = o
		s.Name = typeNameString(o.Obj())
		if ta := o.TypeArgs(); ta != nil {
			for i := 0; i < ta.Len(); i++ {
				s.Args = append(s.Args, SpecOf(ta.At(i)))
			}
		}

	default:
		panic(o)
	}
	return
}
