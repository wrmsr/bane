//go:build !nodev

package impl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wrmsr/bane/core/def"
	gg "github.com/wrmsr/bane/core/go/gen"
	tyu "github.com/wrmsr/bane/core/go/types"
	"github.com/wrmsr/bane/core/maps"
	rtu "github.com/wrmsr/bane/core/runtime"
	"github.com/wrmsr/bane/core/slices"
	stru "github.com/wrmsr/bane/core/strings"
)

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
	ps  *def.PackageSpec
	mod string

	imps     map[string]string
	usedImps map[string]struct{}
}

func newTypeImporter(ps *def.PackageSpec, mod string, deps []string) *typeImporter {
	ts := CollectTypeNames(ps)

	pkgSet := maps.MakeSet[string]()
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
		ps:  ps,
		mod: mod,

		imps:     imps,
		usedImps: make(map[string]struct{}),
	}
}

func (ti *typeImporter) importPkg(pkg string) (string, bool) {
	in, ok := ti.imps[pkg]
	if ok {
		ti.usedImps[pkg] = struct{}{}
	}
	return in, ok
}

func (ti *typeImporter) imports() []gg.Import {
	var imps []gg.Import
	for k, v := range ti.imps {
		if _, ok := ti.usedImps[k]; !ok {
			continue
		}
		_, n, _ := stru.LastCut(k, "/")
		if n != v {
			imps = append(imps, gg.Import{Spec: k, Alias: gg.NewIdent(v)})
		} else {
			imps = append(imps, gg.Import{Spec: k})
		}
	}

	sysImps, restImps := slices.Partition(func(i gg.Import) bool {
		return !strings.Contains(i.Spec, "/")
	}, imps)
	depImps, prjImps := slices.Partition(func(i gg.Import) bool {
		return !strings.HasPrefix(i.Spec, ti.mod)
	}, restImps)

	impSecs := slices.Filter(func(s []gg.Import) bool { return len(s) > 0 }, [][]gg.Import{sysImps, depImps, prjImps})
	return slices.DeepFlatten[gg.Import](slices.Interleave(impSecs, []gg.Import{{Spec: ""}}))
}

func (ti *typeImporter) importedType(ty any) gg.Type {
	var sb strings.Builder
	var rec func(ty tyu.Spec)
	rec = func(ty tyu.Spec) {
		pn := ty.Parse()
		if pn.Pkg != "" && pn.Pkg != ti.ps.Name() {
			in, ok := ti.importPkg(pn.Pkg)
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
	return gg.TypeOf(sb.String())
}
