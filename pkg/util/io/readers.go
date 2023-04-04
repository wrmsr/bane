package io

import "io"

//

type OffsetReaderAt struct {
	Off int64
	R   io.ReaderAt
}

var _ io.ReaderAt = OffsetReaderAt{}

func (r OffsetReaderAt) ReadAt(p []byte, off int64) (n int, err error) {
	return r.R.ReadAt(p, off+r.Off)
}

//

type ReaderAtReader struct {
	Off int64
	R   io.ReaderAt
}

var _ io.Reader = &ReaderAtReader{}

func (r *ReaderAtReader) Read(p []byte) (n int, err error) {
	n, err = r.R.ReadAt(p, r.Off)
	r.Off += int64(n)
	return
}
