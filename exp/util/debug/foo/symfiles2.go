// Copyright 2022 The Go Authors. All rights reserved. Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.
package main

import (
	"bytes"
	"debug/elf"
	"debug/macho"
	"encoding/binary"
	"errors"
	"io"
	"os"
	"strings"

	eu "github.com/wrmsr/bane/pkg/util/errors"
	iou "github.com/wrmsr/bane/pkg/util/io"
)

//

type machoSymFile2 struct {
	name string
}

var _ SymFile = machoSymFile2{}

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

var readChunkSize = uint64(128 * 1024)

func (f2 machoSymFile2) ForEachSym(fn func(s FileSym) bool) (ret bool, err error) {
	r, err := os.Open(f2.name)
	if err != nil {
		return
	}
	defer eu.AppendInvoke(&err, eu.Close(r))

	sr := io.NewSectionReader(r, 0, 1<<63-1)

	var f macho.File

	// Read and decode Mach magic to determine byte order, size. Magic32 and Magic64 differ only in the bottom bit.
	var ident [4]byte
	if _, err = r.ReadAt(ident[0:], 0); err != nil {
		return
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

	if err := binary.Read(sr, f.ByteOrder, &f.FileHeader); err != nil {
		return false, err
	}

	offset := int64(machoFileHeaderSize32)
	if f.Magic == macho.Magic64 {
		offset = machoFileHeaderSize64
	}
	dat, err := iou.SafeReadDataAt(r, uint64(f.Cmdsz), offset, readChunkSize)
	if err != nil {
		return false, err
	}
	c := iou.SafeSliceCap((*macho.Load)(nil), uint64(f.Ncmd), readChunkSize)
	if c < 0 {
		return false, errors.New("too many load commands")
	}
	f.Loads = make([]macho.Load, 0, c)
	bo := f.ByteOrder
	for i := uint32(0); i < f.Ncmd; i++ {
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
			// FIXME: block caching file wrapper
			if _, err := r.ReadAt(strtab, int64(hdr.Stroff)); err != nil {
				return false, err
			}
			var symsz int
			if f.Magic == macho.Magic64 {
				symsz = 16
			} else {
				symsz = 12
			}
			// FIXME: block caching file wrapper
			symdat, err := iou.SafeReadDataAt(r, uint64(hdr.Nsyms)*uint64(symsz), int64(hdr.Symoff), readChunkSize)
			if err != nil {
				return false, err
			}
			bo := f.ByteOrder
			c := iou.SafeSliceCap((*macho.Symbol)(nil), uint64(hdr.Nsyms), readChunkSize)
			if c < 0 {
				return false, errors.New("too many symbols")
			}
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
				if !fn(FileSym{
					Name: name,
					Addr: n.Value,
				}) {
					return false, nil
				}
			}
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

type elfSymFile2 struct {
	name string
}

var _ SymFile = elfSymFile2{}

const (
	elfSeekStart   int = 0
	elfSeekCurrent int = 1
	elfSeekEnd     int = 2
)

func (f2 elfSymFile2) ForEachSym(fn func(s FileSym) bool) (ret bool, err error) {
	r, err := os.Open(f2.name)
	if err != nil {
		return
	}
	defer eu.AppendInvoke(&err, eu.Close(r))

	sr := io.NewSectionReader(r, 0, 1<<63-1)
	// Read and decode ELF identifier
	var ident [16]uint8
	if _, err = r.ReadAt(ident[0:], 0); err != nil {
		return
	}
	if ident[0] != '\x7f' || ident[1] != 'E' || ident[2] != 'L' || ident[3] != 'F' {
		return false, errors.New("bad magic number")
	}

	f := new(elf.File)
	f.Class = elf.Class(ident[elf.EI_CLASS])
	switch f.Class {
	case elf.ELFCLASS32:
	case elf.ELFCLASS64:
		// ok
	default:
		return false, errors.New("unknown ELF class")
	}

	f.Data = elf.Data(ident[elf.EI_DATA])
	switch f.Data {
	case elf.ELFDATA2LSB:
		f.ByteOrder = binary.LittleEndian
	case elf.ELFDATA2MSB:
		f.ByteOrder = binary.BigEndian
	default:
		return false, errors.New("unknown ELF data encoding")
	}

	f.Version = elf.Version(ident[elf.EI_VERSION])
	if f.Version != elf.EV_CURRENT {
		return false, errors.New("unknown ELF version")
	}

	f.OSABI = elf.OSABI(ident[elf.EI_OSABI])
	f.ABIVersion = ident[elf.EI_ABIVERSION]

	// Read ELF file header
	var phoff int64
	var phentsize, phnum int
	var shoff int64
	var shentsize, shnum, shstrndx int
	switch f.Class {
	case elf.ELFCLASS32:
		hdr := new(elf.Header32)
		sr.Seek(0, elfSeekStart)
		if err = binary.Read(sr, f.ByteOrder, hdr); err != nil {
			return
		}
		f.Type = elf.Type(hdr.Type)
		f.Machine = elf.Machine(hdr.Machine)
		f.Entry = uint64(hdr.Entry)
		if v := elf.Version(hdr.Version); v != f.Version {
			return false, errors.New("mismatched ELF version")
		}
		phoff = int64(hdr.Phoff)
		phentsize = int(hdr.Phentsize)
		phnum = int(hdr.Phnum)
		shoff = int64(hdr.Shoff)
		shentsize = int(hdr.Shentsize)
		shnum = int(hdr.Shnum)
		shstrndx = int(hdr.Shstrndx)
	case elf.ELFCLASS64:
		hdr := new(elf.Header64)
		sr.Seek(0, elfSeekStart)
		if err = binary.Read(sr, f.ByteOrder, hdr); err != nil {
			return
		}
		f.Type = elf.Type(hdr.Type)
		f.Machine = elf.Machine(hdr.Machine)
		f.Entry = hdr.Entry
		if v := elf.Version(hdr.Version); v != f.Version {
			return false, errors.New("mismatched ELF version")
		}
		phoff = int64(hdr.Phoff)
		phentsize = int(hdr.Phentsize)
		phnum = int(hdr.Phnum)
		shoff = int64(hdr.Shoff)
		shentsize = int(hdr.Shentsize)
		shnum = int(hdr.Shnum)
		shstrndx = int(hdr.Shstrndx)
	}

	if shoff < 0 {
		return false, errors.New("invalid shoff")
	}
	if phoff < 0 {
		return false, errors.New("invalid phoff")
	}

	if shoff == 0 && shnum != 0 {
		return false, errors.New("invalid ELF shnum for shoff=0")
	}

	if shnum > 0 && shstrndx >= shnum {
		return false, errors.New("invalid ELF shstrndx")
	}

	var wantPhentsize, wantShentsize int
	switch f.Class {
	case elf.ELFCLASS32:
		wantPhentsize = 8 * 4
		wantShentsize = 10 * 4
	case elf.ELFCLASS64:
		wantPhentsize = 2*4 + 6*8
		wantShentsize = 4*4 + 6*8
	}
	if phnum > 0 && phentsize < wantPhentsize {
		return false, errors.New("invalid ELF phentsize")
	}

	// Read program headers
	f.Progs = make([]*elf.Prog, phnum)
	for i := 0; i < phnum; i++ {
		off := phoff + int64(i)*int64(phentsize)
		sr.Seek(off, elfSeekStart)
		p := new(elf.Prog)
		switch f.Class {
		case elf.ELFCLASS32:
			ph := new(elf.Prog32)
			if err = binary.Read(sr, f.ByteOrder, ph); err != nil {
				return
			}
			p.ProgHeader = elf.ProgHeader{
				Type:   elf.ProgType(ph.Type),
				Flags:  elf.ProgFlag(ph.Flags),
				Off:    uint64(ph.Off),
				Vaddr:  uint64(ph.Vaddr),
				Paddr:  uint64(ph.Paddr),
				Filesz: uint64(ph.Filesz),
				Memsz:  uint64(ph.Memsz),
				Align:  uint64(ph.Align),
			}
		case elf.ELFCLASS64:
			ph := new(elf.Prog64)
			if err = binary.Read(sr, f.ByteOrder, ph); err != nil {
				return
			}
			p.ProgHeader = elf.ProgHeader{
				Type:   elf.ProgType(ph.Type),
				Flags:  elf.ProgFlag(ph.Flags),
				Off:    ph.Off,
				Vaddr:  ph.Vaddr,
				Paddr:  ph.Paddr,
				Filesz: ph.Filesz,
				Memsz:  ph.Memsz,
				Align:  ph.Align,
			}
		}
		if int64(p.Off) < 0 {
			return false, errors.New("invalid program header offset")
		}
		if int64(p.Filesz) < 0 {
			return false, errors.New("invalid program header file size")
		}
		p.sr = io.NewSectionReader(r, int64(p.Off), int64(p.Filesz))
		p.ReaderAt = p.sr
		f.Progs[i] = p
	}

	// If the number of sections is greater than or equal to SHN_LORESERVE
	// (0xff00), shnum has the value zero and the actual number of section
	// header table entries is contained in the sh_size field of the section
	// header at index 0.
	if shoff > 0 && shnum == 0 {
		var typ, link uint32
		sr.Seek(shoff, elfSeekStart)
		switch f.Class {
		case elf.ELFCLASS32:
			sh := new(elf.Section32)
			if err = binary.Read(sr, f.ByteOrder, sh); err != nil {
				return
			}
			shnum = int(sh.Size)
			typ = sh.Type
			link = sh.Link
		case elf.ELFCLASS64:
			sh := new(elf.Section64)
			if err = binary.Read(sr, f.ByteOrder, sh); err != nil {
				return
			}
			shnum = int(sh.Size)
			typ = sh.Type
			link = sh.Link
		}
		if elf.SectionType(typ) != elf.SHT_NULL {
			return false, errors.New("invalid type of the initial section")
		}

		if shnum < int(elf.SHN_LORESERVE) {
			return false, errors.New("invalid ELF shnum contained in sh_size")
		}

		// If the section name string table section index is greater than or
		// equal to SHN_LORESERVE (0xff00), this member has the value
		// SHN_XINDEX (0xffff) and the actual index of the section name
		// string table section is contained in the sh_link field of the
		// section header at index 0.
		if shstrndx == int(elf.SHN_XINDEX) {
			shstrndx = int(link)
			if shstrndx < int(elf.SHN_LORESERVE) {
				return false, errors.New("invalid ELF shstrndx contained in sh_link")
			}
		}
	}

	if shnum > 0 && shentsize < wantShentsize {
		return false, errors.New("invalid ELF shentsize")
	}

	// Read section headers
	c := iou.SafeSliceCap((*elf.Section)(nil), uint64(shnum), readChunkSize)
	if c < 0 {
		return false, errors.New("too many sections")
	}
	f.Sections = make([]*elf.Section, 0, c)
	names := make([]uint32, 0, c)
	for i := 0; i < shnum; i++ {
		off := shoff + int64(i)*int64(shentsize)
		sr.Seek(off, elfSeekStart)
		s := new(elf.Section)
		switch f.Class {
		case elf.ELFCLASS32:
			sh := new(elf.Section32)
			if err = binary.Read(sr, f.ByteOrder, sh); err != nil {
				return
			}
			names = append(names, sh.Name)
			s.SectionHeader = elf.SectionHeader{
				Type:      elf.SectionType(sh.Type),
				Flags:     elf.SectionFlag(sh.Flags),
				Addr:      uint64(sh.Addr),
				Offset:    uint64(sh.Off),
				FileSize:  uint64(sh.Size),
				Link:      sh.Link,
				Info:      sh.Info,
				Addralign: uint64(sh.Addralign),
				Entsize:   uint64(sh.Entsize),
			}
		case elf.ELFCLASS64:
			sh := new(elf.Section64)
			if err = binary.Read(sr, f.ByteOrder, sh); err != nil {
				return
			}
			names = append(names, sh.Name)
			s.SectionHeader = elf.SectionHeader{
				Type:      elf.SectionType(sh.Type),
				Flags:     elf.SectionFlag(sh.Flags),
				Offset:    sh.Off,
				FileSize:  sh.Size,
				Addr:      sh.Addr,
				Link:      sh.Link,
				Info:      sh.Info,
				Addralign: sh.Addralign,
				Entsize:   sh.Entsize,
			}
		}
		if int64(s.Offset) < 0 {
			return false, errors.New("invalid section offset")
		}
		if int64(s.FileSize) < 0 {
			return false, errors.New("invalid section size")
		}
		s.sr = io.NewSectionReader(r, int64(s.Offset), int64(s.FileSize))

		if s.Flags&elf.SHF_COMPRESSED == 0 {
			s.ReaderAt = s.sr
			s.Size = s.FileSize
		} else {
			// Read the compression header.
			switch f.Class {
			case elf.ELFCLASS32:
				ch := new(elf.Chdr32)
				if err = binary.Read(s.sr, f.ByteOrder, ch); err != nil {
					return
				}
				s.compressionType = elf.CompressionType(ch.Type)
				s.Size = uint64(ch.Size)
				s.Addralign = uint64(ch.Addralign)
				s.compressionOffset = int64(binary.Size(ch))
			case elf.ELFCLASS64:
				ch := new(elf.Chdr64)
				if err = binary.Read(s.sr, f.ByteOrder, ch); err != nil {
					return
				}
				s.compressionType = elf.CompressionType(ch.Type)
				s.Size = ch.Size
				s.Addralign = ch.Addralign
				s.compressionOffset = int64(binary.Size(ch))
			}
		}

		f.Sections = append(f.Sections, s)
	}

	panic("implement me")
}