package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"unsafe"

	"github.com/wrmsr/bane/pkg/util/check"
	rtu "github.com/wrmsr/bane/pkg/util/runtime"
)

//

var marker = 0

//

func debugWait() {
	//log.Println(os.Getpid())
	//bufio.NewScanner(os.Stdin).Scan()
}

func main() {
	debugWait()

	pn := rtu.GetCaller(0)
	pu := reflect.ValueOf(&marker).Pointer()

	name := os.Args[0]
	os := ""

	//name = "/Users/spinlock/src/wrmsr/bane/bin/linuxlogfoo"
	//os = "linux"

	var f SymFile
	var clo Closer

	f, clo = check.Must2(OpenSymFile(name, os, true))
	defer func() {
		if clo != nil {
			check.Must(clo())
		}
	}()

	var ofs uintptr
	var lp uintptr
	mn := fmt.Sprintf("%s.%s", pn.Pkg, "marker")

	check.Must1(f.ForEachSym(func(s FileSym) bool {
		sn := s.Name.String()
		if sn == mn {
			ofs = pu - uintptr(s.Addr)
		}
		if sn == "log.std" {
			lp = uintptr(s.Addr)
		}
		return ofs == 0 || lp == 0
	}))

	var std **log.Logger = (**log.Logger)(unsafe.Pointer(lp + ofs))
	(*std).Println("hi")

	debugWait()
}
