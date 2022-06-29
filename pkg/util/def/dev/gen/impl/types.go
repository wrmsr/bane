//go:build !nodev

package impl

import (
	"go/types"
	"strings"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	"github.com/wrmsr/bane/pkg/util/def"
	rtu "github.com/wrmsr/bane/pkg/util/runtime"
)

//

type TypeRef struct {
	Name string
	Args []TypeRef

	Origin types.Type
}

func (tr TypeRef) String() string {
	return tr.Origin.String()
}

func (tr TypeRef) Parse() rtu.ParsedName {
	return rtu.ParseName(tr.Name)
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

func typeRef(o any) (tr TypeRef) {
	switch o := o.(type) {

	case *types.Basic:
		tr.Origin = o
		tr.Name = o.String()

	case *types.Named:
		tr.Origin = o
		tr.Name = typeNameString(o.Obj())
		if ta := o.TypeArgs(); ta != nil {
			for i := 0; i < ta.Len(); i++ {
				tr.Args = append(tr.Args, typeRef(ta.At(i)))
			}
		}

	default:
		panic(o)
	}
	return
}

//

func CollectTypeNames(ps *def.PackageSpec) ctr.MutSet[rtu.ParsedName] {
	ret := ctr.NewMutSet[rtu.ParsedName](nil)

	var doType func(ty TypeRef)
	doType = func(tr TypeRef) {
		ret.Add(tr.Parse())
		for _, a := range tr.Args {
			doType(a)
		}
	}

	doField := func(fs *def.FieldSpec) {
		if fs.Type() != nil {
			if tr, ok := fs.Type().(TypeRef); ok {
				doType(tr)
			}
		}
	}

	doStruct := func(ss *def.StructSpec) {
		for _, fs := range ss.Fields() {
			doField(fs)
		}
	}

	for _, ss := range ps.Structs() {
		doStruct(ss)
	}

	return ret
}
