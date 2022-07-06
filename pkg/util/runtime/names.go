package runtime

import (
	"strings"

	"github.com/wrmsr/bane/pkg/util/slices"
)

//

type ParsedName struct {
	Pkg string `json:"pkg"`
	Obj string `json:"obj"`
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
	Builtin  []string `json:"builtin"`
	External []string `json:"external"`
	Internal []string `json:"internal"`
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
