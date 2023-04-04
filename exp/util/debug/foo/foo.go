package main

import (
	"bufio"
	"debug/elf"
	"debug/macho"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime"
	"unsafe"

	"github.com/wrmsr/bane/pkg/util/check"
	rtu "github.com/wrmsr/bane/pkg/util/runtime"
)

//

type ExeSym struct {
	Name string
	Addr uint64
}

type ExeFile interface {
	ForEachSym(fn func(s ExeSym) bool) (bool, error)
}

//

type machoExeFile struct {
	f *macho.File
}

var _ ExeFile = machoExeFile{}

func (f machoExeFile) ForEachSym(fn func(s ExeSym) bool) (bool, error) {
	for _, s := range f.f.Symtab.Syms {
		if !fn(ExeSym{
			Name: s.Name,
			Addr: s.Value,
		}) {
			return false, nil
		}
	}
	return true, nil
}

//

type elfExeFile struct {
	f *elf.File
}

var _ ExeFile = elfExeFile{}

func (f elfExeFile) ForEachSym(fn func(s ExeSym) bool) (bool, error) {
	syms, err := f.f.Symbols()
	if err != nil {
		return false, err
	}

	for _, s := range syms {
		if !fn(ExeSym{
			Name: s.Name,
			Addr: s.Value,
		}) {
			return false, nil
		}
	}
	return true, nil
}

//

type UnsupportedPlatformError struct{}

var _ error = UnsupportedPlatformError{}

func (e UnsupportedPlatformError) Error() string {
	return "unsupported platform"
}

type Closer func() error

func OpenExeFile(name string) (ExeFile, Closer, error) {
	switch runtime.GOOS {
	case "darwin":
		f, err := macho.Open(name)
		if err != nil {
			return nil, nil, err
		}
		return machoExeFile{f: f}, f.Close, nil
	case "linux":
		f, err := elf.Open(name)
		if err != nil {
			return nil, nil, err
		}
		return elfExeFile{f: f}, f.Close, nil
	}
	return nil, nil, UnsupportedPlatformError{}
}

//

var marker = 0

//

func main() {
	pn := rtu.GetCaller(0)
	pu := reflect.ValueOf(&marker).Pointer()

	f, clo := check.Must2(OpenExeFile(os.Args[0]))
	defer func() {
		clo()
	}()

	var ofs uintptr
	var lp uintptr
	mn := fmt.Sprintf("%s.%s", pn.Pkg, "marker")

	check.Must1(f.ForEachSym(func(s ExeSym) bool {
		if s.Name == mn {
			ofs = pu - uintptr(s.Addr)
		}
		if s.Name == "log.std" {
			lp = uintptr(s.Addr)
		}
		return ofs == 0 || lp == 0
	}))

	var std **log.Logger = (**log.Logger)(unsafe.Pointer(lp + ofs))
	(*std).Println("hi")

	log.Println(os.Getpid())
	bufio.NewScanner(os.Stdin).Scan()
}
