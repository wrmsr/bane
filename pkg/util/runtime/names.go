package runtime

import "strings"

type ParsedName struct {
	Pkg string
	Obj string
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
