// Copyright 2022 The Go Authors. All rights reserved. Use of this source code is governed by a BSD-style license that
// can be found in the LICENSE file.
package io

import (
	"io"
	"reflect"
)

// ReadData reads n bytes from the input stream, but avoids allocating all n bytes if n is large. This avoids crashing
// the program by allocating all n bytes in cases where n is incorrect.
//
// The error is io.EOF only if no bytes were read. If an io.EOF happens after reading some but not all the bytes,
// ReadData returns io.ErrUnexpectedEOF.
func SafeReadData(r io.Reader, n uint64, chunkSize uint64) ([]byte, error) {
	if int64(n) < 0 || n != uint64(int(n)) {
		// n is too large to fit in int, so we can't allocate a buffer large enough. Treat this as a read failure.
		return nil, io.ErrUnexpectedEOF
	}

	if n < chunkSize {
		buf := make([]byte, n)
		_, err := io.ReadFull(r, buf)
		if err != nil {
			return nil, err
		}
		return buf, nil
	}

	var buf []byte
	buf1 := make([]byte, chunkSize)
	for n > 0 {
		next := n
		if next > chunkSize {
			next = chunkSize
		}
		_, err := io.ReadFull(r, buf1[:next])
		if err != nil {
			if len(buf) > 0 && err == io.EOF {
				err = io.ErrUnexpectedEOF
			}
			return nil, err
		}
		buf = append(buf, buf1[:next]...)
		n -= next
	}
	return buf, nil
}

// ReadDataAt reads n bytes from the input stream at off, but avoids allocating all n bytes if n is large. This avoids
// crashing the program by allocating all n bytes in cases where n is incorrect.
func SafeReadDataAt(r io.ReaderAt, n uint64, off int64, chunkSize uint64) ([]byte, error) {
	if int64(n) < 0 || n != uint64(int(n)) {
		// n is too large to fit in int, so we can't allocate a buffer large enough. Treat this as a read failure.
		return nil, io.ErrUnexpectedEOF
	}

	if n < chunkSize {
		buf := make([]byte, n)
		_, err := r.ReadAt(buf, off)
		if err != nil {
			// io.SectionReader can return EOF for n == 0, but for our purposes that is a success.
			if err != io.EOF || n > 0 {
				return nil, err
			}
		}
		return buf, nil
	}

	var buf []byte
	buf1 := make([]byte, chunkSize)
	for n > 0 {
		next := n
		if next > chunkSize {
			next = chunkSize
		}
		_, err := r.ReadAt(buf1[:next], off)
		if err != nil {
			return nil, err
		}
		buf = append(buf, buf1[:next]...)
		n -= next
		off += int64(next)
	}
	return buf, nil
}

// SliceCap returns the capacity to use when allocating a slice. After the slice is allocated with the capacity, it
// should be built using append. This will avoid allocating too much memory if the capacity is large and incorrect.
//
// A negative result means that the value is always too big.
//
// The element type is described by passing a pointer to a value of that type. This would ideally use generics, but this
// code is built with the bootstrap compiler which need not support generics. We use a pointer so that we can handle
// slices of interface type.
func SafeSliceCap(v any, c uint64, chunkSize uint64) int {
	if int64(c) < 0 || c != uint64(int(c)) {
		return -1
	}
	typ := reflect.TypeOf(v)
	if typ.Kind() != reflect.Ptr {
		panic("SliceCap called with non-pointer type")
	}
	size := uint64(typ.Elem().Size())
	if size > 0 && c > (1<<64-1)/size {
		return -1
	}
	if c*size > chunkSize {
		c = chunkSize / size
		if c == 0 {
			c = 1
		}
	}
	return int(c)
}
