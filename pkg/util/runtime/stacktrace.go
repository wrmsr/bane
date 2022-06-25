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
