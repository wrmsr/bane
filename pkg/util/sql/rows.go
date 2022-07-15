package sql

import (
	"reflect"

	rfl "github.com/wrmsr/bane/pkg/util/reflect"
	sqb "github.com/wrmsr/bane/pkg/util/sql/base"
	bt "github.com/wrmsr/bane/pkg/util/types"
)

//

type RowScanner func(dst ...any) error

func ScanMap(cs []sqb.Column, scan RowScanner) (map[string]any, error) {
	vvs := make([]reflect.Value, len(cs))
	dvs := make([]any, len(cs))
	for i, c := range cs {
		sv := reflect.New(c.Type)
		vvs[i] = sv.Elem()
		dvs[i] = sv.Interface()
	}
	if err := scan(dvs...); err != nil {
		return nil, err
	}
	m := make(map[string]any, len(cs))
	for i, c := range cs {
		m[c.Name] = vvs[i].Interface()
	}
	return m, nil
}

func MakeAnyColumns(names ...string) []sqb.Column {
	cs := make([]sqb.Column, len(names))
	for i, n := range names {
		cs[i] = sqb.Column{Name: n, Type: rfl.Any()}
	}
	return cs
}

//

type Row struct {
}

//

type riState int8

const (
	riNew riState = iota
	riReady
	riLoaded
	riDone
)

type RowsIterator struct {
	b  sqb.Rows
	st riState
	e  error
}

func Iter(b sqb.Rows) *RowsIterator {
	return &RowsIterator{b: b}
}

func (ri *RowsIterator) Columns() ([]sqb.Column, error) {
	return ri.b.Columns()
}

func (ri *RowsIterator) Close() error {
	return ri.b.Close()
}

var _ bt.Iterator[bt.Result[Row]] = &RowsIterator{}

func (ri *RowsIterator) Iterate() bt.Iterator[bt.Result[Row]] {
	return ri
}

func (ri *RowsIterator) HasNext() bool {
	if ri.st == riDone {
		return false
	}
	if ri.st == riNew {
		if err := ri.b.Err(); err != nil {
			ri.e = err
			ri.st = riLoaded
			return true
		}
	}
	if ri.st != riLoaded {
		if ri.b.Next() {
			ri.st = riLoaded
		} else {
			ri.st = riDone
		}
		if err := ri.b.Err(); err != nil {
			ri.e = err
			ri.st = riLoaded
		}
	}
	return ri.st == riLoaded
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
	ri.st = riReady
	return bt.Ok(Row{})
}
