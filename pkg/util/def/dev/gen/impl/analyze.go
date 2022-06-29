package impl

import (
	"go/types"
	"strings"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	"github.com/wrmsr/bane/pkg/util/def"
	rtu "github.com/wrmsr/bane/pkg/util/runtime"
)

func AnalyzeTypes(ps *def.PackageSpec) ctr.MutSet[rtu.ParsedName] {
	ret := ctr.NewMutSet[rtu.ParsedName](nil)

	doString := func(s string) {
		if !strings.Contains(s, ".") {
			return
		}
		ret.Add(rtu.ParseName(s))
	}

	doTypeName := func(tn *types.TypeName) {
		var sb strings.Builder
		var s string
		s = tn.Pkg().Path()
		if s != "" {
			sb.WriteString(s)
			sb.WriteByte('.')
		}
		sb.WriteString(tn.Name())
		doString(sb.String())
	}

	var doType func(ty types.Type)
	doType = func(ty types.Type) {
		switch o := ty.(type) {
		case *types.Basic:
			doString(o.String())
		case *types.Named:
			doTypeName(o.Obj())
			if ta := o.TypeArgs(); ta != nil {
				for i := 0; i < ta.Len(); i++ {
					doType(ta.At(i))
				}
			}
		default:
			panic(o)
		}
	}

	doField := func(fs *def.FieldSpec) {
		if fs.Type() != nil {
			if tr, ok := fs.Type().(TypeRef); ok {
				doType(tr.Type)
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
