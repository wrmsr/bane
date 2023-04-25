package reading

import (
	"errors"
	"io"
)

type ByteReader struct {
	s []byte
	i int64
}

var _ io.Reader = &ByteReader{}

func (r *ByteReader) Len() int {
	if r.i >= int64(len(r.s)) {
		return 0
	}
	return int(int64(len(r.s)) - r.i)
}

func (r *ByteReader) Size() int64 { return int64(len(r.s)) }

func (r *ByteReader) Tell() int64 { return r.i }

func (r *ByteReader) Read(b []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[r.i:])
	r.i += int64(n)
	return
}

func (r *ByteReader) ReadAt(b []byte, off int64) (n int, err error) {
	if off < 0 {
		return 0, errors.New("negative offset")
	}
	if off >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(b, r.s[off:])
	if n < len(b) {
		err = io.EOF
	}
	return
}

func (r *ByteReader) ReadByte() (byte, error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	b := r.s[r.i]
	r.i++
	return b, nil
}

func (r *ByteReader) Skip(n int64) error {
	if (r.i + n) > int64(len(r.s)) {
		return io.EOF
	}
	r.i += n
	return nil
}

func NewByteReader(b []byte) *ByteReader {
	return &ByteReader{b, 0}
}
