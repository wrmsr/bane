package sql

import (
	sqb "github.com/wrmsr/bane/pkg/util/sql/base"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

type riState int8

const (
	riNew riState = iota
	riReady
	riLoaded
	riDone
)

type RowsIterator struct {
	br sqb.Rows
	st riState
	rr RowReader

	e error
	r Row
}

func IterRows(br sqb.Rows) *RowsIterator {
	return &RowsIterator{br: br}
}

func (ri *RowsIterator) Columns() ([]sqb.Column, error) {
	return ri.br.Columns()
}

func (ri *RowsIterator) Close() error {
	return ri.br.Close()
}

var _ bt.Iterator[bt.Result[Row]] = &RowsIterator{}

func (ri *RowsIterator) Iterate() bt.Iterator[bt.Result[Row]] {
	return ri
}

func (ri *RowsIterator) HasNext() bool {
	if ri.st == riDone {
		return false
	}
	if ri.st == riLoaded {
		return true
	}

	setErr := func(err error) bool {
		ri.e = err
		ri.st = riLoaded
		return true
	}

	if ri.st == riNew {
		if err := ri.br.Err(); err != nil {
			return setErr(err)
		}

		rr, err := NewRowReader(ri.br)
		if err != nil {
			return setErr(err)
		}
		ri.rr = rr
	} else if ri.st != riReady {
		panic("internal state error")
	}

	nxt := ri.br.Next()
	if err := ri.br.Err(); err != nil {
		return setErr(err)
	}

	if !nxt {
		ri.st = riDone
		return false
	}

	r, err := ri.rr.Read()
	if err != nil {
		return setErr(err)
	}

	ri.r = r
	ri.st = riLoaded
	return true
}

func (ri *RowsIterator) Next() bt.Result[Row] {
	if ri.st == riReady {
		ri.HasNext()
	}

	if ri.st == riDone {
		panic(bt.IteratorExhaustedError{})
	} else if ri.st != riLoaded {
		panic("internal state error")
	}

	if ri.e != nil {
		ri.st = riDone
		return bt.Err[Row](ri.e)
	}

	r := ri.r
	ri.r = Row{}
	ri.st = riReady
	return bt.Ok(r)
}
