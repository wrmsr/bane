package io

import (
	"bufio"
	"io"

	"github.com/wrmsr/bane/core/sync/pools"
)

//

type PooledBufioWriter struct {
	u io.Writer
	p pools.Pool[*bufio.Writer]
	b *bufio.Writer
}

var _ io.Writer = &PooledBufioWriter{}

func NewPooledBufioWriter(
	u io.Writer,
	p pools.Pool[*bufio.Writer],
	b *bufio.Writer,
) *PooledBufioWriter {
	return &PooledBufioWriter{
		u: u,
		p: p,
		b: b,
	}
}

func (pw *PooledBufioWriter) Pooled() *bufio.Writer {
	if pw.p == nil {
		return nil
	}
	return pw.b
}

func (pw *PooledBufioWriter) Get() *bufio.Writer {
	if pw.b == nil {
		if pw.p != nil {
			pw.b = pw.p.Get()
			pw.b.Reset(pw.u)
		} else {
			pw.b = bufio.NewWriter(pw.u)
		}
	}
	return pw.b
}

func (pw *PooledBufioWriter) Write(p []byte) (n int, err error) {
	panic("implement me")
}

func (pw *PooledBufioWriter) Flush() error {
	if pw.b == nil {
		return nil
	}

	if pw.p != nil {
		err := pw.b.Flush()
		pw.b.Reset(io.Discard)
		pw.p.Put(pw.b)
		pw.b = nil
		return err

	} else {
		if pw.b.Buffered() < 1 {
			return nil
		}
		return pw.b.Flush()
	}
}
