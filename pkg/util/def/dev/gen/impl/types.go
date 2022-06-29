//go:build !nodev

package impl

import (
	"fmt"
	"go/types"
	"strconv"
	"strings"

	ctr "github.com/wrmsr/bane/pkg/util/container"
	"github.com/wrmsr/bane/pkg/util/def"
	gg "github.com/wrmsr/bane/pkg/util/gogen"
	its "github.com/wrmsr/bane/pkg/util/iterators"
	rtu "github.com/wrmsr/bane/pkg/util/runtime"
	"github.com/wrmsr/bane/pkg/util/slices"
	stru "github.com/wrmsr/bane/pkg/util/strings"
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

//

type typeImporter struct {
	ps   *def.PackageSpec
	imps map[string]string
}

func newTypeImporter(ps *def.PackageSpec) *typeImporter {
	ts := CollectTypeNames(ps)

	pkgSet := ctr.NewMutSet[string](nil)
	pkgSet.Add(def.X_defPackageName())
	ts.ForEach(func(pn rtu.ParsedName) bool {
		if pn.Pkg != "" && pn.Pkg != ps.Name() {
			pkgSet.Add(pn.Pkg)
		}
		return true
	})

	pkgs := its.Seq[string](pkgSet)
	slices.Sort(pkgs)

	imps := make(map[string]string, len(pkgs))
	rimps := make(map[string]string, len(pkgs))
	for _, pkg := range pkgs {
		_, pn, _ := stru.LastCut(pkg, "/")
		for i := 1; ; i++ {
			s := pn
			if i > 1 {
				s += strconv.Itoa(i)
			}
			if rimps[s] == "" {
				rimps[s] = pkg
				imps[pkg] = s
				break
			}
		}
	}

	return &typeImporter{
		ps:   ps,
		imps: imps,
	}
}

func (ti *typeImporter) importedType(ty any) gg.Type {
	var sb strings.Builder
	var rec func(tr TypeRef)
	rec = func(tr TypeRef) {
		pn := tr.Parse()
		if pn.Pkg != "" && pn.Pkg != ti.ps.Name() {
			in, ok := ti.imps[pn.Pkg]
			if !ok {
				panic(fmt.Errorf("import not found: %s", pn.Pkg))
			}
			if in != "" {
				sb.WriteString(in)
				sb.WriteString(".")
			}
			sb.WriteString(pn.Obj)
		} else {
			sb.WriteString(pn.Obj)
		}

		if len(tr.Args) > 0 {
			sb.WriteString("[")
			for i, a := range tr.Args {
				if i > 0 {
					sb.WriteString(", ")
				}
				rec(a)
			}
			sb.WriteString("]")
		}
	}
	rec(ty.(TypeRef))
	return gg.NewNameType(gg.NewIdent(sb.String()))
}
