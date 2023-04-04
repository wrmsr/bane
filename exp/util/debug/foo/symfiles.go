package main

import (
	"debug/elf"
	"debug/macho"
	"runtime"
)

//

type FileSym struct {
	Name string
	Addr uint64
}

type SymFile interface {
	ForEachSym(fn func(s FileSym) bool) (bool, error)
}

//

type machoSymFile struct {
	f *macho.File
}

var _ SymFile = machoSymFile{}

func (f machoSymFile) ForEachSym(fn func(s FileSym) bool) (bool, error) {
	for _, s := range f.f.Symtab.Syms {
		if !fn(FileSym{
			Name: s.Name,
			Addr: s.Value,
		}) {
			return false, nil
		}
	}
	return true, nil
}

//

type elfSymFile struct {
	f *elf.File
}

var _ SymFile = elfSymFile{}

func (f elfSymFile) ForEachSym(fn func(s FileSym) bool) (bool, error) {
	syms, err := f.f.Symbols()
	if err != nil {
		return false, err
	}

	for _, s := range syms {
		if !fn(FileSym{
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

func OpenSymFile(name string, os string, singleShot bool) (SymFile, Closer, error) {
	if os == "" {
		os = runtime.GOOS
	}
	switch os {
	case "darwin":
		if singleShot {
			return machoSymFile2{name: name}, nil, nil
		}
		f, err := macho.Open(name)
		if err != nil {
			return nil, nil, err
		}
		return machoSymFile{f: f}, f.Close, nil
	case "linux":
		f, err := elf.Open(name)
		if err != nil {
			return nil, nil, err
		}
		return elfSymFile{f: f}, f.Close, nil
	}
	return nil, nil, UnsupportedPlatformError{}
}
