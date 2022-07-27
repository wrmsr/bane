//go:build !nodev

package impl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wrmsr/bane/pkg/util/def"
	gg "github.com/wrmsr/bane/pkg/util/go/gen"
	tyu "github.com/wrmsr/bane/pkg/util/go/types"
	"github.com/wrmsr/bane/pkg/util/maps"
	opt "github.com/wrmsr/bane/pkg/util/optional"
	rtu "github.com/wrmsr/bane/pkg/util/runtime"
	"github.com/wrmsr/bane/pkg/util/slices"
	stru "github.com/wrmsr/bane/pkg/util/strings"
)

//

//

func CollectTypeNames(ps *def.PackageSpec) maps.Set[rtu.ParsedName] {
	ret := maps.MakeSet[rtu.ParsedName]()

	var doType func(ty tyu.Spec)
	doType = func(ty tyu.Spec) {
		ret.Add(ty.Parse())
		for _, a := range ty.Args {
			doType(a)
		}
	}

	doField := func(fs *def.FieldSpec) {
		if fs.Type() != nil {
			doType(tyu.SpecOf(fs.Type()))
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

	doEnum := func(es *def.EnumSpec) {
		doType(tyu.SpecOf(es.Ty()))
	}

	for _, es := range ps.Enums() {
		doEnum(es)
	}

	return ret
}

//

type typeImporter struct {
	ps   *def.PackageSpec
	mod  string
	imps map[string]string
}

func newTypeImporter(ps *def.PackageSpec, mod string, deps []string) *typeImporter {
	ts := CollectTypeNames(ps)

	pkgSet := maps.MakeSet[string]()
	pkgSet.Add(def.X_defPackageName())
	for _, d := range deps {
		pkgSet.Add(d)
	}
	for pn := range ts {
		if pn.Pkg != "" && pn.Pkg != ps.Name() {
			pkgSet.Add(pn.Pkg)
		}
	}

	pkgs := pkgSet.Slice()
	slices.Sort(pkgs)

	imps := make(map[string]string, len(pkgs))
	rimps := make(map[string]string, len(pkgs))
	addImp := func(pkg string) {
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

	depSet := maps.MakeSetOf(deps)
	depPkgs, nonDepPkgs := slices.Partition(depSet.Contains, pkgs)
	slices.Apply(addImp, depPkgs)
	slices.Apply(addImp, nonDepPkgs)

	return &typeImporter{
		ps:   ps,
		mod:  mod,
		imps: imps,
	}
}

func (ti *typeImporter) imports() []gg.Import {
	var imps []gg.Import
	for k, v := range ti.imps {
		_, n, _ := stru.LastCut(k, "/")
		if n != v {
			imps = append(imps, gg.NewImportAs(k, opt.Just[gg.Ident](gg.NewIdent(v))))
		} else {
			imps = append(imps, gg.NewImport(k))
		}
	}

	sysImps, restImps := slices.Partition(func(i gg.Import) bool {
		return !strings.Contains(i.Spec, "/")
	}, imps)
	depImps, prjImps := slices.Partition(func(i gg.Import) bool {
		return !strings.HasPrefix(i.Spec, ti.mod)
	}, restImps)

	impSecs := slices.Filter(func(s []gg.Import) bool { return len(s) > 0 }, [][]gg.Import{sysImps, depImps, prjImps})
	return slices.DeepFlatten[gg.Import](slices.Interleave(impSecs, []gg.Import{gg.NewImport("")}))
}

func (ti *typeImporter) importedType(ty any) gg.Type {
	var sb strings.Builder
	var rec func(ty tyu.Spec)
	rec = func(ty tyu.Spec) {
		pn := ty.Parse()
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

		if len(ty.Args) > 0 {
			sb.WriteString("[")
			for i, a := range ty.Args {
				if i > 0 {
					sb.WriteString(", ")
				}
				rec(a)
			}
			sb.WriteString("]")
		}
	}
	rec(tyu.SpecOf(ty))
	return gg.NewNameType(gg.NewIdent(sb.String()))
}
