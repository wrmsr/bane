package main

import "io"

///

type BytesOrString struct {
	b []byte
	s string
}

func (bs BytesOrString) Bytes() []byte {
	if bs.b != nil {
		return bs.b
	}
	if bs.s == "" {
		return nil
	}
	return []byte(bs.s)
}

func (bs BytesOrString) String() string {
	if bs.s != "" {
		return bs.s
	}
	if len(bs.b) < 1 {
		return ""
	}
	return string(bs.b)
}

///

type SliceView[T any] interface {
	IsNil() bool
	Len() int
	At(i int) (T, error)
	Extract(l, r int) ([]T, error)
}

//

type sliceView[T any] struct {
	s []T
}

var _ SliceView[int] = sliceView[int]{}

func (v sliceView[T]) IsNil() bool {
	return v.s == nil
}

func (v sliceView[T]) Len() int {
	return len(v.s)
}

func (v sliceView[T]) At(i int) (T, error) {
	return v.s[i], nil
}

func (v sliceView[T]) Extract(l, r int) ([]T, error) {
	if r < 0 {
		return v.s[l:], nil
	}
	return v.s[l:r], nil
}

//

type readerAtSliceView struct {
	r io.ReaderAt
	l int
}

var _ SliceView[byte] = readerAtSliceView{}

func (v readerAtSliceView) IsNil() bool {
	return v.l < 1
}

func (v readerAtSliceView) Len() int {
	return v.l
}

func (v readerAtSliceView) At(i int) (byte, error) {
	var b [1]byte
	_, err := v.r.ReadAt(b[:], int64(i))
	if err != nil {
		return 0, err
	}
	return b[0], nil
}

func (v readerAtSliceView) Extract(l, r int) ([]byte, error) {
	var c int
	if r < 0 {
		c = v.l - l
	} else {
		c = r - l
	}
	b := make([]byte, c)
	_, err := v.r.ReadAt(b, int64(l))
	if err != nil {
		return nil, err
	}
	return b, nil
}

//

type sliceViewReaderAt struct {
	sv SliceView[byte]
}

var _ io.ReaderAt = sliceViewReaderAt{}

func (r sliceViewReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	b, err := r.sv.Extract(int(off), int(off)+len(p))
	if err != nil {
		return
	}
	n = len(b)
	copy(p, b)
	return
}
