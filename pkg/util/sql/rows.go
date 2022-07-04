package sql

import (
	its "github.com/wrmsr/bane/pkg/util/iterators"
	sqb "github.com/wrmsr/bane/pkg/util/sql/base"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

type Row struct {
}

type RowsIterator struct {
	b sqb.Rows
}

func Iter(b sqb.Rows) *RowsIterator {
	return &RowsIterator{b: b}
}

func (ri *RowsIterator) Columns() ([]sqb.Column, error) {
	return ri.Columns()
}

func (ri *RowsIterator) Close() error {
	return ri.Close()
}

var _ its.Iterator[bt.Result[Row]] = &RowsIterator{}

func (ri *RowsIterator) Iterate() its.Iterator[bt.Result[Row]] {
	panic("implement me")
}

func (ri *RowsIterator) HasNext() bool {
	panic("implement me")
}

func (ri *RowsIterator) Next() bt.Result[Row] {
	panic("implement me")
}
