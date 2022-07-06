package runtime

import (
	"strings"

	"github.com/wrmsr/bane/pkg/util/slices"
)

//

type ParsedName struct {
	Pkg string `json:"pkg,omitempty"`
	Obj string `json:"obj,omitempty"`
}

func ParseName(name string) ParsedName {
	sl := strings.LastIndexByte(name, '/')
	if sl < 0 {
		return ParsedName{Obj: name}
	}
	dl := strings.IndexByte(name[sl:], '.') + sl
	if dl < 0 {
		return ParsedName{Obj: name}
	}
	return ParsedName{
		Pkg: name[:dl],
		Obj: name[dl+1:],
	}
}

//

type SortedNames struct {
	Builtin  []string `json:"builtin,omitempty"`
	External []string `json:"external,omitempty"`
	Internal []string `json:"internal,omitempty"`
}

func SortNames(pfx string, ns []string) (ret SortedNames) {
	for _, n := range slices.Sorted(ns) {
		if !strings.ContainsRune(n, '/') || !strings.ContainsRune(n, '.') {
			ret.Builtin = append(ret.Builtin, n)
			continue
		}
		pn := ParseName(n)
		if !strings.HasPrefix(pn.Pkg, pfx) {
			ret.External = append(ret.External, n)
		} else {
			ret.Internal = append(ret.Internal, n)
		}
	}

	return
}

//

type ParsedGenericName struct {
	Base ParsedName
	Args []ParsedName
}

func ParseGenericName(name string) ParsedGenericName {
	// Optional[github.com/wrmsr/bane/pkg/util/container.Map[int,github.com/wrmsr/bane/pkg/util/optional.Optional[string]]]
	panic("TODO")
}
