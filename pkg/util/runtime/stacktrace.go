package runtime

import (
	"runtime"
)

type StackFrame struct {
	Pc   uintptr
	Name string

	File string
	Line int
}

func GetStackFrame(ofs int) StackFrame {
	var pcs [1]uintptr
	_ = runtime.Callers(
		ofs+2,
		pcs[:],
	)
	frs := runtime.CallersFrames(pcs[:])
	fr, _ := frs.Next()
	return StackFrame{
		Pc:   fr.PC,
		File: fr.File,
		Line: fr.Line,
		Name: fr.Function,
	}
}

func GetStackTrace(n, ofs int) []StackFrame {
	pcs := make([]uintptr, n)
	nf := runtime.Callers(
		ofs+2,
		pcs,
	)
	pcs = pcs[:nf]
	frs := runtime.CallersFrames(pcs)
	sfs := make([]StackFrame, 0, nf)
	for {
		fr, more := frs.Next()
		sfs = append(sfs, StackFrame{
			Pc:   fr.PC,
			File: fr.File,
			Line: fr.Line,
			Name: fr.Function,
		})
		if !more {
			break
		}
	}
	return sfs
}

func GetCaller(ofs int) ParsedName {
	sf := GetStackFrame(ofs + 1)
	return ParseName(sf.Name)
}
