// Copyright 2022 The Go Authors. All rights reserved. Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.
package main

import (
	"bytes"
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
			if _, err := r.ReadAt(strtab, int64(hdr.Stroff)); err != nil {
				return false, err
			}
			var symsz int
			if f.Magic == macho.Magic64 {
				symsz = 16
			} else {
				symsz = 12
			}
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
