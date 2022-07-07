package runtime

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/wrmsr/bane/pkg/util/slices"
	stru "github.com/wrmsr/bane/pkg/util/strings"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type ParsedName struct {
	Pkg string `json:"pkg,omitempty"`
	Obj string `json:"obj,omitempty"`
}

func (n ParsedName) String() string {
	if n.Pkg == "" {
		return n.Obj
	}
	return n.Pkg + "." + n.Obj
}

func ParseName(name string) ParsedName {
	if strings.IndexByte(name, '[') >= 0 {
		panic("generics not supported")
	}
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
	Args []ParsedGenericName
}

func (n ParsedGenericName) String() string {
	if len(n.Args) < 1 {
		return n.Base.String()
	}
	var sb strings.Builder
	sb.WriteString(n.Base.String())
	sb.WriteString("[")
	for i, a := range n.Args {
		if i > 0 {
			sb.WriteString(",")
		}
		sb.WriteString(a.String())
	}
	sb.WriteString("]")
	return sb.String()
}

func ParseGenericName(name string) (ParsedGenericName, error) {
	nl, err := stru.ParseNestedList(name, '[', ']', ',')
	if err != nil {
		return ParsedGenericName{}, err
	}

	var rec func([]any) (ParsedGenericName, error)
	rec = func(nl []any) (ParsedGenericName, error) {
		if len(nl) != 1 && len(nl) != 2 {
			return ParsedGenericName{}, fmt.Errorf("invalid length: %d", nl)
		}

		ret := ParsedGenericName{Base: ParseName(nl[0].(string))}

		if len(nl) > 1 {
			al := nl[1].([]any)

			for i := 0; i < len(al); {
				var j = i + 1
				if j < len(al) && bt.Is[[]any](al[j]) {
					j++
				}

				a, err := rec(al[i:j])
				if err != nil {
					return ret, err
				}

				ret.Args = append(ret.Args, a)
				i += j
			}
		}

		return ret, nil
	}

	return rec(nl)
}

//

func ParseTypeName(ty reflect.Type) (ParsedGenericName, error) {
	s := ty.Name()
	if ty.PkgPath() != "" {
		s = ty.PkgPath() + "." + s
	}
	return ParseGenericName(s)
}
