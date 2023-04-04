package main

import (
	"bufio"
	"bytes"
	"debug/elf"
	"debug/macho"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
	"unsafe"

	"github.com/wrmsr/bane/pkg/util/check"
	iou "github.com/wrmsr/bane/pkg/util/io"
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

type machoExeFile2 struct {
	name string
}

var _ ExeFile = machoExeFile2{}

const (
	machoFileHeaderSize32 = 7 * 4
	machoFileHeaderSize64 = 8 * 4
)

func cstring(b []byte) string {
	i := bytes.IndexByte(b, 0)
	if i == -1 {
		i = len(b)
	}
	return string(b[0:i])
}

func (f2 machoExeFile2) ForEachSym(fn func(s ExeSym) bool) (bool, error) {
	r, err := os.Open(f2.name)
	if err != nil {
		return false, err
	}
	sr := io.NewSectionReader(r, 0, 1<<63-1)

	var f macho.File

	// Read and decode Mach magic to determine byte order, size.
	// Magic32 and Magic64 differ only in the bottom bit.
	var ident [4]byte
	if _, err := r.ReadAt(ident[0:], 0); err != nil {
		return false, err
	}
	be := binary.BigEndian.Uint32(ident[0:])
	le := binary.LittleEndian.Uint32(ident[0:])
	switch macho.Magic32 &^ 1 {
	case be &^ 1:
		f.ByteOrder = binary.BigEndian
		f.Magic = be
	case le &^ 1:
		f.ByteOrder = binary.LittleEndian
		f.Magic = le
	default:
		return false, errors.New("invalid magic number")
	}

	// Read entire file header.
	if err := binary.Read(sr, f.ByteOrder, &f.FileHeader); err != nil {
		return false, err
	}

	// Then load commands.
	offset := int64(machoFileHeaderSize32)
	if f.Magic == macho.Magic64 {
		offset = machoFileHeaderSize64
	}
	dat, err := iou.SafeReadDataAt(r, uint64(f.Cmdsz), offset)
	if err != nil {
		return false, err
	}
	c := iou.SafeSliceCap((*macho.Load)(nil), uint64(f.Ncmd))
	if c < 0 {
		return false, errors.New("too many load commands")
	}
	f.Loads = make([]macho.Load, 0, c)
	bo := f.ByteOrder
	for i := uint32(0); i < f.Ncmd; i++ {
		// Each load command begins with uint32 command and length.
		if len(dat) < 8 {
			return false, errors.New("command block too small")
		}
		cmd, siz := macho.LoadCmd(bo.Uint32(dat[0:4])), bo.Uint32(dat[4:8])
		if siz < 8 || siz > uint32(len(dat)) {
			return false, errors.New("invalid command block size")
		}
		var cmddat []byte
		cmddat, dat = dat[0:siz], dat[siz:]
		offset += int64(siz)
		var s *macho.Segment
		switch cmd {

		case macho.LoadCmdSymtab:
			var hdr macho.SymtabCmd
			b := bytes.NewReader(cmddat)
			if err := binary.Read(b, bo, &hdr); err != nil {
				return false, err
			}
			strtab := make([]byte, hdr.Strsize)
			if _, err := r.ReadAt(strtab, int64(hdr.Stroff)); err != nil {
				return false, err
			}
			var symsz int
			if f.Magic == macho.Magic64 {
				symsz = 16
			} else {
				symsz = 12
			}
			symdat, err := iou.SafeReadDataAt(r, uint64(hdr.Nsyms)*uint64(symsz), int64(hdr.Symoff))
			if err != nil {
				return false, err
			}
			bo := f.ByteOrder
			c := iou.SafeSliceCap((*macho.Symbol)(nil), uint64(hdr.Nsyms))
			if c < 0 {
				return false, errors.New("too many symbols")
			}
			symtab := make([]macho.Symbol, 0, c)
			b2 := bytes.NewReader(symdat)
			for i := 0; i < int(hdr.Nsyms); i++ {
				var n macho.Nlist64
				if f.Magic == macho.Magic64 {
					if err := binary.Read(b2, bo, &n); err != nil {
						return false, err
					}
				} else {
					var n32 macho.Nlist32
					if err := binary.Read(b2, bo, &n32); err != nil {
						return false, err
					}
					n.Name = n32.Name
					n.Type = n32.Type
					n.Sect = n32.Sect
					n.Desc = n32.Desc
					n.Value = uint64(n32.Value)
				}
				if n.Name >= uint32(len(strtab)) {
					return false, errors.New("invalid name in symbol table")
				}
				// We add "_" to Go symbols. Strip it here. See issue 33808.
				name := cstring(strtab[n.Name:])
				if strings.Contains(name, ".") && name[0] == '_' {
					name = name[1:]
				}
				symtab = append(symtab, macho.Symbol{
					Name:  name,
					Type:  n.Type,
					Sect:  n.Sect,
					Desc:  n.Desc,
					Value: n.Value,
				})
			}
			st := new(macho.Symtab)
			st.LoadBytes = macho.LoadBytes(cmddat)
			st.Syms = symtab
			f.Loads = append(f.Loads, st)
			f.Symtab = st
		}
		if s != nil {
			if int64(s.Offset) < 0 {
				return false, errors.New("invalid section offset")
			}
			if int64(s.Filesz) < 0 {
				return false, errors.New("invalid section file size")
			}
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

func debugWait() {
	log.Println(os.Getpid())
	bufio.NewScanner(os.Stdin).Scan()
}

func main() {
	debugWait()

	pn := rtu.GetCaller(0)
	pu := reflect.ValueOf(&marker).Pointer()

	name := os.Args[0]

	var f ExeFile
	var clo Closer

	//f, clo = check.Must2(OpenExeFile(name))
	f = machoExeFile2{name: name}

	defer func() {
		if clo != nil {
			check.Must(clo())
		}
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

	debugWait()
}
