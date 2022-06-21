package runtime

import "strings"

func ParseName(name string) ParsedName {
	sl := strings.LastIndexByte(name, '/')
	if sl < 0 {
		sl = 0
	}
	dl := strings.IndexByte(name[sl:], '.') + sl
	return ParsedName{
		Pkg: name[:dl],
		Obj: name[dl+1:],
	}
}
